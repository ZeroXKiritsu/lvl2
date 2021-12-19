package main

import (
	"errors"
	"strconv"
)

/*
=== Задача на распаковку ===
Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)
В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.
Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func unpackStrings(s string) (string, error) {
	if len(s) == 0 {
		return "", nil
	}

	currentString := []rune(s)[0]
	res := ""

	_, err := strconv.ParseInt(string(currentString), 10, 32)
	if err == nil {
		return "", errors.New("wrong type of data")
	}

	for _, v := range s[1:] {
		str := string(v)
		count, err := strconv.Atoi(str)
		if err != nil {
			res += string(currentString)
			currentString = v
		}
		if count == 0 {
			continue
		}
		for j := 0; j < count-1; j++ {
			res += string(currentString)
		}
	}
	res += string(currentString)
	return res, nil
}

func main() {

}