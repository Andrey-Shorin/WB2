package main

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

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type SortUtil struct { //util
	Filename string
	K        int
	N        bool
	R        bool
	U        bool
	data     [][]string
}

func main() {

	K := flag.Int("k", 1, "column for sorting")
	N := flag.Bool("n", false, "numerical sort")
	U := flag.Bool("u", false, "remove duplicates")
	R := flag.Bool("r", false, "reversed order")

	flag.Parse()
	if len(flag.Args()) != 1 {
		return
	}
	su := SortUtil{Filename: flag.Args()[0], K: *K, N: *N, U: *U, R: *R}

	err := readFile(&su)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func readFile(su *SortUtil) error {

	file, err := os.Open(su.Filename)
	if err != nil {
		return err
	}
	defer file.Close()

	err = su.Process(file)
	if err != nil {
		return err
	}
	su.Sort()
	su.Print()

	return nil
}

func (su *SortUtil) Process(file *os.File) error {
	var hasEnoughCols bool

	su.data = make([][]string, 0)

	scanner := bufio.NewScanner(file)
	i := 0

	for scanner.Scan() {
		line := scanner.Text()
		su.data = append(su.data, make([]string, 0, 10))
		su.data[i] = append(su.data[i], line)
		su.data[i] = append(su.data[i], strings.Split(line, " ")...)
		if len(su.data[i])-1 >= su.K {
			hasEnoughCols = true
		}
		i++
	}

	if !hasEnoughCols {
		su.K = 0
	}

	if su.U {
		su.data = su.removeDupls()
	}

	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func (su *SortUtil) removeDupls() [][]string {
	mp := make(map[string]struct{}, len(su.data))
	res := make([][]string, 0, len(su.data))
	var str string
	var err error
	for _, v := range su.data {
		if len(v) <= su.K {
			str = ""
			if su.N {
				str = "0"
			}
		} else {
			str = v[su.K]
			if su.N {
				_, err = strconv.Atoi(str)
				if err != nil {
					str = "0"
				}
			}
		}

		if _, ok := mp[str]; !ok {
			mp[str] = struct{}{}
			res = append(res, v)
		}
	}
	return res
}

func (su *SortUtil) lessNum(i, j int) bool {
	var num1, num2 int
	var err error
	if len(su.data[i]) > su.K {
		num1, err = strconv.Atoi(su.data[i][su.K])
		if err != nil {
			num1 = 0
		}
	}
	if len(su.data[j]) > su.K {
		num2, err = strconv.Atoi(su.data[j][su.K])
		if err != nil {
			num2 = 0
		}
	}
	if num1 != num2 {
		return num1 < num2
	}
	return su.data[i][0] < su.data[j][0]
}

func (su *SortUtil) lessString(i, j int) bool {
	if len(su.data[i]) <= su.K && len(su.data[j]) <= su.K {
		return su.data[i][0] < su.data[j][0]
	}
	if len(su.data[i]) <= su.K && len(su.data[j]) > su.K {
		return true
	}
	if len(su.data[i]) > su.K && len(su.data[j]) <= su.K {
		return false
	}
	return su.data[i][su.K] < su.data[j][su.K]
}

func (su *SortUtil) Sort() {
	suLess := su.lessString

	if su.N {
		suLess = su.lessNum
	}

	if su.R {
		sort.Slice(su.data, func(i, j int) bool {
			return !suLess(i, j)
		})
		return
	}
	sort.Slice(su.data, suLess)

}

func (su *SortUtil) Print() {
	for i := 0; i < len(su.data); i++ {
		if i+1 == len(su.data) {
			fmt.Printf("%s", su.data[i][0])
		} else {
			fmt.Printf("%s\n", su.data[i][0])
		}
	}
}
