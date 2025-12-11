package main

import "testing"

func TestGeneratorChannels(t *testing.T) {
	tests := []struct {
		name    string
		k       int
		lenWant int
	}{{
		name:    "check generator",
		k:       5,
		lenWant: 5,
	}}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ch := GeneratorChannels(test.k)

			if len(ch) != test.lenWant {
				t.Errorf("expect %d got %d", test.lenWant, len(ch))
			}
		})
	}
}

func TestMerge(t *testing.T) {
	tests := []struct {
		name    string
		k       int
		lenWant int
	}{
		{
			name:    "check merge",
			k:       5,
			lenWant: 10,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ch1 := GeneratorChannels(test.k)
			ch2 := GeneratorChannels(test.k)

			var res []int

			resCh := Merge(ch1, ch2)

			for v := range resCh {
				res = append(res, v)
			}

			if len(res) != test.lenWant {
				t.Errorf("expect %d got %d", test.lenWant, len(res))
			}
		})
	}
}
