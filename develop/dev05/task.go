package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

/*
=== Утилита grep ===
Реализовать утилиту фильтрации (man grep)
Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

var (
	after       int
	before      int
	contextText int
	countBool   bool
	ignoreCase  bool
	invert      bool
	fixed       bool
	lineNum     bool
	filePath    string
)

func main() {
	flag.IntVar(&after, "A", 0, "печатать +N строк после совпадения")
	flag.IntVar(&before, "B", 0, "печатать +N строк до совпадения")
	flag.IntVar(&contextText, "C", 0, "(A+B) печатать ±N строк вокруг совпадения")
	flag.BoolVar(&countBool, "c", false, "(количество строк)")
	flag.BoolVar(&ignoreCase, "i", false, "(игнорировать регистр)")
	flag.BoolVar(&invert, "v", false, "(вместо совпадения, исключать)")
	flag.BoolVar(&fixed, "F", false, "точное совпадение со строкой, не паттерн")
	flag.BoolVar(&lineNum, "n", false, "печатать номер строки")
	flag.Parse()

	args := os.Args[len(os.Args)-1]
	t, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	text := string(t)

	if lineNum {
		fmt.Println("Find on", LineNum(args, text)+1, "line")
	}

	if countBool {
		count := Count(text)
		fmt.Println("Count of lines", count)
	}

	if ignoreCase {
		text = IgnoreCase(args, text, after, before, contextText)
	}

	if after > 0 {
		text = After(args, text, after)
	}

	if before > 0 {
		text = Before(args, text, before)
	}

	if contextText > 0 {
		text = ContextText(args, text, contextText)
	}

	if invert {
		text = Invert(args, text)
	}

	if fixed {
		fmt.Println(Fixed(args, text))
	}

	fmt.Println(text)
}

func After(str, file string, count int) (res string) {
	expr, err := regexp.Compile(str)
	if err != nil {
		fmt.Println(err)
	}
	rows := strings.Split(file, "\n")
	for i, v := range rows {
		if expr.MatchString(v) {
			res += v + "\n"
			if !(i == len(rows)-1) {
				for j, k := range rows {
					if j > i && j <= i+count {
						res += k + "\n"
					}
				}
			}
		}
	}
	return
}

func Before(str, file string, count int) (res string) {
	expr, err := regexp.Compile(str)
	if err != nil {
		fmt.Println(err)
	}
	rows := strings.Split(file, "\n")
	for i, v := range rows {
		if expr.MatchString(v) {
			for j, k := range rows {
				if j < i && j >= i-count {
					res += k + "\n"
				}
			}
			res += v
		}
	}
	return
}

func ContextText(str, file string, count int) (res string) {
	expr, err := regexp.Compile(str)
	if err != nil {
		fmt.Println(err)
	}
	rows := strings.Split(file, "\n")
	for i, v := range rows {
		if expr.MatchString(v) {
			for j, k := range rows {
				if j < i && j >= i-count {
					res += k + "\n"
				}
			}
			res += v + "\n"
			for j, k := range rows {
				if j > i && j <= i+count {
					if j == len(rows)-1 || j == i+count {
						res += k
					} else {
						res += k + "\n"
					}
				}
			}
		}
	}
	return
}

func Count(file string) (count int) {
	rows := strings.Split(file, "\n")
	for range rows {
		count++
	}
	return
}

func IgnoreCase(str string, file string, after int, before int, context int) (res string) {
	str = strings.ToLower(str)
	file = strings.ToLower(file)

	if after > 0 {
		res = After(str, file, after)
		after = 0
	}
	if before > 0 {
		res = Before(str, file, before)
		before = 0
	}

	if context > 0 {
		res = ContextText(str, file, contextText)
		contextText = 0
	}
	return res
}

func Invert(str, file string) (res string) {
	expr, err := regexp.Compile(str)
	if err != nil {
		fmt.Println(err)
	}
	rows := strings.Split(file, "\n")
	for i, v := range rows {
		if expr.MatchString(v) {
			continue
		}
		if i == len(rows)-1 {
			res += v
		} else {
			res += v + "\n"
		}
	}
	return
}

func Fixed(str, file string) bool {
	rows := strings.Split(file, "\n")
	for _, v := range rows {
		if v == str {
			return true
		}
	}
	return false
}

func LineNum(str, file string) int {
	expr, err := regexp.Compile(str) // Ф-ция Compile возвращает регулярное выражение
	if err != nil {
		fmt.Println(err)
	}
	rows := strings.Split(file, "\n")
	for i, v := range rows {
		if expr.MatchString(v) {
			return i
		}
	}
	return -1
}
