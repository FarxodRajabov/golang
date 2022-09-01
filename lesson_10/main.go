package main

import (
	"fmt"
	"sync"
	"time"
)

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// func main() {
// 	//             --------->x
// 	// ------------|------->
// 	ch := make(chan string)

// 	var wg sync.WaitGroup

// 	wg.Add(2)
// 	go func() {
// 		time.Sleep(3 * time.Second)
// 		// fmt.Println("Go")

// 		ch <- "Go"

// 		wg.Done()
// 		// ch <- 0
// 	}()

// 	go func() {
// 		// time.Sleep(2 * time.Second)
// 		// fmt.Println("Go2")

// 		fmt.Println(<-ch)

// 		wg.Done()
// 		// ch <- 0
// 	}()

// 	fmt.Println("main")
// 	wg.Wait()

// 	// <-ch
// 	defer fmt.Println("Defer")

// 	return
// }

type Counter struct {
	mu sync.Mutex
	c  map[string]int
}

func (c *Counter) Inc(key string) {
	c.mu.Lock()
	c.c[key]++
	c.mu.Unlock()
}
func (c *Counter) Value(key string) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.c[key]
}

func main() {
	key := "test"

	c := Counter{c: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc(key)
	}

	time.Sleep(time.Second * 3)
	fmt.Println(c.Value(key))
}