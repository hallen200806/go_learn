package main

func make_fibonacci(ch chan<- int,quit <-chan bool)  {
	
}
func main() {

	ch := make(chan int)

	quit := make(chan bool)

	// 生产者，产生数字
	make_fibonacci(ch, quit)
	
	// 消费者读取内容


}
