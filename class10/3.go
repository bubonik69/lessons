package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sync"
)

func main() {
	ch := make(chan string, 10)

	var wg sync.WaitGroup
	wg.Add(2)
	go reader(ch, &wg)
	go writer(ch, &wg)

	wg.Wait()
}

func reader(ch chan<- string, wg *sync.WaitGroup) {
	f, err := os.Open("./files/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	content, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range content {
		fmt.Println("Send: ", string(v))
		ch <- string(v)
	}

	close(ch)

	wg.Done()
}

func writer(ch <-chan string, wg *sync.WaitGroup) {
	f, err := os.Create("files/output.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	for {
		v, ok := <-ch
		if !ok {
			break
		}

		fmt.Println("Received: ", v)

		_, err := f.WriteString(v)
		if err != nil {
			log.Fatal(err)
		}
	}

	wg.Done()
}
