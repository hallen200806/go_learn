package main

import "fmt"

func main() {
	// 比如int64，我想改成bigint
	type bigint int64
	var a bigint
	a = 10
	fmt.Println(a)
	fmt.Printf("%T\n",a)  // main.bigint main包.bigint
}
