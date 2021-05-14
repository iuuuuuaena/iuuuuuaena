package util

import (
	"encoding/json"
	entity "github.com/iuuuuuaena/entity"
)

// 1. 爬取

// 2. 新建 README.md

// 3. 写入

func Unmarshal(data string) *entity.ResObj {
	// 解析 JSON 数据使用 json.Unmarshal([]byte(JSON_DATA),JSON对应的结构体) ,也就是说我们在解析 JSON 的时候需要确定 JSON 的数据结构
	obj := &entity.ResObj{}
	err := json.Unmarshal([]byte(data), &obj)
	DropErr(err)
	return obj
}

// 创建一个错误处理函数，避免过多的 if err != nil{} 出现
func DropErr(e error) {
	if e != nil {
		panic(e)
	}
}
