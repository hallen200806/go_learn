package main

import (
	"encoding/json"
	"fmt"
)

type IT struct {
	Company string `json:"company"`  // 二次编码
	Subject []string `json:"subject"`  // 次字段不会输出到屏幕
	IsOk bool	`json:"isok"`  // 转换成字符串后输出
	Price float64	`json:"price"`
}


func main() {

	/*
	json转为结构体：
		err := json.Unmarshal([]byte(json_context),&type_tmp)

		注意：这里只有一个参数，必须转成byte类型的切片，必须使用&

	如果只需要json中的一部分转为结构体，只需要在结构体中定义需要的即可
	 */

	json_context := `
	{
	 "company": "网易",
	 "isok": true,
	 "price": 360.36,
	 "subject": [
	  "go",
	  "python",
	  "java"
	 ]
	}

	`
	var type_tmp IT
	err := json.Unmarshal([]byte(json_context),&type_tmp)   // 注意，这里只有一个参数，，必须转成byte类型的切片，必须使用&
	if err != nil{
		return
	}
	fmt.Println(type_tmp)

}
