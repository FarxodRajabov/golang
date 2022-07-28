package main

import "fmt"

type structHere struct {
	N1, N2 int
}

func (s *structHere) Sum() int {
	return s.N1 + s.N2
}

type InterfaceHere interface {
	Sum() int
}

func main() {
	var a InterfaceHere
	sh := structHere{1, 2}
	a = &sh
	// fmt.Println(a.Sum())
	os := otherStruct{5, 7}
	a = os
	fmt.Println(a.Sum())
	var i InterfaceHere = otherStruct{55, 44}

	fmt.Println(i.Sum())
}

type otherStruct struct {
	N1, N2 int
}

func (o otherStruct) Sum() int {
	return o.N1 + o.N2
}
