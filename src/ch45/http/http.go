package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		str := "hello world"
		rw.Write([]byte(str))
	})
	http.HandleFunc("/time", func(rw http.ResponseWriter, req *http.Request) {
		t := time.Now()
		timeStr := fmt.Sprintf("time: %s", t)
		rw.Write([]byte(timeStr))
	})

	http.ListenAndServe(":8888", nil)
}
