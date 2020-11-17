package main

import "net/http"

type ProxyHandler struct {}

func (*ProxyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request)  {

}

func main() {

}