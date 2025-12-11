package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {

	ch := make(chan int)

	go Generator(ch, 10)

	for v := range ch {
		fmt.Println(v)
	}

}

func Generator(ch chan int, n int) {

	for i := 1; i <= n; i++ {
		v := rand.IntN(10)

		ch <- v
	}
	close(ch)
}
