package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welocme to video on slices")


	// We don't declare size here: only in array we do that
	var fruitList = []string{"Apple", "Tomoto", "Peach"}
	fmt.Printf("type of fruitList is %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)

	highScores := make([]int, 4)
	highScores[0] = 234
	highScores[1] = 235
	highScores[2] = 236
	highScores[3] = 237

	highScores = append(highScores, 555, 666, 321)

	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))
	sort.Ints(highScores)
	fmt.Println(highScores)
	




}
