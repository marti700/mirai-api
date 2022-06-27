package main

import (
	"net/http"
	"mirai-api/reqhandler"
)

func main (){
	http.HandleFunc("/regression", reqhandler.HandleRegression)
	http.HandleFunc("/decisiontree/regression", reqhandler.HandleDecisionTreeRegressor)
	http.HandleFunc("/decisiontree/classification", reqhandler.HandleDecisionTreeClassifier)
	http.ListenAndServe(":9090", nil)
}