package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	//no inheritance inn golang; No super or parent

	prabhat := User{"Prabhat", "prabhat123@gmail.com", true, 18}
	fmt.Println(prabhat)
	fmt.Printf("Pabhat details are : %+v\n", prabhat)
	fmt.Printf("Name is %v and email is %v.", prabhat.Name, prabhat.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
