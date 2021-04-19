package hw4

import (
	"fmt"
	"math/rand"
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

func Hw4task1() {
	s1 := make([]int, 0, 0) // создаем слайс
	appendSlice(&s1, 1000)  // заполняем его 1000 рандомными элементами
	sortSlice(true, s1)     // сортируем его в порядке убывания
}
