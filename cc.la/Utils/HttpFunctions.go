package Utils

import (
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

func SetProxyHeader(w http.ResponseWriter, req *http.Request, url string)  {
	getHeader := w.Header()//获取响应头
	newreq,_ := http.NewRequest(req.Method,url,req.Body)
	CloneHeader(req.Header, &newreq.Header)//请求头设置到代理请求头里面
	newreq.Header.Add("x-forwarded-for",req.RemoteAddr)

	//newresp,_ := http.DefaultClient.Do(newreq)
	dt := &http.Transport{

		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}
	newresp,_ := dt.RoundTrip(newreq)

	CloneHeader(newresp.Header, &getHeader)//响应头设置到代理响应头里面
	w.WriteHeader(newresp.StatusCode)
	defer newresp.Body.Close()
	res,_ := ioutil.ReadAll(newresp.Body)
	w.Write(res)
}
