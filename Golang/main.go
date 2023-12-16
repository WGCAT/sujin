package main

import "fmt"

func main() {
	// 	// go의 타입
	// 	var a string
	//  var b int64
	// 	var c float64

	// 	var d uint64
	// 	var e bool

	// 	var f interface{}  //any와 같다
	// 	var g any

	a := 0
	if a < 10 || a == 8 {

	}

	fmt.Print(sum(5, 7))
}

// func sum(a int, b int) int {
// 	return a + b
// }

func sum(a, b int) (int, int, int) {

	return a, b, a + b
}
