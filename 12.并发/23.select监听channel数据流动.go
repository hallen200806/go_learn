package main

func main() {
	/*
	select:
		监听channel上的数据流动
		和switch语法类似，由select开始一个新的选择块，每个选择块由case语句来描述

	限制：
		每个case语句里必须是一个IO操作

	结构：
		select {
		case <-chan1:
			// 如果channel成功读取到时数据，则进行该case处理语句
		case chan2<- 1:
			// 如果channe2成功写入数据，则进行该case处理语句
		default:
			// 如果上面都没有成功，则进行default处理流程

		}

	 */



}
