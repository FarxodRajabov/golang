package main

import "fmt"

func main() {
	var a interface{} = true
	// str, ok := a.(string)

	// fmt.Println(str, ok)

	switch a.(type) {
	case int:
		fmt.Println("a is int")
	case string:
		fmt.Println("a is str")
	default:
		fmt.Printf("a is %T\n", a)
	}

}
