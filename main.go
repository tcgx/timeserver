package main

import (
	"encoding/json"
	"github.com/labstack/gommon/log"
	"net/http"
	"time"
)

var resJson *res

func init() {
	resJson = &res{}
}

type (
	res struct {
		Timestamp int64 `json:"timestamp"`
	}
)

func getTime(w http.ResponseWriter, req *http.Request) {

	resJson.Timestamp = time.Now().UnixNano()/1000000
	resByte, err := json.Marshal(resJson)
	if err != nil {
		panic(err.Error())
	}
	w.Header().Set("content-type", "application/json;charset=utf-8")
	w.Write(resByte)
}

func main() {
	http.HandleFunc("/", getTime)
	err := http.ListenAndServe(":1323", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
