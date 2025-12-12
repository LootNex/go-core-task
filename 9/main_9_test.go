package main

import (
	"math"
	"reflect"
	"testing"
)

func TestConverter(t *testing.T) {

	tests := []struct {
		name      string
		inputCH   chan uint8
		input     []uint8
		expectRes []float64
	}{
		{
			name:      "check converter",
			inputCH:   make(chan uint8),
			input:     []uint8{0, 1, 2, 10},
			expectRes: []float64{math.Pow(0, 3), math.Pow(1, 3), math.Pow(2, 3), math.Pow(10, 3)},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			var res []float64

			go func() {
				for _, v := range test.input {
					test.inputCH <- v
				}
				close(test.inputCH)
			}()
			ch2 := Converter(test.inputCH)

			for num := range ch2 {
				res = append(res, num)
			}

			if !reflect.DeepEqual(res, test.expectRes) {
				t.Errorf("expected %v got %v", test.expectRes, res)
			}

		})
	}

}
