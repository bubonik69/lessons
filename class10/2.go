package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	ch := make(chan int)
	wg.Add(2)
	go consumer(ch, &wg)
	go producer(ch, &wg)

	wg.Wait()
	fmt.Println(runtime.NumCPU())
}

func consumer(ch chan int, wg *sync.WaitGroup) {
	for {
		val, ok := <-ch
		if !ok {
			break
		}
		fmt.Println("recieve", val)
	}

	fmt.Println("consumer finished")
	wg.Done()
}
func producer(ch chan int, wg *sync.WaitGroup) {
	for i := 0; i < 4; i++ {
		fmt.Println("send", i)
		ch <- i
	}
	close(ch)
	fmt.Println("producer finished")
	wg.Done()
}


content, err := ioutil.ReadFile("golangcode.txt")
if err != nil {
log.Fatal(err)
}

// Convert []byte to string and print to screen

f.
text := string(content)
fmt.Println(text)