package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating of our Pizza:")

	//comma ok || err ok 
	//In go there is no try catch : 
	// errors will be treated as variables
	input, _ := reader.ReadString('\n') //_ is variables
	fmt.Println("Thanks for rating, ", input)
	fmt.Printf("Type of this rating is %T \n", input)

}
