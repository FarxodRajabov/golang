package main

import "fmt"

func main() {
	pointers()
}

func pointers() {
	a := "Hello World"
	b := 42

	fmt.Println(a)
	fmt.Println(b)

	p := &a
	fmt.Println(p)
}
