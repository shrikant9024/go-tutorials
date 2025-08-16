package main

import "fmt"

func main() {

	// var age = 12

	// if age >= 18 {
	// 	fmt.Println("Person is an adult")
	// } else if age >= 12 {
	// 	fmt.Println("Person is an teen")
	// } else {
	// 	fmt.Println("Person is a kid")
	// }

	if age := 12; age >= 18 {
		fmt.Println("Person is adult")
	} else {
		fmt.Println("Person is teen or kid")
	}
}
