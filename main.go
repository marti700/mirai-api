package main

import (
	"mirai-api/reqhandler"
	"net/http"
)

func main() {
	http.HandleFunc("/regression", reqhandler.HandleRegression)
	http.HandleFunc("/decisiontree/regression", reqhandler.HandleDecisionTreeRegressor)
	http.HandleFunc("/decisiontree/classification", reqhandler.HandleDecisionTreeClassifier)
	http.HandleFunc("/train", reqhandler.Handle)
	http.ListenAndServe(":9090", nil)
}
