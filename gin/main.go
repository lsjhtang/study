package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
	. "topic/src"
	"topic/src/App"
	"topic/src/App/Services"
)

type Product struct {
	gorm.Model
	Code string
	Price uint
}

func main2()  {
	conn := RedisDefaultPool.Get()
	fmt.Println(redis.String(conn.Do("get","abc")))
}

func main() {
	/*defer DB.Close()
	// Migrate the schema
	//db.AutoMigrate(&Product{})
	DB.AutoMigrate(&TopicQuery{})

	// Create
	DB.Create(&TopicQuery{1,"张三",1,10})

	// Read
	var product []Product
	DB.First(&product) // find product with id 1
	DB.Find(&product, "code = ?", "L1212") // find product with code l1212

	// Update - update product's price to 2000
	DB.Model(&product).Update("Price", 2000)

	// Delete - delete product
	DB.Delete(&product)*/

	r := gin.Default()
	v1 := r.Group("/v1/prods")
	BookService := &Services.BookService{}
	CreteBookListResponse := Services.CreteBookListResponse()
	bookListHandler := App.RegisterHandler(
		Services.BookListEndPoint(BookService),
		Services.CreteBookListRequest(),
		CreteBookListResponse,
		)
	v1.GET("", bookListHandler)

	bookDetailHandler := App.RegisterHandler(
		Services.BookDetailEndPoint(BookService),
		Services.CreteBookDetailRequest(),
		CreteBookListResponse,
	)
	v1.GET(":id", bookDetailHandler)


	v1.Use(MustLogin())
	{
		v1.POST("", NewTopic)
		v1.DELETE(":topic_id", DelTopic)
	}


	r.GET("/ping/:abc", func(c *gin.Context) {
		c.JSON(200, gin.H{"hello": len(c.Params)})
	})
	r.Run("0.0.0.0:8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}