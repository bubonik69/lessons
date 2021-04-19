package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
)

func appendSlice(s *[]int, elements int) {
	//заполняет слайс из n элементов рандомными числами
	for i := 0; i < elements; i++ {
		*s = append(*s, rand.Int())

	}
}

func sortSlice(descending bool, s []int) {
	//var s2 [10]int
	var min, temp int
	fmt.Println(s, cap(s), len(s))
	// сортировка по возрастанию
	switch descending {
	case false:
		for i := 0; i < len(s); i++ {
			min = s[i]
			for j := i; j < len(s); j++ {
				if s[j] < min {
					temp = min
					min = s[j]
					s[j] = temp
					s[i] = min
				}
			}
		}
		for i := 1; i < len(s); i++ {
			fmt.Println(s[i])
		}
	case true:
		for i := 0; i < len(s); i++ {
			min = s[i]
			for j := i; j < len(s); j++ {
				if s[j] > min {
					temp = min
					min = s[j]
					s[j] = temp
					s[i] = min
				}
			}
		}
		for i := 1; i < len(s); i++ {
			fmt.Println(s[i])
		}
	}

}

func hw4task1() {
	s1 := make([]int, 0, 0) // создаем слайс
	appendSlice(&s1, 1000)  // заполняем его 1000 рандомными элементами
	sortSlice(true, s1)     // сортируем его в порядке убывания
}

//************************************************************************************************
//********************************		task2					**********************************

func calculatingWords(text string, countlist int) {
	var textSlice []string
	var ss []int

	dictionary := make(map[string]int)
	dictionary2 := make(map[int][]string)
	text2 := strings.ToLower(text)
	textSlice = strings.Fields(text2)
	for _, k := range textSlice {
		word := strings.Trim(k, ".,!?-''")
		_, ok := dictionary[word]
		if !ok {
			dictionary[word] = 1
		} else {
			dictionary[word] += 1
		}
	}
	for k, v := range dictionary {
		dictionary2[v] = append(dictionary2[v], k)

	}
	for k, _ := range dictionary2 {
		ss = append(ss, k)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(ss)))
	//fmt.Println(dictionary)
	//fmt.Println(dictionary2)
	//fmt.Println(ss)
	for i := 0; i < countlist; i++ {
		k := ss[i]
		fmt.Printf(" Слово -%s, встречается - %d раз\n", dictionary2[k], k)
	}
}
func hw4task2() {
	text := "Вы вы вы вы вы вы вы вы мы мы мы мы мы мы в в в в в где де до до дом и и и и и или или или или a  a a? a, она она она я"
	calculatingWords(text, 3)
}
func main() {

	hw4task1()
	hw4task2()
	//() // замечание - если одно и то же слово встречается один
}
