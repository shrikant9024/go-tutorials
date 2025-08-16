package main

import (
	"fmt"
	"maps"
)

func main() {
	mp := make(map[string]string)
	mp["phone"] = "cobalt"
	mp["phones"] = "asd"

	delete(mp, "phone")
	clear(mp)
	fmt.Println(mp)
	fmt.Println(len(mp))

	mpp := map[int]int{1: 2, 3: 1, 2: 1}
	fmt.Println(mpp[2])

	v, ok := mpp[2]

	fmt.Println(v)
	if ok {
		fmt.Println("element is in map")
	} else {
		fmt.Println("element is not there")
	}

	m1 := map[string]int{"phone": 2, "cable": 3, "animal": 2}
	m2 := map[string]int{"phone": 2, "cable": 3, "animal": 2}

	fmt.Println(maps.Equal(m1, m2))
}
