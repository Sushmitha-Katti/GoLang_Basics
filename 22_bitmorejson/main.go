package main

import (
	"encoding/json"
	"fmt"
)

type course struct{
	Name string `json:"coursename"` //alias name
	Price int
	Platforms string `json:"website"` 
	Password string `json:"-"`
	Tags []string `json:"tags,omitempty"` //if nil then it doesn't show up
}

func main() {
	fmt.Println("Welcome to Json video")
	// EncodeJson()
	DecodeJson()
}

func EncodeJson(){
	lcoCourses := []course{
		{"ReactJs BootCamp", 299, "learncodeonline.in", "abc123", []string{"web-dev", "js"} },
		{"MERN BootCoam", 199, "learncodeonline.in", "bcd123", []string{"web-dev", "js"} },

		{"Angulary BootCoam", 299, "learncodeonline.in", "abc123", nil },

	}

	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)
}

func DecodeJson(){

	jsonDataFormWeb := []byte(`
	{
			"coursename": "ReactJs BootCamp",
			"Price": 299,
			"website": "learncodeonline.in",
			"tags": [
					"web-dev",
					"js"
			]
	}
	`)

	var lcoCourse course 

	checkValid := json.Valid(jsonDataFormWeb)

	if checkValid{
		fmt.Println("Json was valid")
		json.Unmarshal(jsonDataFormWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)

	}else {
		fmt.Println("Invalid Json")
	}

	//some cases where you just want to add data to key value

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFormWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and Type is %T\n", k, v, v )
	}



}