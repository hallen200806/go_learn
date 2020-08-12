package main

import "fmt"

func test02(i int)  {
	var arr[3]int
	arr[i] = 999
	//panic("runtime error: index out of range") // 这里手动触发panic错误
	fmt.Println(arr)

}
func main() {
	// panic函数是程序发生了致命错误，使程序中断
	//panic("hello")
	// panic: runtime error: index out of range [3] with length 3
	test02(3) // arr长度是3，这里传个3就会报错：角标越界

	
}
