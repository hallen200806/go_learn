package main

import "fmt"

type hero1 struct {
	name string
	age int
	power int
}

func testmap(mapa map[int]hero1)(map[int]hero1)  {
	fmt.Println(mapa)
	mapa[101] = hero1{name:"张三2222",age:29,power:111}
	return mapa
}
func main() {
	// 参数作为地址传递，形参改变了实参也会变
	mapa := make(map[int]hero1)
	mapa[101] = hero1{name:"张三",age:18,power:97}
	mapa[102] = hero1{name:"张三",age:18,power:97}
	mapa[103] = hero1{name:"张三",age:18,power:97}
	fmt.Println(mapa)
	mapb := testmap(mapa)
	fmt.Println(mapb)
	fmt.Println(mapa)
}
