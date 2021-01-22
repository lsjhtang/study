package cores

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Goft struct {
	*gin.Engine
	g *gin.RouterGroup
}

func Ignite() *Goft { //这就是所谓的构造函数，ignite有 发射、燃烧， 很有激情。符合我们骚动的心情
	g := &Goft{Engine: gin.New()}
	return g
}
func (this *Goft) Launch() { //最终启动函数， 不用run，没有逼格
	//config:=InitConfig()
	var port int32 = 8080
	if config := this.beanFactory.GetBean(new(SysConfig)); config != nil {
		port = config.(*SysConfig).Server.Port
	}

	this.Run(fmt.Sprintf(":%d", port))
}
func (this *Goft) Handle(httpMethod, relativePath string, handler interface{}) *Goft {
	if h := Convert(handler); h != nil {
		this.g.Handle(httpMethod, relativePath, h)
	}
	return this // 大功告成
}

//我们弄个方法名叫做Attach ，代表加入中间件。 参数 么。。。。
func (this *Goft) Attach(f Fairing) *Goft {
	this.Use(func(context *gin.Context) {
		err := f.OnRequest(context)
		if err != nil {
			context.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		} else {
			context.Next()
		}
	})
	return this
}

//设定数据库连接对象
func (this *Goft) Beans(beans ...interface{}) *Goft {
	this.beanFactory.setBean(beans...)
	return this
}

// 加一个group 参数
func (this *Goft) Mount(group string, classes ...IClass) *Goft { // 这是挂载， 后面还需要加功能。
	this.g = this.Group(group)
	for _, class := range classes {
		class.Build(this) //这一步是关键 。 这样在main里面 就不需要 调用了
		this.beanFactory.inject(class)
	}
	return this
}

//定时任务
func (this *Goft) Task(spec string, cmd func()) *Goft {
	_, err := getCronTask().AddFunc(spec, cmd)
	if err != nil {
		Error(err)
	}
	return this
}
