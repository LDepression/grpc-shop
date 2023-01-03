package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic/v7"
	"log"
	"os"
)

const goodsMapping = `{
	"mappings" : {
		"properties" : {
			"name" : {
				"type" : "text",
				"analyzer":"ik_max_word"
			},
			"id" : {
				"type" : "integer"
			}
		}
	}
}`

type Account struct {
	AccountNumber int32  `json:"account_number"`
	FirstName     string `json:"firstname"`
}

func main() {
	//初始化一个连接
	//如果后面设置为true的话,那么就会是内网地址或者是docker的ip地址,但是我们应该用的是虚拟机的地址
	//导致服务连接不上
	logger := log.New(os.Stdout, "mxshop", log.LstdFlags)
	client, err := elastic.NewClient(elastic.SetURL("http://192.168.28.100:9200"), elastic.SetSniff(false),
		elastic.SetTraceLog(logger))
	if err != nil {
		panic(err)
	}

	q := elastic.NewMatchQuery("address", "street")
	result, err := client.Search().Index("user").Query(q).Do(context.Background())
	if err != nil {
		panic(err)
	}
	total := result.Hits.TotalHits.Value
	fmt.Println("查询到的总数为:", total)
	for _, value := range result.Hits.Hits {
		account := Account{}
		err := json.Unmarshal(value.Source, &account)
		if err != nil {
			return
		}
		//if jsonData, err := value.Source.MarshalJSON(); err == nil {
		//	fmt.Println(string(jsonData))
		//} else {
		//	panic(err)
		//}
		fmt.Println(account)
	}
	//account := Account{AccountNumber: 11245, FirstName: "imooc bobby"}
	//put1, err := client.Index().
	//	Index("myuser").
	//	BodyJson(account).
	//	Do(context.Background())
	//if err != nil {
	//	// Handle error
	//	panic(err)
	//}
	//fmt.Printf("Indexed tweet %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)

	createIndex, err := client.CreateIndex("mygoods").BodyString(goodsMapping).Do(context.Background())
	if err != nil {
		// Handle error
		panic(err)
	}
	if !createIndex.Acknowledged {
		// Not acknowledged
	}
}
