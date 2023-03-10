package main

import (
    "fmt"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

type User struct {
    ID   uint
    Name string
    Age  int
}

func main() {
    // 连接MySQL数据库
    dsn := "username:password@tcp(127.0.0.1:3306)/database_name?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        fmt.Printf("Failed to connect to database: %v", err)
        return
    }
    defer db.Close()

    // 自动迁移表结构
    err = db.AutoMigrate(&User{})
    if err != nil {
        fmt.Printf("Failed to migrate table: %v", err)
        return
    }

    // 创建用户
    user := User{Name: "Alice", Age: 18}
    result := db.Create(&user)
    if result.Error != nil {
        fmt.Printf("Failed to create user: %v", result.Error)
        return
    }
    fmt.Printf("User created: %v\n", user)

    // 查询用户
    var users []User
    result = db.Find(&users)
    if result.Error != nil {
        fmt.Printf("Failed to query users: %v", result.Error)
        return
    }
    fmt.Printf("Users found: %v\n", users)

    // 更新用户
    user.Age = 20
    result = db.Save(&user)
    if result.Error != nil {
        fmt.Printf("Failed to update user: %v", result.Error)
        return
    }
    fmt.Printf("User updated: %v\n", user)

    // 删除用户
    result = db.Delete(&user)
    if result.Error != nil {
        fmt.Printf("Failed to delete user: %v", result.Error)
        return
    }
    fmt.Printf("User deleted: %v\n", user)
}

