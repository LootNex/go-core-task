package main

import (
	"reflect"
	"testing"
)

func TestFindDifference(t *testing.T) {

	tests := []struct {
		name    string
		slice1  []string
		slice2  []string
		resWant []string
	}{
		{
			name:    "no difference",
			slice1:  []string{"a", "b", "c"},
			slice2:  []string{"a", "b", "c"},
			resWant: []string{},
		},
		{
			name:    "simple difference",
			slice1:  []string{"a", "b", "c"},
			slice2:  []string{"a", "b"},
			resWant: []string{"c"},
		},
		{
			name:    "multiple differences",
			slice1:  []string{"a", "b", "x", "y"},
			slice2:  []string{"x", "z"},
			resWant: []string{"a", "b", "y"},
		},
		{
			name:    "slice2 empty returns all slice1",
			slice1:  []string{"a", "b", "c"},
			slice2:  []string{},
			resWant: []string{"a", "b", "c"},
		},
		{
			name:    "slice1 empty returns empty",
			slice1:  []string{},
			slice2:  []string{"a", "b"},
			resWant: []string{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := FindDifference(test.slice1, test.slice2)

			if !reflect.DeepEqual(got, test.resWant) {
				t.Errorf("wrong result, expected %v got %v", test.resWant, got)
			}
		})
	}
}
