package main

import "fmt"

func addNumbers(nums ...int) int {
	total := 0

	for _, v := range nums {
		total = total + v
	}

	return total

}
func main() {

	arr := []int{3, 4, 2, 54, 22, 542, 54245, 234}

	// result := addNumbers(2, 32, 534, 235, 2345, 235)
	result := addNumbers(arr...)

	fmt.Println(result)

}
