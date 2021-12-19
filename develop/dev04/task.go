package main

import (
	"sort"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===
Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.
Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	
}

func GetMap(s []string) map[string][]string {
	m := make(map[string][]string)
	res := make(map[string][]string)
	for _, i := range s {
		str := strings.ToLower(i)
		sorts := SortString(str)
		if _, ok := m[sorts]; ok {
			m[sorts] = append(m[sorts], str)
		} else {
			m[sorts] = []string{str}
		}
	}

	for _, v := range m {
		if len(v) <= 1 {
			continue
		}
		res[v[0]] = v[1:]
	}
	return res
}

type SortBy []rune

func (a SortBy) Len() int           { return len(a) }
func (a SortBy) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortBy) Less(i, j int) bool { return a[i] < a[j] }

func SortString(s string) string {
	r := []rune(s)
	sort.Sort(SortBy(r))
	return string(r)
}