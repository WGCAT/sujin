package main

import "fmt"

func main() {
	var mathScore int = 90
	var engScore int = 70

	var totalScore = mathScore + engScore
	var avgScore = totalScore / 2

	fmt.Println(totalScore)

	fmt.Println(avgScore)
}
