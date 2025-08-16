package main

import (
	"fmt"
	"slices"
)

func main() {
	var nums = make([]int, 4, 5)
	// nums = append(nums, 2, 4, 5, 65, 2121, 3423, 1)
	nums[3] = 22

	fmt.Println(nums)
	fmt.Println(len(nums))
	fmt.Println(cap(nums))

	var nums2 = make([]int, len(nums))

	copy(nums2, nums)
	fmt.Println(nums2)

	var newarr = make([]int, 0)
	newarr = append(newarr, 3, 4, 5, 2, 1)

	fmt.Println(newarr)
	fmt.Println(newarr[:2]) //slice operator

	var numbers = []int{1, 2, 2, 2, 2}
	fmt.Println(numbers)

	arr1 := []int{1, 2, 3}
	var arr2 = []int{1, 2, 3}

	arr1 = append(arr1, 2)

	fmt.Println(slices.Equal(arr1, arr2))

	//2d
	matrix := [][]int{{1, 2}, {2, 3}}
	fmt.Println(matrix)

}
