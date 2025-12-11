package main

import (
	"reflect"
	"testing"
)

func TestFindSimilarities(t *testing.T) {

	tests := []struct {
		name     string
		slice1   []int
		slice2   []int
		boolWant bool
		resWant  []int
	}{
		{
			name:     "no similarities",
			slice1:   []int{1, 2},
			slice2:   []int{3, 4},
			boolWant: false,
			resWant:  []int{},
		},
		{
			name:     "simple similarity",
			slice1:   []int{1, 2, 3},
			slice2:   []int{3, 4},
			boolWant: true,
			resWant:  []int{3},
		},
		{
			name:     "multiple similarities",
			slice1:   []int{65, 3, 58, 678, 64},
			slice2:   []int{64, 2, 3, 43},
			boolWant: true,
			resWant:  []int{64, 3},
		},
		{
			name:     "slice2 empty",
			slice1:   []int{1, 2, 3},
			slice2:   []int{},
			boolWant: false,
			resWant:  []int{},
		},
		{
			name:     "slice1 empty",
			slice1:   []int{},
			slice2:   []int{1, 2},
			boolWant: false,
			resWant:  []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotBool, gotSlice := FindSimilarities(tt.slice1, tt.slice2)

			if gotBool != tt.boolWant {
				t.Errorf("bool mismatch: want %v got %v", tt.boolWant, gotBool)
			}

			if !reflect.DeepEqual(gotSlice, tt.resWant) {
				t.Errorf("slice mismatch: want %v got %v", tt.resWant, gotSlice)
			}
		})
	}
}
