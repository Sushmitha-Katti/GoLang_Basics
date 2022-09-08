package main

import ("fmt")

func main(){

	fmt.Println("Welcome to video on maps")

	languages := make(map[string]string)

	languages["JS"] = "JavaScript"
	languages["RB"] = "ruby"
	languages["PY"] = "Python"

	fmt.Println("List of langauages: ", languages)
	fmt.Println("Js shorts for: ", languages["JS"])

	// delete
	delete(languages, "RB")
	fmt.Println("List of languages: ", languages)

	// loops

	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}

	// comma ok syntax
	for _, value := range languages {
		fmt.Printf("For key v, value is %v\n", value)
	}


}