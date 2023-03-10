
package main

import (
    "fmt"
    "github.com/Shopify/sarama"
    "github.com/Shopify/sarama/sasl"
)

func main() {
    // 设置Kafka连接信息
    broker := "kafka.example.com:9092"
    topic := "test_topic"
    username := "kafka-user"
    realm := "EXAMPLE.COM"
    keytabPath := "/path/to/keytab"

    // 创建Kerberos配置
    kerberosConfig := &sasl.KerberosConfig{
        ServiceName: "kafka",
        Realm:       realm,
        KeyTabPath:  keytabPath,
        Username:    username,
    }

    // 创建Kafka配置
    kafkaConfig := sarama.NewConfig()
    kafkaConfig.Net.SASL.Enable = true
    kafkaConfig.Net.SASL.Mechanism = sarama.SASLTypeGSSAPI
    kafkaConfig.Net.SASL.KerberosConfig = kerberosConfig
    kafkaConfig.Version = sarama.V2_0_0_0

    // 创建Kafka消费者
    consumer, err := sarama.NewConsumer([]string{broker}, kafkaConfig)
    if err != nil {
        fmt.Println("Error creating Kafka consumer:", err)
        return
    }
    defer consumer.Close()

    // 订阅Kafka topic
    partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
    if err != nil {
        fmt.Println("Error subscribing to Kafka topic:", err)
        return
    }
    defer partitionConsumer.Close()

    // 读取Kafka消息
    for msg := range partitionConsumer.Messages() {
        fmt.Printf("message received: %s\n", string(msg.Value))
    }
}

