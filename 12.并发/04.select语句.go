package main

import (
	"fmt"
	"time"
)

func timeout() (out chan int){
	fmt.Println("timeout")
	return
}

func main() {

	for{
		select {
		case num := <- ch:
			fmt.Println("num:%s",num)
		case <- timeout:
			fmt.Println("timeout")
			
		}
	}
}
