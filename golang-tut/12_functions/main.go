package main

import "fmt"

func addTwo(a int, b int) int {

	return a + b

}

func getLang() (string, string, bool) {
	return "golang", "c++", true
}

// func ProcessIt(fn func(a int) int) {
// 	fn(1)
// }

func ProcessIt() func(a int) int {
	return func(a int) int {
		return 4
	}
}

func main() {
	result := addTwo(3, 76)
	// lang1, lang2, lang3 := getLang()
	lang1, lang2, _ := getLang()

	// fn := func(a int) int {
	// 	return 2
	// }
	ans := ProcessIt()

	// fmt.Println(ans)

	// ans := fn(4)
	fmt.Println(ans(10))

	fmt.Println(result)
	fmt.Println(lang1, lang2)

}
