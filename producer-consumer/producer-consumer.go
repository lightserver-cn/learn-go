package main

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

func main() {
	// 创建 Redis 客户端
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // 如果没有设置密码，留空即可
		DB:       0,  // 使用默认数据库
	})

	// 使用 XInfoGroups 检查消费者组是否存在
	queueName := "my_queue"
	consumerGroup := "my_group"

	// 检查消费者组是否存在
	groupExists := false
	groups, err := client.XInfoGroups(context.Background(), queueName).Result()
	if err != nil && err != redis.Nil {
		fmt.Println("检查消费者组失败:", err)
		return
	}

	for _, group := range groups {
		if group.Name == consumerGroup {
			groupExists = true
			break
		}
	}

	// 如果消费者组不存在，创建消费者组
	if !groupExists {
		_, err := client.XGroupCreateMkStream(context.Background(), queueName, consumerGroup, "0").Result()
		if err != nil {
			fmt.Println("创建消费者组失败:", err)
			return
		}
	}

	go func() {
		// 消费者从队列中接收消息
		consumerID := "consumer1"
		for {
			result, err := client.XReadGroup(context.Background(), &redis.XReadGroupArgs{
				Group:    consumerGroup,
				Consumer: consumerID,
				Streams:  []string{queueName, ">"},
				Count:    1,
				Block:    0, // 0 表示立即返回，阻塞时间可以根据实际需求设置
			}).Result()
			if err != nil {
				fmt.Println("接收消息失败:", err)
				return
			}

			for _, stream := range result {
				for _, message := range stream.Messages {
					fmt.Println("收到消息:", message.Values["message"])

					// 处理消息后确认消息已被消费
					_, err := client.XAck(context.Background(), queueName, consumerGroup, message.ID).Result()
					if err != nil {
						fmt.Println("确认消息失败:", err)
					}
				}
			}
		}
	}()

	// 生产者发送消息到队列
	_, err = client.XAdd(context.Background(), &redis.XAddArgs{
		Stream: queueName,
		Values: map[string]interface{}{
			"message": "第一条消息",
		},
	}).Result()
	if err != nil {
		fmt.Println("发送消息失败:", err)
		return
	}

	// 生产者发送消息到队列
	_, err = client.XAdd(context.Background(), &redis.XAddArgs{
		Stream: queueName,
		Values: map[string]interface{}{
			"message": "第二条消息",
		},
	}).Result()
	if err != nil {
		fmt.Println("发送消息失败:", err)
		return
	}

	time.Sleep(2 * time.Second)

	// 生产者发送消息到队列
	_, err = client.XAdd(context.Background(), &redis.XAddArgs{
		Stream: queueName,
		Values: map[string]interface{}{
			"message": "第三条消息",
		},
	}).Result()
	if err != nil {
		fmt.Println("发送消息失败:", err)
		return
	}

	// 等待任务执行
	time.Sleep(10 * time.Second)

	// 关闭 Redis 连接
	err = client.Close()
	if err != nil {
		fmt.Println("关闭 Redis 连接失败:", err)
	}
}
