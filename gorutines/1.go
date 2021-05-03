package main

import "fmt"

type mymap map[string]int

func (m *mymap) AddValue(s string, i int) {
	m = make(map[string]int)
	m[s] = i
	return m
}

func (m mymap) String() string {
	k := make(mymap)
	k = m
	//value, ok := m[]
	return fmt.Sprintf("my Map - %v", k)
}
func main() {

	var k mymap
	k = k.AddValue("Hi", 55)
	//go func() {
	//	k = k.AddValue("Hi", 55)
	//}()
	//go func() {
	//	k = k.AddValue("Mike", 11)
	//}()
	//go func() {
	//	k = k.AddValue("sis", 13)
	//}()
	fmt.Print(k)
}
