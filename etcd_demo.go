package main

import (
	"context"
	"fmt"
	"time"

	"github.com/coreos/etcd/clientv3"
)

func main() {
	// 创建etcd客户端连接
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"}, // etcd节点地址
		DialTimeout: 5 * time.Second,             // 连接超时时间
	})
	if err != nil {
		panic(err)
	}
	defer client.Close()

	// 设置etcd键值对
	key := "foo"
	value := "bar"
	_, err = client.Put(context.Background(), key, value)
	if err != nil {
		panic(err)
	}

	// 获取etcd键值对
	resp, err := client.Get(context.Background(), key)
	if err != nil {
		panic(err)
	}
	for _, kv := range resp.Kvs {
		fmt.Printf("key: %s, value: %s\n", kv.Key, kv.Value)
	}

	// 删除etcd键值对
	_, err = client.Delete(context.Background(), key)
	if err != nil {
		panic(err)
	}
}

