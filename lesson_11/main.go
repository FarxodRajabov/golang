package main

import (
	"fmt"
	"log"
)

// import (
// 	"time"
// )

// func main() {
// 	go test()
// 	time.Sleep(300 * time.Millisecond)
// }

// func test() int {
// 	a := []int{1, 2, 3}
// 	return a[3]
// }

func main() {
	devide(4, 0)
	fmt.Println("222")
}

func devide(a, b int) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("panic happened:", err)
		}
	}()
	fmt.Println(a / b)
}
