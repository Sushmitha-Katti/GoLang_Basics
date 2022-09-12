package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("HTTP  Json Resquest")
	// PerformGetRequest()
	// PerformPostJsonRequest()
	PerformPostFormRequest()
}

func PerformGetRequest(){
	const myurl = "http://localhost:8000/get"

	response, err := http.Get(myurl)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status Code: response.StatusCode")

	fmt.Println("Content Length is: ", response.ContentLength)

	content, _ := ioutil.ReadAll(response.Body)

	var responseString strings.Builder
	byteCount, _ := responseString.Write(content)

	fmt.Println("ByteCode is:", byteCount)
	fmt.Println(responseString.String())

	// fmt.Println(string(content))

}


func PerformPostJsonRequest() {

	const myurl = "http://localhost:8000/post"
	
	requestBody := strings.NewReader(`
		{
			"coursename" :"go lang",
			"price": 0,
			"platforms": "learncode.io"
		}
	`)

	response, err := http.Post(myurl, "application/json", requestBody)
	
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))


}

func PerformPostFormRequest(){
	const myurl = "http://localhost:8000/postform"

	// formdata
	
	data := url.Values{}

	data.Add("firstname", "hitesh")

	data.Add("lastname", "choudary")

	data.Add("email", "email@e.com")

	response, err := http.PostForm(myurl, data)

	if err != nil {
		panic(err)
	}

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))
}