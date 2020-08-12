package main

func main() {

	/*
	gorountine运行在相同的地址空间，因此访问共享内存必须做好同步，goroutine奉行通过通信来共享内容，而不是共享内存来通信

	引用类型channel是CSP模式的具体实现，用于多个goroutine通讯，其内部实现了同步，确保并发安全

	channel类型：
		和map类似，channel也是一个对应make创建的底层数据结构的引用

		当我们复制一个channel或用于函数参数传递时，我们只是拷贝了一个channel引用，因此调用者将被引用同一个channel对象
		和其他的引用类型一样，channel的0值也是nil

	创建通道：
		chan 类型的空值是 nil，声明后需要配合 make 后才能使用。

	定义channel：
		定义一个channel时，也需要定义发送到channel的值的类型。channel可以使用make()函数来创建
			make(chan Type)  // 等价于make(chan Type,0)
			make(chan Type, capacity)
				当capacity=0时，channel是无缓冲阻塞读写的，
				当capacity>0时，channel有缓冲，是非阻塞的，直到写满capacity个元素才阻塞写入
			注意：
				必须是chan名，不能变

		// 声明不带缓冲的通道
		ch1 := make(chan string)

		// 声明带10个缓冲的通道
		ch2 := make(chan string, 10)

		// 声明只读通道
		ch3 := make(<-chan string)

		// 声明只写通道
		ch4 := make(chan<- string)

	注意：

		不带缓冲的通道，进和出都会阻塞。

		带缓冲的通道，进一次长度 +1，出一次长度 -1，如果长度等于缓冲长度时，再进就会阻塞。


	通道 ch 是可以进行遍历的，遍历的结果就是接收到的数据：
		for data := range ch {
		}

	有缓冲和无缓冲的区别：
		无缓冲的  就是一个送信人去你家门口送信 ，你不在家 他不走，你一定要接下信，他才会走。
			无缓冲保证信能到你手上

		有缓冲的 就是一个送信人去你家仍到你家的信箱 转身就走 ，除非你的信箱满了 他必须等信箱空下来。
			有缓冲的 保证 信能进你家的邮箱
	 */
}
