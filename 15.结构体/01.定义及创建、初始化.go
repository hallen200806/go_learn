package main

import "fmt"

type Dog struct {
	Id int
	Name string
	Age int
}



func test_struct()  {
	// 初始化1
	var dog Dog
	dog.Id = 1
	dog.Name = "ali"
	dog.Age = 2
	fmt.Println(dog)

	// 初始化2
	dog2 := Dog{Id:2,Name:"tom",Age:3}
	fmt.Println(dog2)

	// 初始化3
	dog3 := new(Dog)
	dog3.Id = 3
	dog3.Name = "hmm"
	dog3.Age = 4
	fmt.Println(*dog3)   // 指针

	
}
func main() {
	/*

	 */

	 test_struct()
	 
}
