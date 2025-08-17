package main

import (
	"fmt"
	"time"
)

// func printCount(count int) {
// 	fmt.Println(count)
// }
//used go keyword to do non blcker task to perfrom parallely on multiple threads(concurrent)

func main() {

	for i := 0; i < 10; i++ {
		// go printCount(i)
		go func(i int) {
			fmt.Println(i)
		}(i)
	}

	time.Sleep(time.Second * 2)

}
