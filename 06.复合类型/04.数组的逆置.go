package main

import "fmt"

func main() {
	a := [10]int{0,1,2,3,4,5,6,7,8,9}
	fmt.Println(a)
	i := 0
	j := len(a)-1
	for i<j {
		a[i],a[j] = a[j],a[i]
		i++
		j--
	}
	fmt.Println(a)
}
