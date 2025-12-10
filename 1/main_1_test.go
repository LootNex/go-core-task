package main

import (
	"crypto/sha256"
	"reflect"
	"testing"
)

func TestCheckType(t *testing.T) {
	tests := []struct {
		name    string
		val     any
		wantRes reflect.Type
	}{
		{
			name:    "check string",
			val:     "Golang",
			wantRes: reflect.TypeOf(""),
		},
		{
			name:    "check int",
			val:     42,
			wantRes: reflect.TypeOf(0),
		},
		{
			name:    "check float64",
			val:     3.14,
			wantRes: reflect.TypeOf(0.0),
		},
		{
			name:    "check bool",
			val:     true,
			wantRes: reflect.TypeOf(true),
		},
		{
			name:    "check complex64",
			val:     1 + 2i,
			wantRes: reflect.TypeOf(1 + 2i),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := CheckType(test.val)
			if res != test.wantRes {
				t.Errorf("wrong result, expected %v got %v", test.wantRes, res)
			}
		})
	}
}

func TestConvertToStr(t *testing.T) {
	tests := []struct {
		name    string
		val     any
		wantRes string
	}{
		{
			name:    "check string",
			val:     "Golang",
			wantRes: "Golang",
		},
		{
			name:    "check int",
			val:     42,
			wantRes: "42",
		},
		{
			name:    "check float64",
			val:     3.14,
			wantRes: "3.14",
		},
		{
			name:    "check bool",
			val:     true,
			wantRes: "true",
		},
		{
			name:    "check complex64",
			val:     complex64(1 + 2i),
			wantRes: "(1+2i)",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := ConvertToStr(test.val)
			if res != test.wantRes {
				t.Errorf("wrong result, expected %v got %v", test.wantRes, res)
			}
		})
	}

}

func TestConvertToRune(t *testing.T) {

	tests := []struct {
		name    string
		val     string
		wantRes []rune
	}{
		{
			name:    "ASCII string",
			val:     "Golang",
			wantRes: []rune{'G', 'o', 'l', 'a', 'n', 'g'},
		},
		{
			name:    "Unicode string (Russian)",
			val:     "Привет",
			wantRes: []rune{'П', 'р', 'и', 'в', 'е', 'т'},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := ConvertToRune(test.val)
			if !reflect.DeepEqual(res, test.wantRes) {
				t.Errorf("wrong result, expected %v got %v", test.wantRes, res)
			}
		})
	}
}

func TestHash(t *testing.T) {
	tests := []struct {
		name string
		val  string
	}{
		{
			name: "hash ascii",
			val:  "Go",
		},
		{
			name: "hash unicode",
			val:  "Привет",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := Hash([]rune(test.val))

			want := sha256.Sum256([]byte(test.val))

			if res != want {
				t.Errorf("wrong result, expected %v got %v", want, res)
			}
		})
	}
}
