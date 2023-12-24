package main

import (
	"fmt"
)

func main() {

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
		} else if i > (lineCnt1/2) && i < lineCnt1 {
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
		if i < (lineCnt2 / 2) {
			spaceCnt -= 1
			starCnt += 2
			fmt.Println()
		} else if i == (lineCnt2 / 2) {
			fmt.Println()
		} else if i > (lineCnt2/2) && i < lineCnt2 {
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

