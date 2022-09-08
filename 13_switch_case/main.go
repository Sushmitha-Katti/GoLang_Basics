package main

import (
	"math/rand"
	"fmt"
	"time"
)

func main() {
	fmt.Println("Switch in go lang")
	rand.Seed(time.Now().UnixNano())

	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value of dice is ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice Value is 1 and you can open")
	
	case 2:
		fmt.Println("you can move 2 spot")
	case 3:
		fmt.Println("3 spot")
		fallthrough // goes to next case as well
	case 4:
		fmt.Println("move 4 spot")
	case 5:
		fmt.Println("move 5 spot")
	case 6:
		fmt.Println("move 6 and roll dice again")
	}



}