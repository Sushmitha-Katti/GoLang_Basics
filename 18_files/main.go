package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main(){
	fmt.Println("Files in Go")
	content := "This needs to go in a file - LearnCodeOnline.in"

	file, err := os.Create("./mylcogofile.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)

	if err != nil {
		panic(err)
	}
	fmt.Println("Length is: ", length)
	defer file.Close()
	readFile("./mylcogofile.txt")

}

func readFile(filename string){
	databytes, err := ioutil.ReadFile(filename)
	checkkNilErr(err)

	fmt.Println(string(databytes))
	
}

func checkkNilErr(err error){
	if err != nil {
		panic(err)
	}
}