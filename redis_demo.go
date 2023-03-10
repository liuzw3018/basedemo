package main

import (
    "fmt"
    "github.com/go-redis/redis"
)

func main() {
    // 创建Redis客户端
    client := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "", // Redis密码
        DB:       0,  // Redis数据库编号
    })

    // 设置Redis键值对
    err := client.Set("mykey", "Hello Redis", 0).Err()
    if err != nil {
        fmt.Println("Error setting key-value: ", err)
    }

    // 获取Redis键值对
    value, err := client.Get("mykey").Result()
    if err != nil {
        fmt.Println("Error getting key-value: ", err)
    } else {
        fmt.Println("Value: ", value)
    }

    // 删除Redis键值对
    err = client.Del("mykey").Err()
    if err != nil {
        fmt.Println("Error deleting key-value: ", err)
    }
}

