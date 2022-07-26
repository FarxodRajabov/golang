package main

import "fmt"

type Point struct {
	X int
	Y int
	Z string
}

func main() {

	// _-_-_-_-_ for(range) _-_-_-_-_

	// arr := []string{"a", "b", "c"}
	// for i, l := range arr {
	// 	fmt.Println(i, l)
	// }

	// _-_-_-_-_ Map() _-_-_-_-_

	// pointsMap := map[string]Point{}
	// nuwPointsMap := make(map[string]Point)

	// pointsMap["a"] = Point{
	// 	X: 1,
	// 	Y: 2,
	// }

	// fmt.Println(pointsMap)
	// fmt.Println(pointsMap["a"])

	newMap := map[string]Point{
		"x": {323, 222, "221112"},
		"a": {323, 222, "222222"},
		"b": {323, 222, "2222"},
	}

	for k, v := range newMap {
		fmt.Println(k)
		fmt.Println(v)
	}
}
