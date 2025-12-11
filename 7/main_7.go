package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
)

func main() {

	var channels []chan int

	for i := 1; i <= 5; i++ {
		ch := GeneratorChannels(5)
		channels = append(channels, ch)
	}

	channel := Merge(channels...)

	for v := range channel {
		fmt.Println(v)
	}

}

func GeneratorChannels(k int) chan int {

	ch := make(chan int, k)

	for i := 1; i <= k; i++ {
		n := rand.IntN(10)
		ch <- n
	}

	close(ch)
	return ch

}

func Merge(chans ...chan int) chan int {
	res := make(chan int)

	var wg sync.WaitGroup

	wg.Add(len(chans))

	for _, ch := range chans {
		go func(ch chan int) {
			defer wg.Done()
			for v := range ch {
				res <- v
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(res)
	}()

	return res

}
