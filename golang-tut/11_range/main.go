package main

import "fmt"

func main() {

	//slices
	arr := []int{1, 2, 3}
	sum := 0
	for i, num := range arr {
		fmt.Println(num, i)
		sum += num
	}
	fmt.Println(sum)

	//map

	mpp := map[string]string{"fname": "John", "lname": "doe"}
	for k := range mpp {
		fmt.Println(k)
	}

	//string
	word := "golang"

	for i, c := range word {
		fmt.Println(i, string(c))
	}
}
