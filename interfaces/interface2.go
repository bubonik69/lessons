package main

import (
	"fmt"
	"math"
)

type mm int
type Figure interface {
	Area() float64
	Perimeter() float64
	fmt.Stringer
}
type square struct {
	side mm
}

type circle struct {
	radius mm
}
type rectangle struct {
	side1 mm
	side2 mm
}
type rhombus struct {
	side1 mm
	side2 mm
}

// описание поведения методов квадрата
func (f square) Area() float64 {
	return float64(f.side * f.side)
}
func (f square) Perimeter() float64 {
	return float64(f.side * 4)
}

func (f square) String() string {
	return fmt.Sprintf("This is square. Side is %d.", f.side)
}

// описание поведения методов круга
func (f circle) Area() float64 {
	return 3.14 * float64(f.radius*f.radius)
}
func (f circle) Perimeter() float64 {
	return 2 * math.Pi * float64(f.radius)
}
func (f circle) String() string {
	return fmt.Sprintf("This is circle. Radius is %d.", f.radius)
}

// описание поведения методов четырехугольника
func (m rectangle) Area() float64 {
	return float64(m.side1 * m.side2)
}
func (m rectangle) Perimeter() float64 {
	return float64(m.side1*2 + m.side2*2)
}
func (m rectangle) String() string {
	return fmt.Sprintf("This is rectangle, sides %d,%d.", m.side1, m.side2)
}

func validFigure(f []Figure) map[Figure]int {
	ret := make(map[Figure]int)
	for _, fig := range f {
		_, ok := ret[fig]
		if fig.Area() > 0 {
			if !ok {
				ret[fig] = 1
			} else {
				ret[fig] += 1
			}
		}
	}
	return ret

}

func main() {
	k := []Figure{square{side: 5}, square{side: 10}}
	k = append(k, square{side: 11})
	k = append(k, circle{radius: 11})
	k = append(k, rectangle{side1: 11, side2: 3})
	k = append(k, rectangle{side1: 11, side2: 3})
	k = append(k, rectangle{})
	k = append(k, circle{})
	k = append(k, rectangle{side1: 1, side2: 3})
	for _, s := range k {
		fmt.Printf("%s , Area = %v, Perimeter %v\n", s, s.Area(), s.Perimeter())
	}
	myMap := make(map[Figure]int)
	myMap = validFigure(k)
	fmt.Println(myMap)
	//fmt.Println(k)
}
