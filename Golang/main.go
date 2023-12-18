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
	// fmt.Print(star(11))

	// for i := 1; i <= lineCnt; i++ {
	// 	for j := 1; j <= spaceCnt; j++ {
	// 		fmt.Print(" ")
	// 	}
	// 	for j := 1; j <= starCnt; j++ {
	// 		fmt.Print("*")
	// 	}
	// 	if i < (lineCnt / 2) {
	// 		spaceCnt -= 1
	// 		starCnt += 2
	// 		fmt.Println()

	// 	} else {
	// 		spaceCnt += 1
	// 		starCnt -= 2
	// 		fmt.Println()

	// 	}

	// }
	var lineCnt1 int = 10
	var lineCnt2 int = 8
	var lineCnt3 int = 7
	var spaceCnt int = 3
	var starCnt int = 1
	for i := 1; i <= lineCnt1; i++ {

		for j := 1; j <= starCnt; j++ {
			fmt.Print("*")
		}
		if i < lineCnt1/2 {
			starCnt += 1
			fmt.Println()
		} else if i == lineCnt1/2 {
			fmt.Println()
		} else if i > lineCnt1/2 && i < lineCnt1 {
			starCnt -= 1
			fmt.Println()
		} else if i == lineCnt1 {
			fmt.Println()
		}
	}
	for i := 1; i <= lineCnt2; i++ {
		for j := 1; j <= spaceCnt; j++ {
			fmt.Print(" ")
		}
		for j := 1; j <= starCnt; j++ {
			fmt.Print("*")
		}
		if i < lineCnt2/2 {
			spaceCnt -= 1
			starCnt += 2
			fmt.Println()
		} else if i == lineCnt2/2 {
			fmt.Println()
		} else if i > lineCnt2/2 && i < lineCnt2 {
			spaceCnt += 1
			starCnt -= 2
			fmt.Println()
		} else if i == lineCnt2 {
			fmt.Println()
		}
	}

	for i := 1; i <= lineCnt3; i++ {
		for j := 1; j <= spaceCnt; j++ {
			fmt.Print(" ")
		}
		for j := 1; j <= starCnt; j++ {
			fmt.Print("*")
		}
		if i < lineCnt3/2+1 {
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

//		return a, b, a + b
//	}
// func star(a int) int {
// 	var lineCnt int = a
// 	var spaceCnt int = lineCnt - 1
// 	var starCnt int = 1
// 	if a%2 == 0 {
// 		fmt.Println("도형의 높이가 홀수가 아닙니다")
// 	} else {
// 		for i := 1; i <= lineCnt; i++ {
// 			for j := 1; j <= spaceCnt; j++ {
// 				fmt.Print(" ")
// 			}
// 			for j := 1; j <= starCnt; j++ {
// 				fmt.Print("*")
// 			}
// 			if i < (lineCnt/2)+1 {
// 				spaceCnt -= 1
// 				starCnt += 2
// 				fmt.Println()

// 			} else {
// 				spaceCnt += 1
// 				starCnt -= 2
// 				fmt.Println()

// 			}

// 		}

// 	}

// 	fmt.Println("입력한 도형의 높이")
// 	return a
// }
