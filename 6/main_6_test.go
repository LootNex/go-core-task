package main

import "testing"

func TestGenerator(t *testing.T) {
	tests := []struct {
		name  string
		ch    chan int
		count int
	}{
		{
			name:  "check generator",
			ch:    make(chan int),
			count: 10,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			go Generator(test.ch, test.count)
			res := []int{}
			for v := range test.ch {
				res = append(res, v)
			}

			if len(res) != test.count {
				t.Errorf("wrong result, expected len res = %d", test.count)
			}
		})
	}

}
