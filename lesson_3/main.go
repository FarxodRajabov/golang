package main

import "fmt"

// type Point struct {
// 	X int
// 	Y int
// 	S string
// }

// func main() {
// 	// pointers()
// 	// structs()

// 	p1 := Point{
// 		X: 2,
// 		Y: 3,
// 		S: "Hello",
// 	}
// 	// p1.method()
// 	p2 := &p1
// 	p2.method()

// }
// func (p Point) method() {
// 	fmt.Println(p.X)
// 	fmt.Println(p.Y)
// 	fmt.Println(p.S)
// }

// func structs() {
// 	// p := Point{
// 	// 	X: 0,
// 	// 	Y: 2,
// 	// 	S: "Hello",
// 	// }
// 	// fmt.Println(p)
// 	// fmt.Println(p.X)
// }

// func pointers() {
// 	// a := "Hello World"
// 	// b := 42

// 	// fmt.Println(a)
// 	// fmt.Println(b)

// 	// p := &a
// 	// fmt.Println(p)
// 	// fmt.Println(*p)
// 	// *p = "GOGOGO"
// 	// fmt.Println(*p)
// 	// g := &b
// 	// *g = *g / 2
// 	// fmt.Println(*g)

// }

// ------Array-------

// func main() {
// 	var a [2]string
// 	a[0] = "Hello"
// 	a[1] = "World"
// 	fmt.Println(a)

// 	numbers := [3]int{1, 2, 3}
// 	// numbers := [...]int{1, 2, 3}
// 	fmt.Println(numbers)
// 	fmt.Println(numbers[1])
// }

// -----slices-----

func main() {
	// letters := []string{"a", "b", "c"}
	// letters[0] = "1"
	// letters = append(letters, "d", "e")
	// fmt.Println(letters)

	createSlice := make([]string, 3)
	createSlice[0] = "1"
	createSlice[1] = "2"
	createSlice[2] = "3"
	createSlice = append(createSlice, "4")
	fmt.Println(createSlice)

}
