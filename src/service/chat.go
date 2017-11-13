package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const appUrl = "http://www.tuling123.com/openapi/api"
const apiKey = "9eb7bb3807414cdca1fc8216b7976b10"

//发送请求参数
type Msg struct {
	Key    string `json:"key"`
	Info   string `json:"info"`
	Loc    string `json:"loc"`
	Userid string `json:"userid"`
}

//请求回传结果
type Result struct {
	Code int    `json:"code"`
	Text string `json:"text"`
	Url  string `json:"url"`
}

func main() {
	//chat("你是哪个星座")
	Chat("你认识赵益吗")
	time.Sleep(time.Second * 10)
}
func Chat(info string) *Result {
	fmt.Println("问：", info)
	result, err := askRobot(info)
	if err != nil {
		fmt.Println("error:", err)
	}
	return result
}
func askRobot(info string) (*Result, error) {
	msgJson, err := json.Marshal(Msg{apiKey, info, "沈阳市", "123456"})
	if err != nil {
		fmt.Println("json marsha1 error:", err)
	}
	req, err := http.NewRequest("POST", appUrl, bytes.NewBuffer(msgJson))
	req.Header.Set("Content-Type", "application/json")
	//设置请求代理
	url, _ := url.Parse("http://wang.zhuang:Freudia.1@proxy.neusoft.com:8080")
	proxy := http.ProxyURL(url)
	//创建client（带代理的）
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: proxy,
		},
	}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.Status == "200 OK" {
		body, _ := ioutil.ReadAll(resp.Body)
		result := &Result{}
		err := json.Unmarshal(body, result)
		if err != nil {
			fmt.Println("json marsha1 error:", err)
		}
		return result, nil
	} else {
		return &Result{}, errors.New("fail")
	}
}
