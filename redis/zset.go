package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"strconv"
	"time"
)

func main() {
	xKey := "zset-test"
	ctx := context.Background()
	// 不能重复
	id := strconv.FormatInt(time.Now().Unix(), 10)

	// 设置任务的详细信息
	// 添加任务，设置id
	err := client.ZAdd(ctx, xKey, redis.Z{
		Score:  1000,
		Member: 1000,
	}).Err()
	if err != nil {
		log.Println("Failed to set task details:", err)
		return
	}

	result, err := client.XDel(ctx, xKey, id).Result()
	if err != nil {
		return
	}

	fmt.Println(result)
}
