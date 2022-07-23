package main

import "fmt"

// func main() {
// 	s, s1, s2 := test1()
// 	a := test2()
// 	fmt.Println(s, s1, s2, a)
// }

// func test1() (a, b, c string) {
// 	a = "hello"
// 	b = "world"
// 	c = "!!!"
// 	return a, b, c
// }

// func test2() (a int) {
// 	a = 1
// 	return a
// }

func main() {
	// sum := 0

	// for i := 0; i < 10; i++ {
	// 	sum += i
	// }

	// fmt.Println(sum)

	// for {    // =======> while
	// 	sum += 10
	// }
	// fmt.Println(sum)

	// for sum < 100 {
	// 	sum += 20
	// 	fmt.Println(sum)
	// }

	// a := 0
	// for a < 11 {
	// 	if a == 9 {
	// 		fmt.Println(fmt.Sprintf(" a is %d", a))
	// 	} else {
	// 		fmt.Println(fmt.Sprintf(" a is not  %d", a))
	// 	}
	// 	a++
	// }

	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("Hello")
}
