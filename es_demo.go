package main

import (
    "context"
    "fmt"
    "github.com/elastic/go-elasticsearch/v7"
    "github.com/elastic/go-elasticsearch/v7/esapi"
    "encoding/json"
)

func main() {
    // 创建Elasticsearch客户端
    es, err := elasticsearch.NewDefaultClient()
    if err != nil {
        fmt.Println("Error creating the client: ", err)
        return
    }

    // 创建索引
    indexName := "my_index"
    createIndexRequest := esapi.IndicesCreateRequest{
        Index: indexName,
    }
    res, err := createIndexRequest.Do(context.Background(), es)
    if err != nil {
        fmt.Println("Error creating the index: ", err)
        return
    }
    defer res.Body.Close()

    // 添加文档
    docID := "1"
    doc := map[string]interface{}{
        "title":  "My Document",
        "author": "John Doe",
    }
    docJSON, err := json.Marshal(doc)
    if err != nil {
        fmt.Println("Error marshaling the document: ", err)
        return
    }
    addDocRequest := esapi.IndexRequest{
        Index:      indexName,
        DocumentID: docID,
        Body:       bytes.NewReader(docJSON),
    }
    res, err = addDocRequest.Do(context.Background(), es)
    if err != nil {
        fmt.Println("Error adding the document: ", err)
        return
    }
    defer res.Body.Close()

    // 获取文档
    getDocRequest := esapi.GetRequest{
        Index:      indexName,
        DocumentID: docID,
    }
    res, err = getDocRequest.Do(context.Background(), es)
    if err != nil {
        fmt.Println("Error getting the document: ", err)
        return
    }
    defer res.Body.Close()

    var docResult map[string]interface{}
    err = json.NewDecoder(res.Body).Decode(&docResult)
    if err != nil {
        fmt.Println("Error decoding the document result: ", err)
        return
    }

    fmt.Println("Document found: ", docResult)
}

