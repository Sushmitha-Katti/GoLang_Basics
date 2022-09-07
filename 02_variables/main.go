package main

import "fmt"

// Variable starting with Capital letter means it is public else private
const LoginToken string = "gibberish"

func main() {
	var username bool = false
	fmt.Println("username")
	fmt.Printf("Variable is of type: %T \n", username)

	var smallVal uint8 = 255 
	fmt.Println((smallVal))
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float32 = 255.44565667678 //prints only 4-5 characters
	fmt.Println((smallFloat))
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	var bigFloat float64 = 255.44565667678
	fmt.Println((bigFloat))
	fmt.Printf("Variable is of type: %T \n", bigFloat)


	var anotherVarible int
	fmt.Println((anotherVarible))
	fmt.Printf("Variable is of type: %T \n", anotherVarible)


	//implicit type
	var website = "learncodeonline.in"
	fmt.Println((website))

	//no var style - allowed only inside methods not as global variable
	numberofUser := 3000
	fmt.Println((numberofUser))
}
