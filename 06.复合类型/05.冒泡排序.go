package main

import "fmt"

func main() {

	a := [10]int{0,9,2,6,4,5,3,7,8,1}
	fmt.Println(a)
	for i:=0;i<len(a)-1 ;i++  {
		for j:=0;j<len(a)-1-i;j++{
			if a[j]>a[j+1] {
				a[j],a[j+1] = a[j+1],a[j]
			}
		}
	}

	fmt.Println(a)
	
}
