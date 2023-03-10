package main

import (
    "fmt"
    "github.com/segmentio/kafka-go"
)

func main() {
    // 设置Kafka连接信息
    broker := "localhost:9092"
    topic := "test_topic"

    // 创建Kafka读取器
    r := kafka.NewReader(kafka.ReaderConfig{
        Brokers: []string{broker},
        Topic:   topic,
        GroupID: "test_group",
    })

    // 读取Kafka消息
    for {
        m, err := r.ReadMessage(context.Background())
        if err != nil {
            fmt.Printf("error while reading message: %v\n", err)
            break
        }
        fmt.Printf("message received: %s\n", string(m.Value))
    }

    // 关闭Kafka读取器
    r.Close()
}

