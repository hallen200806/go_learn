package main

import (
	"net"
	"fmt"
	"strings"
)

// socket的服务端
func main() {

	// 监听
	listener , err := net.Listen("tcp","127.0.0.1:8000")
	if err != nil{
		fmt.Println(err)
		return
	}

	defer listener.Close()
	//阻塞等待用户连接
	for {
		conn , err := listener.Accept()
		if err != nil{
			fmt.Println(err)
			continue
		}

		// 接收用户的请求
		buf := make([]byte,1204)  // 1024大小的缓冲区
		n , err1 := conn.Read(buf)
		if err1 != nil{
			fmt.Println(err1)
			continue
		}

		fmt.Println(string(buf[:n]))
		// 给客户端发送数据
		conn.Write([]byte(strings.ToUpper(string(buf[:n]))))

		defer conn.Close()	// 关闭当前用户连接
	}


}
