package main

import "fmt"

//  _-_-_-_ func methods, work with callback _-_-_-_

// func math(callback func(int, int) int, str string) int {
// 	nums := callback(5, 4)
// 	fmt.Println(str)
// 	return nums
// }

// func main() {
// 	sumCallback := func(n1, n2 int) int {
// 		return n1 + n2
// 	}
// 	result := math(sumCallback, "plus")
// 	fmt.Println(result)

// 	multCallback := func(n1, n2 int) int {
// 		return n1 * n2
// 	}

// 	result = math(multCallback, "mult")
// 	fmt.Println(result)
// }

//  _-_-_-_ zamikaniya _-_-_-_

func totalPrice(initPrice int) func(int) int {
	sum := initPrice
	return func(i int) int {
		sum += i
		return sum
	}
}

func main() {
	orderPrice := totalPrice(1)
	fmt.Println(orderPrice(88))
	fmt.Println(orderPrice(100))
}
