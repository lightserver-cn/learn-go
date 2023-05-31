# delay_queue 延迟队列

使用 Redis Stream 实现延时队列的简单示例：

> Redis Stream 可以用作延时队列的实现。通过合理使用 Stream 的特性，我们可以实现具有延时功能的消息队列。下面是一种基本的实现方式：

1. 将消息写入 Stream，并设置一个适当的时间戳作为消息的 score 值。时间戳应该表示消息应该在何时被处理。
2. 消费者使用 XREAD 命令按顺序读取 Stream 中的消息，可以使用 BLOCK 参数来实现阻塞读取。
3. 在消费者读取到消息后，可以根据当前时间和消息的时间戳来判断是否应该立即处理消息，或者将消息重新放入 Stream 中等待后续处理。

在示例代码中，我们使用 Redis Stream 实现了一个简单的延时队列。我们创建了一个 Redis 客户端 client。

```go
// 创建 Redis 客户端
client := redis.NewClient(&redis.Options{
    Addr:     "localhost:6379",
    Password: "", // 如果没有设置密码，留空即可
    DB:       0,  // 使用默认数据库
})
```

首先，我们定义了延时队列的名称 queueName。然后，我们启动了一个协程来处理延时队列中的任务。在协程中，我们通过调用 XRangeN 命令获取延时队列中的任务，并根据任务的时间戳判断是否应该执行。如果任务的时间戳小于等于当前时间戳，表示任务已到期，我们执行具体的任务逻辑，并从延时队列中删除已处理的任务。

```go
// 创建延时队列
queueName := "delayed_queue"

// 启动一个协程，用于处理延时队列中的任务
go func() {
    for {
        now := time.Now().UnixNano() / int64(time.Millisecond) // 当前时间戳（毫秒）
        result, err := client.XRangeN(ctx, queueName, "-", "+", 0).Result()
        if err != nil {
            fmt.Println("获取延时任务失败:", err)
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
```

接下来，我们实现了一个 addJobToQueue 函数，用于向延时队列中添加任务。在函数中，我们计算任务的执行时间戳，并使用 XAdd 命令将任务添加到延时队列中。

```go
func addJobToQueue(client *redis.Client, queueName, jobID string, delay time.Duration) {
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
```

最后，我们添加了几个延时任务到队列，并等待任务执行完成。

```go
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
```

请注意，上述示例代码仅提供了基本的延时队列实现，实际应用中可能需要考虑更多的边界情况和错误处理。同时，该示例中的时间戳处理基于毫秒级精度，根据实际需求可以进行调整。
