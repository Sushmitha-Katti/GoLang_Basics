package main

import ("fmt")


// Array is not used much in go
func main(){

	fmt.Println("Welcome to array in golangs")

	var fruitList [5]string

	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[3] = "Peach"
	fmt.Println("fruit list is: ", fruitList)
	fmt.Println("Fruit list is: ", len(fruitList))

	var vegList = [3]string{"potato", "beans", "mushroom"}

	fmt.Println("Vegy list is, ", len(vegList))


	
}