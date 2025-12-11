package main

import (
	"fmt"
	"math/rand"
)

func main() {

	var slice []int

	for i := 1; i <= 10; i++ {
		n := rand.Intn(100)
		slice = append(slice, n)
	}

	fmt.Println(slice)

	slice = sliceExample(slice)
	fmt.Println(slice)

	slice = addElements(&slice, 10)
	fmt.Println(slice)

	slice = copySlice(slice)
	fmt.Println(slice)

	slice = removeElement(slice, 2)
	fmt.Println(slice)

}

func sliceExample(slice []int) []int {

	var newslice []int

	for _, v := range slice {
		if v%2 == 0 {
			newslice = append(newslice, v)
		}
	}

	return newslice

}

func addElements(slice *[]int, val int) []int {

	*slice = append(*slice, val)

	return *slice
}

func copySlice(slice []int) []int {
	newslice := make([]int, len(slice))

	copy(newslice, slice)

	return newslice
}

func removeElement(slice []int, idx int) []int {

	slice = append(slice[:idx], slice[idx+1:]...)

	return slice
}
