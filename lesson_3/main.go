package main

import "fmt"

type Point struct {
	X int
	Y int
}

func main() {
	pointers()
	structs()
}

func structs() {
	p := Point{
		X: 0,
		Y: 2,
	}
	fmt.Println(p)
}

func pointers() {
	// a := "Hello World"
	// b := 42

	// fmt.Println(a)
	// fmt.Println(b)

	// p := &a
	// fmt.Println(p)
	// fmt.Println(*p)
	// *p = "GOGOGO"
	// fmt.Println(*p)
	// g := &b
	// *g = *g / 2
	// fmt.Println(*g)

}
