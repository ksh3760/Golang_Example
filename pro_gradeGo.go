package main

import (
	"fmt"
)

var gGrdArr [5][5]string // save students grade [id, name, korSco, mathSco, engSco]
var gCnt int = 0         // counting gGrdArr

func main() {
	menu()
}

func menu() { // The function that

	var flag bool = true
	var menuSelector int // menu selector

	for flag {

		fmt.Println("============ [MENU] ============")
		fmt.Println("1)Insert grades \n2)Print grades \n3)Exit")
		fmt.Println("================================")

		menuSelector = 0

		fmt.Print("select menu : ")
		fmt.Scanln(&menuSelector)

		switch menuSelector {
		case 1:
			fmt.Println("1)Insert grades")
			fmt.Println("--------------------------------")
			grade()
			break
		case 2:
			fmt.Println("2)Print grades")
			fmt.Println("--------------------------------")
			outGrade()
			break
		case 3:
			fmt.Println("3)Exit")
			fmt.Println("--------------------------------")
			fmt.Println("Thank you")
			flag = false
			break

		default:
			fmt.Println("default")
		}
	}

} // end func menu

func grade() { // The function that insert student's grades from user

	var numId, name, korSco, mathSco, engSco string

	fmt.Print("insert numId : ")
	fmt.Scan(&numId)

	fmt.Print("insert name : ")
	fmt.Scan(&name)

	fmt.Print("insert Korean score  : ")
	fmt.Scan(&korSco)

	fmt.Print("insert math score  : ")
	fmt.Scan(&mathSco)

	fmt.Print("insert English score  : ")
	fmt.Scan(&engSco)

	gGrdArr[gCnt][0] = numId
	gGrdArr[gCnt][1] = name
	gGrdArr[gCnt][2] = korSco
	gGrdArr[gCnt][3] = mathSco
	gGrdArr[gCnt][4] = engSco

	gCnt += 1

	fmt.Println("------------------------------------------------")
	fmt.Println(numId, "\t|", name, "\t|", korSco, "\t|", mathSco, "\t|", engSco)
	fmt.Println("------------------------------------------------")

} // end func grade

func outGrade() { // The function that prints student's grades

	fmt.Println("-----------------------------------------------")
	fmt.Println(" numId | name | KorScore | mathScore | EngScore")
	fmt.Println("-----------------------------------------------")

	for i := 0; i < gCnt; i++ {
		for j := 0; j < 5; j++ {
			fmt.Print(gGrdArr[i][j], "\t|")
		}
		fmt.Println()
		fmt.Println("-----------------------------------------------")
	}

} // end func outGrade
