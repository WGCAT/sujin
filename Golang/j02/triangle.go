package main

import "fmt"

func main() {

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
		spaceCnt -= 1
		starCnt += 2
		fmt.Println()
	}

}
