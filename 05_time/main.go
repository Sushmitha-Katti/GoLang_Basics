package main

import (
	"fmt" 
	"time"
)

func main() {

	fmt.Println("welcome to time study of golang")

	presentTime := time.Now()

	// for format you have to give this contant time and time given in doc
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdate := time.Date(2020, time.August, 10, 23, 23, 0, 0, time.UTC)

	fmt.Println(createdate)
	fmt.Printf(createdate.Format("01-02-2006 Monday\n"))

}