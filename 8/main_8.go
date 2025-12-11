package main

import (
	"fmt"
	"time"
)

type MyWaitGroup struct {
	sem chan struct{}
}

func NewMyWaitGroup(gorutines int) *MyWaitGroup {
	return &MyWaitGroup{
		sem: make(chan struct{}, gorutines),
	}
}

func (wg *MyWaitGroup) Add(k int) {
	for i := 1; i <= k; i++ {
		wg.sem <- struct{}{}
	}
}

func (wg *MyWaitGroup) Done() {
	<-wg.sem
}

func (wg *MyWaitGroup) Wait() {
	for len(wg.sem) > 0 {
		time.Sleep(time.Second)
	}

}

func main() {

	Mywg := NewMyWaitGroup(100)

	for i := 1; i <= 10; i++ {
		Mywg.Add(1)
		go func() {
			defer Mywg.Done()
			time.Sleep(2 * time.Second)
			fmt.Printf("id %d\n", i)
		}()
	}

	Mywg.Wait()

	fmt.Println("Finish...")

}
