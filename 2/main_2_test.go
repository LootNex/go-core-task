package main

import (
	"reflect"
	"testing"
)

func TestSliceExample(t *testing.T) {

	tests := []struct {
		name    string
		val     []int
		resWant []int
	}{
		{
			name:    "mix even/odd",
			val:     []int{1, 2, 3, 4, 5, 6},
			resWant: []int{2, 4, 6},
		},
		{
			name:    "only odd",
			val:     []int{1, 3, 8, 5, 7},
			resWant: []int{8},
		},
		{
			name:    "only even",
			val:     []int{2, 4, 6, 8},
			resWant: []int{2, 4, 6, 8},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := sliceExample(test.val)
			if !reflect.DeepEqual(res, test.resWant) {
				t.Errorf("wrong result, expected %v got %v", test.resWant, res)
			}
		})
	}
}

func TestAddElements(t *testing.T) {
	tests := []struct {
		name    string
		val     []int
		add     int
		resWant []int
	}{
		{
			name:    "add to normal slice",
			val:     []int{1, 2, 3},
			add:     10,
			resWant: []int{1, 2, 3, 10},
		},
		{
			name:    "add to empty slice",
			val:     []int{},
			add:     5,
			resWant: []int{5},
		},
		{
			name:    "add to one element slice",
			val:     []int{9},
			add:     1,
			resWant: []int{9, 1},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := addElements(&test.val, test.add)
			if !reflect.DeepEqual(res, test.resWant) {
				t.Errorf("wrong result, expected %v got %v", test.resWant, res)
			}
		})
	}
}

func TestCopySlice(t *testing.T) {
	tests := []struct {
		name    string
		val     []int
		resWant []int
	}{
		{
			name:    "copy normal slice",
			val:     []int{1, 2, 3},
			resWant: []int{1, 2, 3},
		},
		{
			name:    "copy empty slice",
			val:     []int{},
			resWant: []int{},
		},
		{
			name:    "copy single element slice",
			val:     []int{5},
			resWant: []int{5},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := copySlice(test.val)
			if !reflect.DeepEqual(res, test.resWant) {
				t.Errorf("wrong result, expected %v got %v", test.resWant, res)
			}
		})
	}
}

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		name    string
		val     []int
		idx     int
		resWant []int
	}{
		{
			name:    "remove middle",
			val:     []int{10, 20, 30, 40},
			idx:     1,
			resWant: []int{10, 30, 40},
		},
		{
			name:    "remove first",
			val:     []int{7, 8, 9},
			idx:     0,
			resWant: []int{8, 9},
		},
		{
			name:    "remove last",
			val:     []int{1, 2, 3, 4},
			idx:     3,
			resWant: []int{1, 2, 3},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := removeElement(test.val, test.idx)
			if !reflect.DeepEqual(res, test.resWant) {
				t.Errorf("wrong result, expected %v got %v", test.resWant, res)
			}
		})
	}
}
