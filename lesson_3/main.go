package main

import "fmt"

type Point struct {
	X int
	Y int
	S string
}

func main() {
	// pointers()
	// structs()

	p1 := Point{
		X: 2,
		Y: 3,
		S: "Hello",
	}
	// p1.method()
	p2 := &p1
	p2.method()

}
func (p Point) method() {
	fmt.Println(p.X)
	fmt.Println(p.Y)
	fmt.Println(p.S)
}

func structs() {
	// p := Point{
	// 	X: 0,
	// 	Y: 2,
	// 	S: "Hello",
	// }
	// fmt.Println(p)
	// fmt.Println(p.X)
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
