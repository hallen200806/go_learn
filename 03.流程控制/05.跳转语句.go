package main

import "fmt"

func main() {
	// break ,continue

	for i:=0;i<10;i++{
		fmt.Println(i)
		if i == 1 {
			fmt.Println("等于1,继续下一次循环")
			continue
		}
		if i ==5{
			fmt.Println("等于5，终止循环")
			break
		}

	}
}
