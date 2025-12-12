package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {

	ch1 := Generator()

	ch2 := Converter(ch1)

	for num := range ch2 {
		fmt.Println(num)
	}

}

func Generator() chan uint8 {
	ch1 := make(chan uint8)

	go func() {
		defer close(ch1)
		for i := 1; i <= 5; i++ {
			n := uint8(rand.Intn(256))
			ch1 <- n
		}
	}()

	return ch1

}

func Converter(ch1 chan uint8) chan float64 {
	ch2 := make(chan float64)

	go func() {
		defer close(ch2)
		for val := range ch1 {

			val_float := float64(val)

			ch2 <- math.Pow(val_float, 3)

		}
	}()

	return ch2

}
