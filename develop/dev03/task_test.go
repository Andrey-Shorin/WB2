package main

import (
	"fmt"
	"os"
	"testing"
)

func TestUnpack(t *testing.T) {
	var test = []SortUtil{
		{Filename: "test.txt", K: 1, N: false, U: false, R: false},
		{Filename: "test.txt", K: 2, N: false, U: false, R: false},
		{Filename: "test.txt", K: 1, N: false, U: true, R: false},
		{Filename: "test.txt", K: 1, N: false, U: false, R: true},
		{Filename: "test.txt", K: 2, N: false, U: false, R: true},
	}
	for i, tt := range test {

		testname := "test1"

		t.Run(testname, func(t *testing.T) {
			old := os.Stdout
			file, err := os.Create(fmt.Sprintf("test/res%d.txt", i))
			if err != nil {
				fmt.Println("Ошибка при создании временного файла:", err)
				return
			}

			os.Stdout = file
			readFile(&tt)
			os.Stdout = old
			file.Close()
			want, _ := os.ReadFile(fmt.Sprintf("test/test%d.txt", i))
			ansv, _ := os.ReadFile(fmt.Sprintf("test/res%d.txt", i))

			if len(want) != len(ansv) {
				t.Errorf("got %s, want %s", want, ansv)
			}

		})

	}
}
