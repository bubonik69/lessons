package main

import (
	"fmt"
	"unicode/utf8"
)

func task1(min int) {
	fmt.Println(" task1")
	//Задача. В переменной min лежит число от 0 до 59. Определите в какую четверть часа попадает это число (в первую, вторую,
	//	третью или четвертую).
	if min < 15 {
		fmt.Println("первая четверть")
	}
	if min >= 15 && min < 30 {
		fmt.Println("вторая четверть")
	}
	if min >= 30 && min < 45 {
		fmt.Println("третяя четверть")
	}
	if min >= 45 {
		fmt.Println("четвертая четверть")
	}

}
func task2(a, b float64) {
	fmt.Println(" task2")
	//Задача2. Если переменная a равна или меньше 1, а переменная b больше или равна 3, то выведите сумму этих переменных,
	//иначе выведите их разность (результат вычитания). Проверьте работу скрипта при a и b, равном 1 и 3, 0 и 6, 3 и 5.
	if a <= 1 && b >= 3 {
		fmt.Println("a+b=", a+b)
	} else {
		fmt.Println("a-b=", a-b)
	}

}

func task3(num int) {
	fmt.Println(" task3")
	//Задача3 Переменная num может принимать 4 значения: 1, 2, 3 или 4. Если она имеет значение '1', то в переменную result
	//запишем 'зима', если имеет значение '2' – 'весна' и так далее. Задачу выполнить с использованием функции, которая просто
	//будет возвращать результат (сезон)
	switch num {
	case 1:
		fmt.Println("зима")
	case 2:
		fmt.Println("весна")
	case 3:
		fmt.Println("лето")
	case 4:
		fmt.Println("осень")
	default:
		fmt.Println("неправильное значение сезона")
	}

}
func main() {
	task1(55)
	task2(1, 3)
	task2(0, 6)
	task2(3, 5)
	task3(1)

	s := []string{"hello word"}
	r, size := utf8.DecodeRuneInString(s[0])
	for _, k := range r {
		fmt.Printf("%v -%v \n", r)
	}

	str := "Hello, 世界"

	for len(str) > 0 {
		r, size := utf8.DecodeRuneInString(str)
		fmt.Printf("%c %v\n", r, size)
		str = str[size:]
	}
}
