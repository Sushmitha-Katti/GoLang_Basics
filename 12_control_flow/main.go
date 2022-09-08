package main
import ("fmt")

func main(){
	fmt.Println("Control Flow")

	loginCount := 10
	var result string

	if loginCount < 10 {
		result = "Regulare user"
	}else if loginCount > 10 {
		result = "Watch Out"
	} else {
		result = "Exactly 10"
	}
	fmt.Println((result))

	if 9%2 == 0{
		fmt.Println("Number is even")
	}else {
		fmt.Println("Number is odd")
	}

	if num := 3; num < 10 {
		fmt.Println("num is less than 10")
	} else {
		fmt.Println("Num is not less tha 10")
	}

	
}