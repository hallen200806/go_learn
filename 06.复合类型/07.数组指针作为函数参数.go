package main

import "fmt"

func arr_test(a *[5]int)  {
	fmt.Println(*a)
}
func main() {

	a := [5]int {1,2,3,4,5}
	arr_test(&a)

}
