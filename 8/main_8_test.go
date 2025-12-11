package main

import (
	"fmt"
	"testing"
)

func TestMyWaitGroup(t *testing.T) {

	wg := NewMyWaitGroup(10)

	k := 2

	wg.Add(k)

	if len(wg.sem) != k {
		t.Errorf("expected len %d got %d", k, len(wg.sem))
	}

	defer close(wg.sem)

	go func() {
		for i := 1; i <= k; i++ {
			wg.Done()
		}
	}()

	fmt.Println("we are waiting...")

	wg.Wait()

	if len(wg.sem) != 0 {
		t.Errorf("expected len %d got %d", 0, len(wg.sem))
	}

	fmt.Println("finish")

}
