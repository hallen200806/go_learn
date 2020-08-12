package main

import "fmt"

func test(a,b int) (result int,err error) { // 这里使用了error接口
	err = nil
	if b == 0 {
		fmt.Println(err)
	}else {
		result = a / b
	}
	return
}
func main() {
	//ret := test(10,20)
	ret,err := test(10,0)  //被除数不能为0,报error
	if err != nil{
		fmt.Println(err)
	}else {
		fmt.Println(ret)
	}


}
