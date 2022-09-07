package main
import ("fmt")

// pointers : the direct refernce to memory address

func main() {
	fmt.Println("Welcome to a class on pointers")

	var ptr *int // poiner def
	fmt.Println("value of pointer is ", ptr)

	myNumber := 23
	var ptr2 = &myNumber

	// prints addresss
	fmt.Println("value of pointer2 is", ptr2)

	// prints value(23)
	fmt.Println("value of pointer2 is", *ptr2)
	
	// operation will be performed on values
	*ptr2 = *ptr2 + 2
	fmt.Println("New value is: ", myNumber) // myNumner = 25

}