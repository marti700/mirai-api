package main

import (
	"net/http"
	"mirai-api/reqhandler"
)

func main (){
	http.HandleFunc("/regression", reqhandler.HandleRegression)
	http.ListenAndServe(":9090", nil)
}