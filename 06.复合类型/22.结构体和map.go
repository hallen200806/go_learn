package main

import "fmt"

type hero struct {
	name string
	age int
	power int
}
func main() {
	mapa := make(map[int]hero)
	mapa[101] = hero{name:"张三1",age:181,power:971}
	mapa[102] = hero{name:"张三2",age:182,power:972}
	mapa[103] = hero{name:"张三3",age:183,power:973}
	fmt.Println(mapa)

	// 删除
	delete(mapa,101)
	fmt.Println(mapa)

	// 使用切片
	// 定义
	mapb := make(map[int][]hero)
	// 添加
	mapb[103] = []hero{
		{name:"张三11",age:11,power:80},
		{name:"李四11",age:19,power:67},
	}
	// 添加
	mapb[104] = append(mapb[104],hero{name:"张三222",age:23,power:56})

	fmt.Println(mapb)
}
