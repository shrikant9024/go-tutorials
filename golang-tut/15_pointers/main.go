package main

import "fmt"

func changeNumber(num *int) {
	*num = 4

	fmt.Println("in the function", *num)

}

func main() {
	num := 1
	changeNumber(&num)

	fmt.Println("Number after change", num)
}
