package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	Name string `bson:"name"`
	Age  int    `bson:"age"`
}

func main() {
	// 创建MongoDB客户端
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(context.Background())

	// 获取MongoDB集合
	collection := client.Database("test").Collection("users")

	// 插入文档
	user := User{Name: "Alice", Age: 18}
	result, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Inserted document with ID: %v\n", result.InsertedID)

	// 查询文档
	filter := bson.M{"name": "Alice"}
	var users []User
	cur, err := collection.Find(context.Background(), filter)
	if err != nil {
		panic(err)
	}
	defer cur.Close(context.Background())
	for cur.Next(context.Background()) {
		var user User
		err := cur.Decode(&user)
		if err != nil {
			panic(err)
		}
		users = append(users, user)
	}
	if err := cur.Err(); err != nil {
		panic(err)
	}
	fmt.Printf("Found documents: %v\n", users)

	// 更新文档
	filter = bson.M{"name": "Alice"}
	update := bson.M{"$set": bson.M{"age": 20}}
	result, err = collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Matched %v documents and updated %v documents.\n", result.MatchedCount, result.ModifiedCount)

	// 删除文档
	filter = bson.M{"name": "Alice"}
	result, err = collection.DeleteOne(context.Background(), filter)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Deleted %v documents.\n", result.DeletedCount)
}

