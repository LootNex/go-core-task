package main

import (
	"reflect"
	"testing"
)

func TestAdd(t *testing.T) {

	tests := []struct {
		name    string
		mapa    StringIntMap
		key     string
		value   int
		resWant StringIntMap
	}{
		{
			name:    "first test",
			mapa:    StringIntMap{Storage: map[string]int{"first": 1, "second": 2}},
			key:     "third",
			value:   3,
			resWant: StringIntMap{Storage: map[string]int{"first": 1, "second": 2, "third": 3}},
		},
		{
			name:    "second test",
			mapa:    StringIntMap{Storage: map[string]int{}},
			key:     "first",
			value:   1,
			resWant: StringIntMap{Storage: map[string]int{"first": 1}},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mapa.Add(test.key, test.value)

			if !reflect.DeepEqual(test.mapa, test.resWant) {
				t.Errorf("wrong result, expected %v got %v", test.resWant, test.mapa)
			}
		})
	}

}

func TestRemove(t *testing.T) {

	tests := []struct {
		name    string
		mapa    StringIntMap
		key     string
		resWant StringIntMap
	}{
		{
			name:    "remove existing key",
			mapa:    StringIntMap{Storage: map[string]int{"first": 1, "second": 2}},
			key:     "first",
			resWant: StringIntMap{Storage: map[string]int{"second": 2}},
		},
		{
			name:    "remove non-existing key",
			mapa:    StringIntMap{Storage: map[string]int{"first": 1}},
			key:     "second",
			resWant: StringIntMap{Storage: map[string]int{"first": 1}},
		},
		{
			name:    "remove from empty map",
			mapa:    StringIntMap{Storage: map[string]int{}},
			key:     "first",
			resWant: StringIntMap{Storage: map[string]int{}},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mapa.Remove(test.key)

			if !reflect.DeepEqual(test.mapa, test.resWant) {
				t.Errorf("wrong result, expected %v got %v", test.resWant, test.mapa)
			}
		})
	}
}

func TestCopy(t *testing.T) {

	tests := []struct {
		name    string
		mapa    StringIntMap
		resWant map[string]int
	}{
		{
			name:    "copy non-empty map",
			mapa:    StringIntMap{Storage: map[string]int{"first": 1, "second": 2}},
			resWant: map[string]int{"first": 1, "second": 2},
		},
		{
			name:    "copy empty map",
			mapa:    StringIntMap{Storage: map[string]int{}},
			resWant: map[string]int{},
		},
		{
			name:    "copy nil storage",
			mapa:    StringIntMap{Storage: nil},
			resWant: map[string]int{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := test.mapa.Copy()

			if !reflect.DeepEqual(got, test.resWant) {
				t.Errorf("wrong result, expected %v got %v", test.resWant, got)
			}

			if len(got) > 0 {
				for k := range got {
					got[k] = got[k] + 100
					break
				}
				if reflect.DeepEqual(got, test.mapa.Storage) {
					t.Errorf("map must be copied, but underlying map was modified too")
				}
			}
		})
	}
}

func TestExists(t *testing.T) {

	tests := []struct {
		name    string
		mapa    StringIntMap
		key     string
		resWant bool
	}{
		{
			name:    "key exists",
			mapa:    StringIntMap{Storage: map[string]int{"first": 1, "second": 2}},
			key:     "first",
			resWant: true,
		},
		{
			name:    "key does not exist",
			mapa:    StringIntMap{Storage: map[string]int{"first": 1}},
			key:     "second",
			resWant: false,
		},
		{
			name:    "empty map",
			mapa:    StringIntMap{Storage: map[string]int{}},
			key:     "first",
			resWant: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := test.mapa.Exists(test.key)

			if got != test.resWant {
				t.Errorf("wrong result, expected %v got %v", test.resWant, got)
			}
		})
	}
}

func TestGet(t *testing.T) {

	tests := []struct {
		name       string
		mapa       StringIntMap
		key        string
		wantVal    int
		wantExists bool
	}{
		{
			name:       "get existing key",
			mapa:       StringIntMap{Storage: map[string]int{"first": 1, "second": 2}},
			key:        "first",
			wantVal:    1,
			wantExists: true,
		},
		{
			name:       "get non-existing key",
			mapa:       StringIntMap{Storage: map[string]int{"first": 1}},
			key:        "second",
			wantVal:    0,
			wantExists: false,
		},
		{
			name:       "get from empty map",
			mapa:       StringIntMap{Storage: map[string]int{}},
			key:        "first",
			wantVal:    0,
			wantExists: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gotVal, gotExists := test.mapa.Get(test.key)

			if gotExists != test.wantExists {
				t.Errorf("wrong exists flag, expected %v got %v", test.wantExists, gotExists)
			}

			if gotVal != test.wantVal {
				t.Errorf("wrong value, expected %v got %v", test.wantVal, gotVal)
			}
		})
	}
}
