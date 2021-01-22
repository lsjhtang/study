package main

import (
	. "activity/config"
	Models "activity/models"
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	BASE = iota + 1
	TASK1
	TASK2
	TASK3
	TASK4
)
const CACHEPREFIX = "202012_platact_7_"

type Property struct {
	GameId    string
	ChannelId map[string]bool
	aa        int
}

func NewProperty() *Property {
	return &Property{
		GameId: "670",
		ChannelId: map[string]bool{
			"8999": true, "8981": true, "8980": true, "8910": true, "8228": true, "8003": true, "7878": true, "7877": true, "7876": true, "6994": true, "6268": true, "6168": true,
			"5924": true, "5832": true, "5017": true, "4928": true, "4527": true, "4526": true, "4525": true, "4524": true, "4523": true, "4522": true, "4515": true, "4285": true,
			"3129": true, "2965": true, "2964": true, "2963": true, "2962": true, "2961": true, "2872": true, "2753": true, "2663": true, "2446": true, "2445": true,
		},
	}
}

func main() {
	r := gin.Default()
	act := r.Group("/1609135992")
	act.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
	})
	act.Use(gin.Logger(), gin.Recovery())
	{
		act.GET("/getUserInfo", getUserInfo)
		act.GET("/task1", task1)
		act.GET("/task2", task2)
		act.GET("/task3", task3)
	}
	r.Run("hd.9187.cn:80") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func getUserInfo(c *gin.Context) {
	checkAll(c, BASE)
	user, err := getUser(c)
	if err != nil {
		c.JSON(http.StatusOK, newResponse(http.StatusBadRequest, "用户信息不存在", DataMap{}))
		return
	}
	num, _ := strconv.Atoi(user["Num"].(string))
	data := DataMap{"user": user, "user_task2": num + 1}
	getUserTask(c, &data)
	c.JSON(http.StatusOK, newResponse(http.StatusOK, "初始化成功", data))
}

func task1(c *gin.Context) {
	checkAll(c, TASK1)
}

func task2(c *gin.Context) {
	checkAll(c, 2)
}

func task3(c *gin.Context) {
	checkAll(c, 2)
	c.JSON(http.StatusOK, newResponse(http.StatusOK, "测试", DataMap{}))
}

func getUserTask(c *gin.Context, m *DataMap) {
	uid := c.Keys["uid"].(string)
	cacheKey := getRedisKey(2, uid)
	newTask, _ := redis.StringMap(Redis.Do("HGETALL", cacheKey))
	if len(newTask) == 0 {
		newTask = map[string]string{"1": "1", "2": "1"}
		Redis.Do("hMset", redis.Args{}.Add(cacheKey).AddFlat(newTask)...)
		Redis.Do("EXPIRE", cacheKey, 86400*30)
	}
	(*m)["user_task1"] = newTask

	cacheKey = getRedisKey(3, uid)
	signTask, _ := redis.String(Redis.Do("hGet", cacheKey, uid))
	if signTask == "" {
		signTask = "1"
		Redis.Do("hSet", uid, signTask)
		Redis.Do("EXPIRE", cacheKey, 86400*30)
	}
	(*m)["user_task2_type"] = signTask

	cacheKey = getRedisKey(4, uid)
	levelTask, _ := redis.StringMap(Redis.Do("HGETALL", cacheKey))
	if len(levelTask) == 0 {
		levelTask = map[string]string{"1": "1", "2": "1", "3": "1"}
		Redis.Do("hMset", redis.Args{}.Add(cacheKey).AddFlat(levelTask)...)
		Redis.Do("EXPIRE", cacheKey, 86400*30)
	}
	(*m)["user_task3"] = levelTask
}

func getUser(c *gin.Context) (DataMap, error) {
	uid := c.Keys["uid"].(string)
	usrKey := getRedisKey(1, uid)
	user, err := Redis.Do("HGETALL", usrKey)
	if len(user.([]interface{})) == 0 {
		outUser := Models.NewUsers()
		createUser := Models.Users{
			Openid:    c.DefaultQuery("openid", ""),
			UserId:    c.Keys["uid"].(string),
			NickName:  c.Query("nick_name"),
			UserName:  c.Keys["userName"].(string),
			CreatedAt: int(time.Now().Unix()),
		}
		DB.Table("202012_platact_7_user").Where("user_id = ?", uid).Attrs(createUser).FirstOrCreate(outUser)
		_, err = Redis.Do("HMSET", redis.Args{}.Add(usrKey).AddFlat(outUser)...)
		Redis.Do("EXPIRE", usrKey, 86400*30)
		return StructToMap(outUser), err
	}
	return SliceToMap(user), nil
}

