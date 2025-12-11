package main

import "fmt"

func main() {

	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig", "null"}

	result := FindDifference(slice1, slice2)

	fmt.Println(result)

}

func FindDifference(slice1, slice2 []string) []string {

	mapa := make(map[string]bool)

	res := []string{}

	for _, v := range slice2 {
		mapa[v] = true
	}

	for _, v := range slice1 {
		if _, ok := mapa[v]; !ok {
			res = append(res, v)
		}
	}

	return res

}
