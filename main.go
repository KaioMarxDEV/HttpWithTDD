package main

import "fmt"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Square struct {
	size float64
}

func (s Square) Area() float64 {
	return s.size * s.size
}

func (s Square) Perimeter() float64 {
	return s.size * 4
}

func printInformation(obj Shape) {
	fmt.Printf("%T\n", obj)
	fmt.Println("Area: ", obj.Area())
	fmt.Println("Perimeter:", obj.Perimeter())
	fmt.Println()
}

func main() {
	squareImplementation := Square{4}
	printInformation(squareImplementation)

	var anotherSquare Shape = Square{8}
	printInformation(anotherSquare)
}
