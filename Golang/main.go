package main

import (
	"fmt"
)

func main() {
	// 	// go의 타입
	// 	var a string
	//  var b int64
	// 	var c float64

	// 	var d uint64
	// 	var e bool

	// 	var f interface{}  //any와 같다
	// 	var g any

	// a := 0
	// if a < 10 || a == 8 {

	// }
	// fmt.Print(sum(5, 7))

	var lineCnt int = 5
	var spaceCnt int = 4
	var starCnt int = 1
	for i := 1; i <= lineCnt; i++ {
		for j := 1; j <= spaceCnt; j++ {
			fmt.Print(" ")
		}
		for j := 1; j <= starCnt; j++ {
			fmt.Print("*")
		}
		if i < (lineCnt/2)+1 {
			spaceCnt -= 1
			starCnt += 2
			fmt.Println()

		} else {
			spaceCnt += 1
			starCnt -= 2
			fmt.Println()

		}

	}

}

// func sum(a int, b int) int {
// 	return a + b
// }

// func sum(a, b int) (int, int, int) {

// 	return a, b, a + b
// }
