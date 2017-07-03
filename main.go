package main

import (
	"encoding/json"
	"github.com/labstack/gommon/log"
	"net/http"
	"time"
)

var resJson *res
var timeChan chan int64

func init() {
	resJson = &res{}
	timeChan = make(chan int64, 1)
}

type (
	res struct {
		Timestamp int64 `json:"timestamp"`
	}
)

func getTime(w http.ResponseWriter, req *http.Request) {
	resByte, err := json.Marshal(resJson)
	if err != nil {
		panic(err.Error())
	}
	w.Header().Set("content-type", "application/json;charset=utf-8")
	w.Write(resByte)
}

func startTimer() {
	for{
		select {
		case <-timeChan:
			resJson.Timestamp = time.Now().Unix
		case <-time.After(time.Duration(1)*time.Millisecond):
			resJson.Timestamp = time.Now().Unix
		}
	}
}

func main() {
	http.HandleFunc("/", getTime)
	timeChan <- time.Now().Unix()
	go startTimer()
	err := http.ListenAndServe(":1325", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

