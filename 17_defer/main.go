package main

import "fmt"
func main(){
	defer fmt.Println("Hello") //after defer keywod. Execution will be put up in stack and executes next line and then comes back
	fmt.Println("World")
	fmt.Println("Hello2")
	fmt.Println("Hello3")
	myDef()

}

func myDef(){
	for i:=0; i<5; i++ {
		defer fmt.Println(i)
	}
}