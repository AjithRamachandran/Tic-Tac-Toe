package main

import ("fmt"
	"math/rand"
	"os"
	"os/exec"
)

const empty string = "-"
var ttt[3][3] string

func checkHorizontal() bool {
	if ((ttt[0][0] == ttt[0][1]) && (ttt[0][1] == ttt[0][2]) && (ttt[0][1] != empty)) || ((ttt[1][0] == ttt[1][1]) && (ttt[1][1] == ttt[1][2]) && (ttt[1][1] != empty)) || ((ttt[2][0] == ttt[2][1]) && (ttt[2][1] == ttt[2][2]) && (ttt[2][2] != empty)) {
		return true
	}
	return false
}

func checkVertical() bool {
	if ((ttt[0][0] == ttt[1][0]) && (ttt[1][0] == ttt[2][0]) && (ttt[0][0] != empty)) || ((ttt[0][1] == ttt[1][1]) && (ttt[1][1] == ttt[2][1]) && (ttt[0][1] != empty)) || ((ttt[0][2] == ttt[1][2]) && (ttt[1][2] == ttt[2][2]) && (ttt[0][2] != empty)) {
		return true
	}
	return false
}

func checkDiagonal() bool {
	if ((ttt[0][0] == ttt[1][1]) && (ttt[1][1] == ttt[2][2]) && (ttt[0][0] != empty)) || ((ttt[0][2] == ttt[1][1]) && (ttt[1][1] == ttt[2][0]) && (ttt[1][1] != empty)) {
		return true
	}
	return false
}

func checkWin() bool {
	if checkHorizontal() || checkVertical() || checkDiagonal() {
		return true
	}
	return false
}

func showVector() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
	for t:=0; t<3; t++ {
	fmt.Println(ttt[t][0], " | ", ttt[t][1], " | ", ttt[t][2])
	}
}

func user() {
	var userx, usery int
	fmt.Println("Enter the coordinate to put 'X'")
	fmt.Scan(&userx, &usery)
	if (userx < 0 || userx > 2) || (usery < 0 || usery > 2) {
		fmt.Println("Please enter a valid input!!")
		user()
	} else if ttt[userx][usery] != empty {
		fmt.Println("Cell is not empty!")
		user()
	} else {
		ttt[userx][usery] = "X"
		if checkWin() {
			showVector()
			fmt.Println("You won!")
			os.Exit(0)
		} else {
			showVector()
		}
	}
}

func comp() {
	var compx= rand.Intn(3)
	var compy= rand.Intn(3)
	if ttt[compx][compy] != empty {
		comp()
	} else {
		ttt[compx][compy] = "O"
		if checkWin() {
			showVector()
			fmt.Println("Computer won!")
			os.Exit(0)
		} else {
			showVector()
		}
	}
}

func main() {

	count := 0
	for i:=0; i<3; i++ {
		for j:=0; j<3; j++ {
			ttt[i][j] = empty
		}
	}
	showVector()
	for true {
		if count%2==0 {
			fmt.Println("Your Turn")
			user()
			count++
		} else {
			fmt.Println("Computer's Turn")
			comp()
			count++
		}
	}
}
