package main

import "fmt"

func test_map(mapa map[string]string) (map[string]string) {
	fmt.Println(mapa)
	mapa["gender"] = "男"
	return mapa
}
func main() {
	/*
	map参数：地址或者引用的传递
		形参的改变会改变实参
	 */
	mapa := make(map[string]string)
	mapa["name"] = "张三"
	fmt.Println(mapa)
	mapb := test_map(mapa)
	fmt.Println(mapb)
	fmt.Println(mapa)
}
