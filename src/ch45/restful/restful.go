package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Student struct {
	Name string `json: "name"`
	Age  int    `json: "age"`
}

var Students map[string]*Student

func init() {
	Students = map[string]*Student{}
	Students["Mike"] = &Student{"Mike", 10}
	Students["Job"] = &Student{"Job", 12}
}

func Index(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	w.Write([]byte("hello"))
}

func GetInfoByName(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	name := p.ByName("name")
	var (
		ok       bool
		info     *Student
		err      error
		infoJson []byte
	)

	if info, ok = Students[name]; !ok {
		w.Write([]byte("{error: \"not found\"}"))
		return
	}
	if infoJson, err = json.Marshal(info); err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err)))
		return
	}
	w.Write([]byte(infoJson))
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/student/:name", GetInfoByName)
	http.ListenAndServe(":8888", router)
}
