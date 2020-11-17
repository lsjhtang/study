package src

import (
	"github.com/gin-gonic/gin"
	"net/http"
	. "topic/src/AppInit"
)

func MustLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, ok:=c.GetQuery("token"); !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"data": "token验证不通过"})
			c.Abort()
		}
		c.Next()
	}
}

func GetTopicDetail(c *gin.Context) {
	//c.JSON(200, gin.H{"data": "获取topic_id="+c.Param("topic_id")+"的帖子"})
	//c.JSON(200, CreateTopic(100, "获取帖子成功"))
	TopicQuery := &TopicQuery{}
	DB.Find(TopicQuery, c.Param("topic_id"))
	c.JSON(200, TopicQuery)
	/*err := c.BindQuery(TopicQuery)
	if err != nil {
		c.JSON(400, gin.H{"msg":err.Error()})
	}else {
		c.JSON(200, TopicQuery)
	}*/
}

func NewTopic(c *gin.Context)  {
	Topic := &Topics{}
	err := c.BindJSON(Topic)
	if err != nil {
		c.JSON(400, gin.H{"msg":err.Error()})
	}else {
		c.JSON(200, Topic)
	}
	//c.JSON(200, gin.H{"data": map[string]string{"status":"新增帖子成功"}})
}

func DelTopic(c *gin.Context)  {
	c.JSON(200, gin.H{"data": map[string]string{"status":"删除帖子成功"}})
}

/*func GetBookList(c *gin.Context)  {
	products := &BookList{}
	DB.Limit(10).Order("book_id desc").Find(products)
	c.JSON(200, products)
}*/
