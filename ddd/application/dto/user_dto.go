package dto

//请求对象
type SimpleUserReq struct {
	ID int `json:"id"`
}

//响应对象
type SimpleUserInfo struct {
	ID       int    `json:"id"`
	UserId   string `json:"user_id"`
	NickName string `json:"nick_name"`
	Money    int    `json:"money"`
}

type UserLogf struct {
	ID   int    `json:"id"`
	Log  string `json:"log"`
	Date string `json:"update_at"`
}

type UserInfo struct {
	ID       int    `json:"id"`
	UserId   string `json:"user_id"`
	RoleId   string `json:"role_id"`
	ServerId string `json:"server_id"`
	Num      int    `json:"num"`
	UseNum   int    `json:"use_num"`
}
