package main

import "fmt"

const cat = "siberian"

func main() {
	var price int = 30
	var name = "shrikant"
	phone := 9818232424
	const ph = 213123

	const (
		port = 3000
		host = 8000
	)

	fmt.Println(price)
	fmt.Println(name)
	fmt.Println(phone)
	fmt.Println(ph)
	fmt.Println(cat)
	fmt.Println(port)
	fmt.Println(host)
}
