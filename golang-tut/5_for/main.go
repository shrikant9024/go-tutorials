package main

import "fmt"

func main() {

	i := 0

	for i < 3 {
		fmt.Println(i)
		i++
	}

	for i := 0; i < 4; i++ {
		fmt.Println(i)
	}

	for i := range 10 {
		fmt.Println(i)
	}
}
