package main

import (
	"net"
	"fmt"
)

// socket的客户端
func main() {
	// 主动连接服务端
	conn , err := net.Dial("tcp","127.0.0.1:8000")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer conn.Close()

	// 发送数据
	conn.Write([]byte("hello hallen!"))

}
