package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pquerna/ffjson/ffjson"
	"sync"
)

type NewsModel struct {
	Id int
	Title string
}

func (news NewsModel) ToJson () string {
	result, err :=ffjson.Marshal(news)
	if err != nil {
		return err.Error()
	} else {
		return string(result)
	}
}


func main() {
	/*news:=NewsModel{1,"test"}
	fmt.Println(news.ToJson())*/
	/*fmt.Printf(services.GetNews()+services.GetUser())*/
	/*aa:="11"
	var u *string
	u = new(string)
	test(&aa)
	fmt.Println(*u)*/
	/*var service services.IService=services.ServiceFactory{}.Create("news")
	fmt.Printf(service.Get(1241))*/

	/*db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/swoole")
	if err !=nil {
		fmt.Printf(err.Error())
		return
	}

	rows, err := db.Query("select user_name,age from test")
	if err !=nil {
		fmt.Printf(err.Error())
		return
	}
	columns, _ :=rows.Columns()
	allRows:=make([]interface{},0)
	filedMap:=make(map[string]string)
	//fmt.Println(allRows,columns)
	for rows.Next()	{
		oenRow :=make([]interface{},len(columns))
		scaRow :=make([]interface{},len(columns))
		for i,_ :=range oenRow{
			scaRow[i] = &oenRow[i]
		}
		rows.Scan(scaRow...)
		for i,v:=range oenRow{
			v,ok:=v.([]byte)
			if ok {
				filedMap[columns[i]] = string(v)
			}
		}
		allRows = append(allRows,filedMap)
	}
	fmt.Println(allRows)*/

	/*user := strings.Split("aa,bb,cc,dd",",")
	age := strings.Split("11,22,33,44",",")
	//fmt.Println(strings.Split(user,","))
	c1, c2 := make(chan bool), make(chan bool)
	ret := make([]string,0)
	go func() {
		for _, v := range user{
			<- c1
			ret = append(ret,v)
			c2 <- true
		}
	}()
	go func() {
		for _, v := range age{
			<- c2
			ret = append(ret,v)
			c1 <- true
		}

	}()
	c1 <- true
	time.Sleep(1)
	fmt.Println(ret)*/
	/*c := make(chan map[int][]byte)
	url := "http://www.jtthink.com/course/%d"
	for i:=1;i<=3;i++ {
		go func(index int) {
			ret ,_ := http.Get(fmt.Sprintf(url,index))
			cnt ,_ := ioutil.ReadAll(ret.Body)
			c<- map[int][]byte{index:cnt}
			//ioutil.WriteFile(fmt.Sprintf("./files/%d",index), cnt,777)
		}(i)
	}

	for k := range c{
		for index,cnt := range k {
			ioutil.WriteFile(fmt.Sprintf("./files/%d.html",index), cnt,777)
		}
	}*/
	/*defer func() {
		err := recover()
		fmt.Println(err)
	}()
	fmt.Println(test())*/
	/*syn := sync.WaitGroup{}
	url := "http://www.jtthink.com/course/%d"
	for i:=1;i<=10;i++ {
		go func(index int) {
			defer func() {
				syn.Done()
			}()
			ret ,_ := http.Get(fmt.Sprintf(url,index))
			cnt ,_ := ioutil.ReadAll(ret.Body)
			ioutil.WriteFile(fmt.Sprintf("./files/%d",index), cnt,777)
			fmt.Println("启动协程",index)
		}(i)
		syn.Add(1)
	}*/
	syn := sync.WaitGroup{}
	mutex := sync.Mutex{}
	list := make([]int,0)
	go func() {
		defer func() {
			syn.Done()
			mutex.Unlock()
		}()
		mutex.Lock()
		for i:=0; i<1000000;i++ {
			list = append(list,i)
		}
		fmt.Println(len(list))
	}()
	go func() {
		defer func() {
			syn.Done()
			mutex.Unlock()
		}()
		mutex.Lock()
		for i:=0; i<1000000;i++ {
			list = append(list,i)
		}
		fmt.Println(len(list))
	}()

	//fmt.Println("任务总数",runtime.NumGoroutine())
	syn.Add(2)
	syn.Wait()
	fmt.Println(len(list))
}

func test() (ret int) {
	a := 1
	defer func() {
		ret = 10
	}()
	panic("111")
	return a
}