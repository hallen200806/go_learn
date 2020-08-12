package main

import (
	"encoding/json"
	"fmt"
)

// 成员变量必须大写
type it struct {
	Company string `json:"company"`  // 二次编码
	Subject []string `json:"-"`  // 次字段不会输出到屏幕
	IsOk bool	`json:",string"`  // 转换成字符串后输出
	Price float64	`json:"price"`
}

func main() {
	/*
	通过结构体生成json：
		1、成员变量必须大写
		2、通过json.Marshal(s)转换：返回两个参数，一个是文本的切片类型，一个是bool类型的err
			buf,err := json.Marshal(s)  // 返回的buf是切片
		3、可以在结构体中重命名key,将首字母大写转为小写：这里的引号是tab键上面的键
			Company string `json:"company"`
			Subject []string `json:"subject"`
			IsOk bool	`json:"is_ok"`
			Price float64	`json:"price"`
	 */

	 // 定义一个结构体变量，同时初始化
	 s := it{"baidu",[]string{"go","python","java"},true,666}

	 // 编码，根据内容生成json文本
	 //buf,err := json.Marshal(s)  // 返回的buf是切片

	 // 格式化json文本
	json_ret,err := json.MarshalIndent(s,""," ")
	 if err != nil{
	 	return
	 }
	 fmt.Println(string(json_ret))   // buf是个切片，需要转一下
}
