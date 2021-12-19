package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

/*
=== Утилита sort ===
Отсортировать строки (man sort)
Основное
Поддержать ключи
-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки
Дополнительное
Поддержать ключи
-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	k := flag.Int("k", 1, "указание колонки для сортировки")
	n := flag.Bool("n", false, "сортировать по числовому значению")
	r := flag.Bool("r", false, "сортировать в обратном порядке")
	u := flag.Bool("u", false, "не выводить повторяющиеся строки")

	flag.Parse()

	args := flag.Args()
	if len(args) != 1 {
		log.Fatalln("usage: [flags] [file]")
	}

	file, err := ioutil.ReadFile(args[0])
	if err != nil {
		log.Fatal(err)
	}

	split := strings.Split(string(file), "\n")

	fmt.Println(sortString(split, *k-1, *n, *r, *u))
}

func Find(slice []string, value interface{}) bool {
	for _, v := range slice {
		if v == value {
			return true
		} else {
			continue
		}
	}
	return false
}

func sortString(arr []string, k int, n, r, u bool) []string {
	var key string
	internhip := make(map[string]string)
	var keySlice []string
	var intSlice []int
	var result []string

	for _, v := range arr {
		split := strings.Split(v, " ")

		if k > len(split) {
			key = split[0]
		} else {
			key = split[k]
		}

		if n {
			index := strings.IndexFunc(key, func(r rune) bool { return r < '0' || r > '9' })
			i := 0
			if index == 0 {
				i--
				convertedInt := strconv.Itoa(i)
				internhip[convertedInt] = v
				break
			}
			convertedInt := key[:index]
			internhip[convertedInt] = v
		} else {
			internhip[key] = v
		}
	}

	for key := range internhip {
		if n {
			intKey, err := strconv.Atoi(key)
			if err != nil {
				log.Fatal(err)
			}
			intSlice = append(intSlice, intKey)
		} else {
			keySlice = append(keySlice, key)
		}
	}

	if r {
		if n {
			sort.Sort(sort.Reverse(sort.IntSlice(intSlice)))
			for _, v := range intSlice {
				keySlice = append(keySlice, strconv.Itoa(v))
			}
		} else {
			sort.Sort(sort.Reverse(sort.StringSlice(keySlice)))
		}
	} else {
		if n {
			sort.Ints(intSlice)
			for _, v := range intSlice {
				keySlice = append(keySlice, strconv.Itoa(v))
			}
		} else {
			sort.Strings(keySlice)
		}
	}

	for _, v := range keySlice {
		str := internhip[v]
		if u {
			if Find(result, str) {
				continue
			} else {
				result = append(result, str)
			}
		}
	}
	return result
}
