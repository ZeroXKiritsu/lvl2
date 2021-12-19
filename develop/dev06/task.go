package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
=== Утилита cut ===
Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные
Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

var (
	fields      string
	delimeter   string
	isSeparated bool
)

func main() {
	flag.StringVar(&fields, "f", "", "выбрать поля (колонки)")
	flag.StringVar(&delimeter, "d", "\t", "Использовать другой разделитель")
	flag.BoolVar(&isSeparated, "s", false, "Выводить только строки с разделителем")
	flag.Parse()
	cutStart(fields, delimeter, isSeparated)
}

func cutStart(f string, d string, s bool) {
	if f == "-1" {
		fmt.Println("укажите флаг -f")
		return
	}
	arrFs := strings.Split(f, ",")
	arrF := []int{}
	for _, v := range arrFs {
		F, _ := strconv.Atoi(v)
		arrF = append(arrF, F)
	}

	cut(arrF, d, s)
}

func cut(f []int, d string, s bool) {
	reader := bufio.NewReader(os.Stdin)
	for {
		text, _ := reader.ReadString('\n')
		str := separate(text, f, d, s)
		if str == "" && s {
			continue
		}
		fmt.Println(str)
	}
}
func separate(buff string, i []int, delimiter string, s bool) (res string) {
	split := strings.Split(string(buff), delimiter)
	if len(split) == 1 && s {
		return ""
	}
	for _, v := range i {
		if v >= len(split) {
			return
		}
		if res != "" {
			res += delimiter
		}
		res += split[v]
	}
	return
}
