package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

// 发送GET请求
// url：         请求地址
// response：    请求返回的内容
func Get(url string) string {

	// 超时时间：5秒
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	var buffer [512]byte
	result := bytes.NewBuffer(nil)
	for {
		n, err := resp.Body.Read(buffer[0:])
		result.Write(buffer[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}

	return result.String()
}

func Get2(url string) []byte{
	resp, err := http.Get(url)
	DropErr(err)
	defer resp.Body.Close()
	if err != nil {
		panic(err)
	}
	all, err := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(all))
	return all
}

func Get3(url string) {
	// 1. 建立 请求客户端，就可以携带请求头header 和 请求体了
	client := &http.Client{}
	// 2. 创建 Get， 请求， 第三个参数 body 为 nil
	request, err := http.NewRequest("Get", url, nil)
	if err != nil {
		panic(err)
	}
	// 3. 添加请求头
	if request != nil {
		request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.212 Safari/537.36")
		request.Header.Add("accept", "application/json, text/javascript, */*; q=0.01")
	} else {
		fmt.Println("request is nil")
	}
	// 4. 正式发送请求
	response, err := client.Do(request)
	// 5. 别忘了关
	defer response.Body.Close()
	if err != nil {
		panic(err)
	}
	all, err := ioutil.ReadAll(response.Body)
	fmt.Println(string(all))
}

// 发送POST请求
// url：         请求地址
// data：        POST请求提交的数据
// contentType： 请求体格式，如：application/json
// content：     请求放回的内容
func Post(url string, data interface{}, contentType string) string {

	// 超时时间：5秒
	client := &http.Client{Timeout: 5 * time.Second}
	jsonStr, _ := json.Marshal(data)
	resp, err := client.Post(url, contentType, bytes.NewBuffer(jsonStr))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	result, _ := ioutil.ReadAll(resp.Body)
	return string(result)
}
