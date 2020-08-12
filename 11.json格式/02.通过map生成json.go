package main

import (
	"fmt"
	"encoding/json"
)

func main() {

	// 创建map
	m := make(map[string]interface{},4)
	m["company"] = "网易"
	m["subject"] = []string{"go","python","java"}
	m["price"] = 360.36
	m["isok"] = true
	//fmt.Println(m)

	//json_ret, err := json.Marshal(m)
	// 格式化json
	json_ret,err := json.MarshalIndent(m,""," ")
	if err != nil{
		return
	}
	fmt.Println(string(json_ret))     // 是切片得转string
}
