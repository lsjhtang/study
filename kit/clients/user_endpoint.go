package clients

type Request struct {
	Uid int `json:"uid"`
}

type Response struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   map[string]interface{}
}
