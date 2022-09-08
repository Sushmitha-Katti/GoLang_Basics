package main

import ("fmt")

func main(){
	fmt.Println("Struct Class")
	// no inheritence in goland: No super or parent

	sushmitha := User{"Sushmitha", "skattti@gmail.com", true, 16}
	fmt.Println((sushmitha))
	fmt.Printf("Sushmitha details are: %+v\n", sushmitha)
	fmt.Printf("Name is %v and email is %v\n", sushmitha.Name, sushmitha.Email)

	sushmitha.GetStatus()
	sushmitha.NewMail()

}
type User struct {
	Name string //capital letters are for public
	Email string
	Status bool
	Age int
}

func (u User) GetStatus(){
	fmt.Println("Is User Active: ", u.Status)

}
func (u User) NewMail(){
	u.Email = "abc@hmail"
	fmt.Println("New EMail: ",  u.Email)
}