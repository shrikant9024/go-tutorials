package main

import (
	"fmt"
)

func main() {

	// i := 0

	// switch i {
	// case 1:
	// 	fmt.Println("print one")
	// case 2:
	// 	fmt.Println("print two")
	// default:
	// 	fmt.Println("the value is 0")
	// }

	// switch time.Now().Weekday() {
	// case time.Saturday, time.Sunday:
	// 	fmt.Println("Its a weekened, lets party")
	// default:
	// 	fmt.Println("Its a weekday, lets work")
	// }

	whoamI := func(i interface{}) {
		switch i.(type) {
		case int:
			fmt.Println("the number is an integer")
		case string:
			fmt.Println("the number is a string")
		case bool:
			fmt.Println("the number is a bool")
		default:
			fmt.Println("other case")

		}
	}
	whoamI(13.2)
}
