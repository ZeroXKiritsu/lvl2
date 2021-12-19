package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMap(t *testing.T) {
	tests := []struct {
		slice []string
		Map   map[string][]string
	}{
		{
			slice: []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"},
			Map: map[string][]string{
				"пятак":  {"пятка", "тяпка"},
				"листок": {"слиток", "столик"},
			},
		},
	}
	
	for _, v := range tests {
		result := GetMap(v.slice)
		if len(result) != len(v.Map) {
			t.Errorf("Incorrect result. Expect: %v, Got %v\n", v.Map, result)
		}
		for key, val := range result {
			assert.Equal(t, val, v.Map[key])
		}
	}
}
