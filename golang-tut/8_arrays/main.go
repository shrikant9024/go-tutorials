package main

import "fmt"

func main() {
	var arr [4]int
	arr[0] = 2
	var chars [2]string
	chars[0] = "hello world"

	nums := [3]int{1, 2, 3}
	matrix := [3][3]int{{1, 1}, {2, 2}, {4, 5}}

	fmt.Println(arr[0])
	fmt.Println(chars[0])
	fmt.Println(arr)
	fmt.Println(nums)
	fmt.Println(matrix)
}