func checkAll(c *gin.Context, t int) {
	uid, err := c.Cookie("plat_uid")
	if err != nil {
		c.JSON(http.StatusOK, newResponse(http.StatusBadRequest, "请先登入1", DataMap{}))
		return
	}

	userName, err := c.Cookie("plat_user_name")
	if err != nil {
		c.JSON(http.StatusOK, newResponse(http.StatusBadRequest, "请先登入2", DataMap{}))
		return
	}
	c.Set("uid", uid)
	c.Set("userName", userName)
	if isWeChatBrowser(c) {
		isOpenid(c)
	}
	switch t {
	case TASK2:
		if c.Query("task_id") == "" {
			c.JSON(http.StatusOK, newResponse(http.StatusBadRequest, "参数错误1", DataMap{}))
			return
		}
		user, _ := getUser(c)
		if !checkBindRole(user) {
			c.JSON(http.StatusOK, newResponse(http.StatusBadRequest, "请先绑定角色", DataMap{}))
			return
		}
	case TASK3:
		if c.Query("task_id") == "" {
			c.JSON(http.StatusOK, newResponse(http.StatusBadRequest, "参数错误1", DataMap{}))
			return
		}
		if c.Query("type") == "" {
			c.JSON(http.StatusOK, newResponse(http.StatusBadRequest, "参数错误2", DataMap{}))
			return
		}

		user, _ := getUser(c)
		if !checkBindRole(user) {
			c.JSON(http.StatusOK, newResponse(http.StatusBadRequest, "请先绑定角色", DataMap{}))
			return
		}
	case TASK4:
		if c.Query("server_id") == "" || c.Query("role_id") == "" {
			c.JSON(http.StatusOK, newResponse(http.StatusBadRequest, "缺少必要参数", DataMap{}))
			return
		}
	}
	isChannel(c)
}

func getRedisKey(t int, id string) string {
	switch t {
	case 1:
		return CACHEPREFIX + "USER_" + id
	case 2:
		return CACHEPREFIX + "NEW_TASK" + id
	case 3:
		return CACHEPREFIX + time.Now().Format("20060102") + "_TASK"
	case 4:
		return CACHEPREFIX + "LEVEL_TASK_" + id
	case 5:
		return CACHEPREFIX + "NEW_TASK_RED"
	case 6:
		return CACHEPREFIX + "_RED_DATE" + id
	case 7:
		return CACHEPREFIX + time.Now().Format("20060102") + "TEMP_" + id
	case 11:
		return CACHEPREFIX + "RED_PKG_SUM"
	case 12:
		return CACHEPREFIX + "RED_PKG"
	default:
		return CACHEPREFIX + "USER_" + id
	}
}

func isOpenid(c *gin.Context) {
	key := "WX|$oPoUoNLOGIN"
	signStr := c.Query("openid") + c.Query("subscribe") + key
	if fmt.Sprintf("%x", md5.Sum([]byte(signStr))) == c.Query("sign") {
		c.JSON(http.StatusOK, newResponse(http.StatusBadRequest, "非法请求", DataMap{}))
		return
	}
}

func checkBindRole(dataMap DataMap) bool {
	if dataMap["role_id"] == "" {
		return false
	}
	return true
}

func isWeChatBrowser(c *gin.Context) bool {
	if ok, _ := regexp.MatchString(".*micromessenger.*", c.Request.Header.Get("User-Agent")); ok {
		return true
	}
	return false
}

