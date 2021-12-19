package main

import "testing"

func TestAfter(t *testing.T) {
	args := "солнце"
	text := "стол рука чашка\nдом солнце игра\nвысоко мышь луч\nцифра жук зебра\nмышь его говори"
	count := 1
	if After(args, text, count) != "дом солнце игра\nвысоко мышь луч\n" {
		t.Error("результат не совпал с ожидаемым значением", After(args, text, count))
	}
}

func TestBefore(t *testing.T) {
	args := "жук"
	text := "стол рука чашка\nдом солнце игра\nвысоко мышь луч\nцифра жук зебра\nмышь его говори"
	count := 3
	if Before(args, text, count) != "стол рука чашка\nдом солнце игра\nвысоко мышь луч\nцифра жук зебра" {
		t.Error("результат не совпал с ожидаемым значением:", Before(args, text, count))
	}
}

func TestContextText(t *testing.T) {
	args := "солнце"
	text := "стол рука чашка\nдом солнце игра\nвысоко мышь луч\nцифра жук зебра\nмышь его говори"
	count := 1
	if ContextText(args, text, count) != "стол рука чашка\nдом солнце игра\nвысоко мышь луч" {
		t.Error("результат не совпал с ожидаемым значением:", ContextText(args, text, count))
	}
}

func TestCount(t *testing.T) {
	args := "стол рука чашка\nдом солнце игра\nвысоко мышь луч\nцифра жук зебра\nмышь его говори"
	if Count(args) != 5 {
		t.Error("результат не совпал с ожидаемым значением:", Count(args))
	}
}

func TestIgnoreCase(t *testing.T) {
	args := "солНЦЕ"
	text := "стол рука чашка\nдом сОлнце игра\nвысоко мышь луч\nцифра жук зебра\nмышь его говори"
	after = 2
	if IgnoreCase(args, text, after, before, contextText) != "дом солнце игра\nвысоко мышь луч\nцифра жук зебра\n" {
		t.Error("результат не совпал с ожидаемым значением:", IgnoreCase(args, text, after, before, contextText))
	}
}

func TestInvert(t *testing.T) {
	args := "мышь"
	text := "стол рука чашка\nдом солнце игра\nвысоко мышь луч\nцифра жук зебра\nмышь его говори"
	if Invert(args, text) != "стол рука чашка\nдом солнце игра\nцифра жук зебра\n" {
		t.Error("результат не совпал с ожидаемым значением:", Invert(args, text))
	}
}