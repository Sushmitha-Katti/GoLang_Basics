package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?sdfdfdgf"

func main() {
	fmt.Println("Welcome to handling URLS in golang")
	fmt.Println("Welcome to handling URLS in goloang")

	//parsing
	result, _ := url.Parse(myurl)

	fmt.Println(result.Scheme)

	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)
	fmt.Println(result.Path)
	fmt.Println(result.Host)

	qparams := result.Query()
	fmt.Printf("The type of query params are: %T\n", qparams)

	for _ , val := range qparams {
		fmt.Println("Param is:", val)
	}


	partsOfUrl := &url.URL{
		Scheme: "https",
		Host: "lco.dev",
		Path: "/tutess",
		RawPath : "user-hitesh",
	}

	fmt.Println(qparams["coursename"])
	fmt.Println(partsOfUrl.String())




	

}

