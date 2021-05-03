package main

import (
	"fmt"
	"sort"
	"strings"
)

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

	calculatingWords(text, 1)
}

func main() {
	hw4task2()
}
