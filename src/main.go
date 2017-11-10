package main

import (
	conf "./conf"
	"encoding/json"
	"net/http"
	"strings"
)

func main() {
	http.Handle("/js/", http.FileServer(http.Dir("template")))
	http.Handle("/html/", http.FileServer(http.Dir("template")))

	conf.Pool = conf.NewPool("127.0.0.1:6379", "", 0)
	http.HandleFunc("/save/", saveHandler)
	http.HandleFunc("/get/", getHandler)
	http.ListenAndServe(":1234", nil)
}
func saveHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		result, _ := json.Marshal(string("参数错误"))
		w.Write(result)
		return
	}
	userid := r.FormValue("userId")
	content := r.FormValue("userValue")
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
