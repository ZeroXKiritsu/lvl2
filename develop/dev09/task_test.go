package main

import "testing"

func BenchmarkWget(b *testing.B) {
	url := "https://www.google.ru/"
	fileName := "downloads.txt"

	for i := 0; i < b.N; i++ {
		wget(url, fileName, 0)
	}
}