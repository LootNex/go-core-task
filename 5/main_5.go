package main

import "fmt"

func main() {

	a := []int{65, 3, 58, 678, 64}
	b := []int{64, 2, 3, 43}

	fmt.Println(FindSimilarities(a, b))
}

func FindSimilarities(a, b []int) (bool, []int) {

	mapa := make(map[int]bool)

	res := []int{}

	for _, v := range a {
		mapa[v] = true
	}

	for _, v := range b {
		if _, ok := mapa[v]; ok {
			res = append(res, v)
		}
	}

	if len(res) == 0 {
		return false, res
	}

	return true, res
}
