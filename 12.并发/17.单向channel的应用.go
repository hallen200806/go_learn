package main

import "fmt"

func producer(out chan<- int)  { // 这里只用到了写
	for i:=0;i<10;i++{
		out <- i
	}
	close(out)
}

func consumer(in <-chan int)  {  // 这里只用到了读

	for num := range in{
		fmt.Println(num)
	}

}
func main() {

	ch := make(chan int)
	// 生产者：生产数字，写入channel
	go producer(ch)		// 地址传递

	// 消费者：从channel读取数字打印
	consumer(ch)
}
