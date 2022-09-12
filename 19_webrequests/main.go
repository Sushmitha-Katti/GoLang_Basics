package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("LCO web request")
	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	fmt.Printf("Response if of type: %T\n", response)

	defer response.Body.Close() //caller's responsibilty to close the connection
	
	databytest, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	content := string(databytest)
	fmt.Println(content)
}