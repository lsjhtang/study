package main

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"log"
)

func main()  {
	client, err := elastic.NewClient(elastic.SetURL("http://192.168.87.128:9200"), elastic.SetSniff(false))
	if err != nil {
		log.Print(err)
	}
	ctx := context.Background()
	/*mapping, abc := client.GetMapping().Index("news").Do(ctx)
	if abc != nil {
		log.Print(abc)
	}
	fmt.Println(mapping)*/
	json := `{"news_title":"test1","news_type":"go","news_status":1}`
	mapping, abc := client.Index().Index("news").Id("100").BodyString(json).Do(ctx)
	if abc != nil {
		log.Print(abc)
	}
	fmt.Println(mapping)
}


