package main

import (
	conf "./conf"
	"encoding/json"
	"net/http"
	"strings"
)

func main() {
	conf.Pool = conf.NewPool("127.0.0.1:6379", "", 0)
	http.HandleFunc("/save/", saveHandler)
	http.HandleFunc("/get/", getHandler)
	http.ListenAndServe(":1234", nil)
}
func saveHandler(w http.ResponseWriter, r *http.Request) {
	println(r.URL.Path)
	path := r.URL.Path
	paths := strings.Split(path, "/")
	if len(paths) != 4 {
		result, _ := json.Marshal(string("url错误"))
		w.Write(result)
		return
	}
	userid := paths[2]
	content := paths[3]
	setResult := conf.SetRedis(userid, content)
	result, _ := json.Marshal(setResult)
	w.Write(result)
}
func getHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	paths := strings.Split(path, "/")
	if len(paths) != 3 {
		result, _ := json.Marshal(string("url错误"))
		w.Write(result)
		return
	}
	userid := paths[2]
	getResult := conf.GetRedis(userid)
	result, _ := json.Marshal(getResult)
	w.Write(result)
}
