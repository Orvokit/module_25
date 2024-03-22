/*
Написать программу для нахождения подстроки в кириллической подстроке. Программа должна запускаться с помощью команды:

 go run main.go --str "строка для поиска" --substr "поиска"

Для реализации такой работы с флагами воспользуйтесь пакетом flags, а для поиска подстроки в строке вам понадобятся руны.

Что нужно сделать
1. Спроектировать алгоритм поиска подстроки.
2. Определить строку и подстроку, используя флаги.
3. Написать алгоритм реализацию для работы со строками UTF-8 (для этого необходимо воспользоваться рунами).



*/

package main

import (
	"flag"
	"fmt"
)

func findStr(str string, substr string) bool {
	strRune := []rune(str)
	substrRune := []rune(substr)

str:
	for i, s := range strRune {
		if s == substrRune[0] {
			for j, v := range substrRune[1:] {
				if strRune[i+j+1] != v {
					continue str
				}
			}
			return true
		}
	}
	return false
}

func main() {
	var str, substr string
	flag.StringVar(&str, "str", " ", "str")
	flag.StringVar(&substr, "substr", "", "substr")

	flag.Parse()

	fmt.Println("Строка в которой ищем подстроку:", str, "\nПодстрока которую ищем:", substr)
	result := findStr(str, substr)
	if result == true {
		fmt.Println("Подстрока содержится в подстроке:", result)
	} else {
		fmt.Println("Подстрока НЕ содержится в подстроке:", result)
	}
}
