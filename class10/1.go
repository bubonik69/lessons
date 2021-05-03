package main

import (
	"fmt"
	"time"
)

func main() {
	//myInt := 0
	//chanel := make(chan int,2)
	//for i := 0; i < 5; i++ {
	//	myInt = i
	//	chanel <- myInt
	//	fmt.Println(<-chanel)
	//
	//}
	ch := make(chan int)
	go sum(33, 44, ch)
	go multiply(ch)
	time.Sleep(1 * time.Second)
}

func sum(a int, b int, ch chan int) {
	fmt.Println(a + b)
	ch <- a + b
}
func multiply(ch chan int) {
	sum := <-ch
	fmt.Println(sum * 2)

}
