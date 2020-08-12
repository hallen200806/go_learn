package main

import (
	"time"
	"fmt"
)

func main() {
	/*
	时间间隔，没过一个间隔执行一次
	 */
	ticker := time.NewTicker(2 * time.Second)

	i := 1
	for {
		<- ticker.C
		fmt.Println("执行了ticker",i,"次")
		if i == 6 {
			break
		}
		i++
	}

}
