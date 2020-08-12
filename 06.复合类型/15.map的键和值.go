package main

import "fmt"

func main() {
	/*
	map的key：
		类型必须支持== 或者！= ，一般建议写基本类型
		切片，函数，切片的结构类型，不能作为map的key
	map的value：
		可以是某一种类型
			mapa := make(map[string]string)
			mapa["name"] = "张三"
		也可以是数组
			mapb := make(map[string][3]int)
			mapb["name"] = [3]int{1,2,3}
	取值：
		for循环取出所有的key和value：
			for k,v := range mapb{
				fmt.Println(k)
				fmt.Println(v)
			}
		取出单独某个key：
			k, v := mapb["张三"]

		注意：如果有改key，咋v为true，否则为false，可以通过这种方式进行判断

	删除：
		delete(mapb,"张三")		前面是哪个map，后面是要删除的key
		注意：如果key不存在，不会报错
	 */
	mapb := make(map[string][3]int)
	mapb["张三"] = [3]int{1,2,3}
	mapb["李四"] = [3]int{4,5,6}
	fmt.Println(mapb)
	// 循环取出
	for k,v := range mapb{
		fmt.Println(k)
		fmt.Println(v)
	}
	// 取出某个key对应的值
	k, v := mapb["张三"]
	fmt.Println(k)
	fmt.Println(v)  // 如果有改key，咋v为true，否则为false

	// 删除某个key
	delete(mapb,"张三")
	fmt.Println(mapb)  // 张三被删除了
}
