package main

import "testing"

func TestCut(t *testing.T) {
	tests := []struct {
		input string
		index []int
		delimeter   string
		want string
		isSeparated bool
	}{
		{
			input: "one	two	three	first	second	third",
			index:   []int{1, 2},
			delimeter: "\t",
			want: "two	three",
			isSeparated: true,
		},
		{
			input:       "one two three first second third",
			index:       []int{1, 2},
			delimeter:     "\t",
			want:  "",
			isSeparated: true,
		},
		{
			input:       "one-two-three-first-second-third",
			index:       []int{1, 2},
			delimeter:     "-",
			want:  "two-three",
			isSeparated: true,
		},
		{
			input:       "one two three first second third",
			index:       []int{0},
			delimeter:     "-",
			want:  "one two three first second third",
			isSeparated: false,
		},
	}

	for _, v := range tests {
		res := separate(v.input, v.index, v.delimeter, v.isSeparated)
		if v.want != res {
			t.Errorf("Incorrect result, want %s, got %s", v.want, res)
		}
	}
}