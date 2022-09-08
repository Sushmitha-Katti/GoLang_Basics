package main

import "fmt"

func main(){
	fmt.Println("Welcome to functions in golang")	
	greeter()
	result := add(3,4)
	fmt.Println("Result is: ", result)
	prosRes := proAdder(2,3,4,5)
	fmt.Println("Pro Result is: ", prosRes)

	// function inside a function is not allowed
}

func greeter(){
	fmt.Println("Namstey from golang")
}

// Annonymous function
// func () {
// 	fmt.Printn("Anonymous")
// }()

func add (val1 int, val2 int) int{
	return val1 + val2
}

func proAdder(values ...int)int {
	total := 0
	for _, val := range values{
		total += val
	}
	return total
}