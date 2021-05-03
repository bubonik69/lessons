package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
	"unicode"
)

func fileReaderByLine(myChan chan string, path string) {
	// os.Open() opens specific file in
	// read-only mode and this return
	// a pointer of type os.
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("failed to open")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}
	for _, line := range text {
		//fmt.Println(each_ln)
		myChan <- line
	}

}

func calcWordToFile(ch chan []string, path string) {
	file, er := ioutil.ReadFile(path)
	if er != nil {
		fmt.Print("error due reading file")
	}
	text := string(file)
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	var newtext []string
	text = strings.ToLower(text)
	newtext = strings.FieldsFunc(text, f)
	fmt.Printf("Fields are: %v, length :%d\n", newtext, len(newtext))
	ch <- newtext
}
func worker3(ch chan []string) {
	myMap := make(map[string]int)
	for {
		v, ok := <-ch
		if !ok {
			fmt.Println("chanel worker's 3 was closed")
			break
		}
		for _, word := range v {
			_, ok := myMap[word]
			if !ok {
				myMap[word] = 1
			} else {
				myMap[word] += 1
			}
			if myMap[word] > 5 {
				fmt.Printf("word %s, met %d\n", word, myMap[word])

			}
		}
		fmt.Println(v)
	}
}

func main() {
	//var wg sync.WaitGroup
	ch1 := make(chan string)
	ch2 := make(chan []string)
	path := "hw10/text.txt"
	//wg.Wait()
	// start worker1 - devide text on the lines
	go fileReaderByLine(ch1, path)
	// tart worker2 - printing lines from channel
	go func() {
		numLIne := 1
		for {
			v, ok := <-ch1
			if !ok {
				fmt.Println("chanel worke's 2 was closed")
				break
			}
			fmt.Println("line - ", numLIne, v)
			numLIne++
		}
	}()
	//tart worker3 - calculate word and send to next worker slice
	go calcWordToFile(ch2, path)
	go worker3(ch2)

	select {
	case <-ch2:
		close(ch1)
		close(ch2)

	}
	time.Sleep(time.Second)
	fmt.Println("end programm")

}

// как сделать вейтгруппу если не знаешь количество передаваемых объектов
// как передать ссылку на файл
