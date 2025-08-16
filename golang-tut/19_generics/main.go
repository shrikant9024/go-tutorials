package main

import "fmt"

func printSlice[T comparable](items []T, name string) {

	for _, item := range items {
		fmt.Println(item, name)
	}

}

type stack[T any] struct {
	element []T
}

func main() {

	names := []string{"abdul", "shikha", "john"}
	numbers := []int{1, 245, 5, 123, 5}

	myStack := stack[string]{
		element: []string{"jhohn", "cenna"},
	}

	fmt.Println(myStack)
	printSlice(names, "Bot")
	printSlice(numbers, "bOB")
}
