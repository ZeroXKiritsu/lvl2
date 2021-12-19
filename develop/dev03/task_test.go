package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortString(t *testing.T) {
	tests := []struct {
		slice    []string
		k        int
		n, r, u  bool
		expected []string
	}{
		{
			slice:    []string{"e aba", "g aca", "c bab", "d bba", "f aaa", "b aaa"},
			k:        0,
			n:        false,
			r:        false,
			u:        false,
			expected: []string{"b aaa", "c bab", "d bba", "e aba", "f aaa", "g aca"},
		},
		{
			slice:    []string{"aa", "ba", "ab", "bb", "ca"},
			k:        0,
			n:        false,
			r:        true,
			u:        false,
			expected: []string{"ca", "bb", "ba", "ab", "aa"},
		},
		{
			slice:    []string{"e aba", "g aca", "c bab", "d bba", "b aaa"},
			k:        1,
			n:        false,
			r:        true,
			u:        false,
			expected: []string{"d bba", "c bab", "g aca", "e aba", "b aaa"},
		},
		{
			slice:    []string{"aa", "ba", "ab", "bb", "ca", "aa", "ca"},
			k:        0,
			n:        false,
			r:        false,
			u:        true,
			expected: []string{"aa", "ab", "ba", "bb", "ca"},
		},
		{
			slice:    []string{"e aba", "g aca", "c bab", "d bba", "f aaa", "b aaa", "g aca", "d bba", "f aaa", "b aaa"},
			k:        0,
			n:        false,
			r:        true,
			u:        true,
			expected: []string{"g aca", "f aaa", "e aba", "d bba", "c bab", "b aaa"},
		},
	}

	for _, v := range tests {
		fmt.Println("before: ", v)
		res := sortString(v.slice, v.k, v.n, v.r, v.u)
		fmt.Println("result: ", res)
		fmt.Println("expected: ", v.expected)
		assert.Equal(t, res, v.expected)
	}
}
