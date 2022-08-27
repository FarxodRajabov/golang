package main

import "fmt"

func main() {
	go say("hello")
}

func say(world string) {
	fmt.Println(world)
}
