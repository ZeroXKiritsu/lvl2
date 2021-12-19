package main

import (
	"errors"
	"testing"
)

func TestUnpackStrings(t *testing.T) {
	tests := []struct {
		str         string
		expected    string
		expectedErr error
	}{
		{
			str:         "a4bc2d5e",
			expected:    "aaaabccddddde",
			expectedErr: nil,
		},
		{
			str:         "abcd",
			expected:    "abcd",
			expectedErr: nil,
		},
		{
			str:         "45",
			expected:    "",
			expectedErr: errors.New("wrong type of data"),
		},
		{
			str:         "",
			expected:    "",
			expectedErr: nil,
		},
	}

	for _, v := range tests {
		result, err := unpackStrings(v.str)
		if result != v.expected && err != v.expectedErr {
			t.Errorf("Incorrect result. Expected: %s, err:%v. Got %s, %v\n", v.expected, v.expectedErr, result, err)
		}
	}
}