func isChannel(c *gin.Context) {
	infos, err := createRolesInfo(fmt.Sprintf("%s", c.Keys["uid"]))
	if err != nil {
		c.JSON(http.StatusOK, newResponse(http.StatusBadRequest, "账号信息不存在", DataMap{}))
		return
	}
	info, _ := infos.(map[string]interface{})
	if info[NewProperty().GameId] == nil {
		c.JSON(http.StatusOK, newResponse(http.StatusBadRequest, "请先创建角色", DataMap{}))
		return
	}

	var roleInfoSort []float64
	roleInfoMap := map[float64]string{}
	for _, v := range info[NewProperty().GameId].([]interface{}) {
		v := v.(map[string]interface{})
		if v["create_time"] == nil {
			roleInfoSort = append(roleInfoSort, v["last_time"].(float64))
			roleInfoMap[v["last_time"].(float64)] = v["role_id"].(string)
		} else {
			roleInfoSort = append(roleInfoSort, v["create_time"].(float64))
			roleInfoMap[v["create_time"].(float64)] = v["role_id"].(string)
		}
	}

	sort.Float64s(roleInfoSort)
	roleChannel, err := getUserChanId(roleInfoMap[roleInfoSort[0]], strconv.FormatFloat(roleInfoSort[0], 'E', -1, 64))
	if err != nil {
		c.JSON(http.StatusOK, newResponse(http.StatusBadRequest, "请点击活动页的 <font color='#dc143c'>下载体验</font> 按钮并注册角色再参与活动", DataMap{}))
		return
	}
	channelId := roleChannel.(map[string]interface{})["channel_id"].(string)
	if !NewProperty().ChannelId[channelId] {
		c.JSON(http.StatusOK, newResponse(http.StatusBadRequest, "请点击活动页的 <font color='#dc143c'>下载体验</font> 按钮并重新 <font color='#dc143c'>注册账号</font> 再参与活动", DataMap{}))
		return
	}
	c.Set("ChannelId", channelId)
}

func getUserChanId(roleId string, createTime string) (interface{}, error) {
	data := map[string]string{
		"role_id":     roleId,
		"create_time": createTime,
		"time":        strconv.FormatInt(time.Now().Unix(), 10),
	}
	getSign(&data)
	data["ct"] = "user"
	data["ac"] = "getUserChanId"

	return getRequest("https://plat.9187.cn/user/", &data)
}

func createRolesInfo(uid string) (interface{}, error) {
	data := map[string]string{
		"user_id": uid,
		"time":    strconv.FormatInt(time.Now().Unix(), 10),
	}
	getSign(&data)
	data["ct"] = "user"
	data["ac"] = "getUserRole"

	return getRequest("https://plat.9187.cn/user/", &data)
}

func getRequest(rawurl string, data *map[string]string) (interface{}, error) {
	params := url.Values{}
	myUrl, _ := url.Parse(rawurl)
	for k, v := range *data {
		params.Set(k, v)
	}
	myUrl.RawQuery = params.Encode()
	res, err := http.Get(myUrl.String())
	if err != nil {
		log.Fatal(err)
	}
	robots, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	rData := &returnData{}
	err = json.Unmarshal(robots, rData)
	if err != nil {
		log.Fatal(err)
	}

	if rData.State == 1 {
		return rData.Data, nil
	} else {
		return nil, errors.New(rData.Msg)
	}
}

func getSign(data *map[string]string) {
	sign := "YBNK7-TVB4L-SD07N-UGBJ8"
	dataSice := make([]string, 0)
	for k, v := range *data {
		dataSice = append(dataSice, k+"="+v)
	}

	sort.Strings(dataSice)

	str := strings.Join(dataSice, "&")
	dataSign := fmt.Sprintf("%x", md5.Sum([]byte(str+sign)))
	b := *data
	b["sign"] = dataSign
	data = &b
}

func StructToMap(obj interface{}) DataMap {
	objType := reflect.TypeOf(obj).Elem()
	objValue := reflect.ValueOf(obj).Elem()

	var data = DataMap{}
	for i := 0; i < objType.NumField(); i++ {
		data[objType.Field(i).Name] = objValue.Field(i).Interface()
	}
	return data
}

func SliceToMap(s interface{}) DataMap {
	slices := s.([]interface{})
	data := DataMap{}
	for i := 0; i < len(slices); i += 2 {
		data[string(slices[(i)].([]byte))] = string(slices[(i + 1)].([]byte))
	}
	return data
}

type DataMap map[string]interface{}
type response struct {
	StatusCode int         `json:"status_code"`
	Msg        string      `json:"msg"`
	Data       interface{} `json:"data"`
}

type returnData struct {
	State int         `json:"state"`
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Data  interface{} `json:"data"`
}

func newResponse(statusCode int, msg string, data DataMap) *response {
	return &response{StatusCode: statusCode, Msg: msg, Data: data}
}
