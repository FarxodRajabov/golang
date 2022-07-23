package main

import "fmt"

func main() {
	// fmt.Println("Hello World")
	// fmt.Println("Hello")

	// var str string   // str = ""
	// var numInt int32 // 0 var numInt = 0
	// var bul bool     // false
	// var numFl float32

	// fmt.Println(str, numInt, bul, numFl)

	// name := "Farhod"
	// fmt.Println(name)

	// var name, age = "Farhod", 22

	// var (
	// 	name      = "Farhod"
	// 	age       = 22
	// 	isMarried = false
	// )
	// isMarried = true
	// fmt.Println(name, age, isMarried)

	var (
		name = "Farhod"
		age  = 22
	)

	c := fmt.Sprintf("My name is %s. and Im %d yers old", name, age)
	fmt.Println(c)
}
