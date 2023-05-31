package main

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

func main() {

	ctx := context.Background()
	// 创建 Redis 客户端
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // 如果没有设置密码，留空即可
		DB:       0,  // 使用默认数据库
	})

	// 创建延时队列
	queueName := "delayed_queue"

	// 启动一个协程，用于处理延时队列中的任务
	go func() {
		for {
			now := time.Now().UnixNano() / int64(time.Millisecond) // 当前时间戳（毫秒）
			result, err := client.XRangeN(ctx, queueName, "-", "+", 1).Result()
			if err != nil {
				if err != redis.Nil {
					fmt.Println("获取延时任务失败:", err)
				}
				continue
			}

			for _, entry := range result {
				jobID := entry.ID
				jobTimestamp, err := strconv.ParseInt(entry.Values["timestamp"].(string), 10, 64)
				if err != nil {
					fmt.Println("解析时间戳失败:", err)
					continue
				}

				if jobTimestamp <= now {
					fmt.Println("执行任务:", jobID)
					// 执行具体的任务逻辑

					// 从延时队列中删除已处理的任务
					_, err := client.XDel(ctx, queueName, jobID).Result()
					if err != nil {
						fmt.Println("删除任务失败:", err)
					}
				}
			}

			// 等待一段时间继续轮询
			time.Sleep(time.Second)
		}
	}()

	// 添加延时任务到队列
	addJobToQueue(client, queueName, "job1", 5*time.Second)
	addJobToQueue(client, queueName, "job2", 10*time.Second)
	addJobToQueue(client, queueName, "job3", 15*time.Second)

	// 等待任务执行
	time.Sleep(20 * time.Second)

	// 关闭 Redis 连接
	err := client.Close()
	if err != nil {
		fmt.Println("关闭 Redis 连接失败:", err)
	}
}

func addJobToQueue(client *redis.Client, queueName, jobID string, delay time.Duration) {
	ctx := context.Background()

	// 计算任务的执行时间戳
	executeAt := time.Now().Add(delay).UnixNano() / int64(time.Millisecond)

	// 将任务添加到延时队列中
	err := client.XAdd(ctx, &redis.XAddArgs{
		Stream: queueName,
		ID:     "*",
		Values: map[string]interface{}{
			"job_id":    jobID,
			"timestamp": strconv.FormatInt(executeAt, 10),
		},
	}).Err()
	if err != nil {
		fmt.Println("添加任务到延时队列失败:", err)
	}

	fmt.Println("添加任务到延时队列:", jobID)
}
