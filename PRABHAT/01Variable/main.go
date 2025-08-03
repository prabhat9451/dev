package main

import "fmt"

const LoginToken string = "qwertyuio" //public

func main() {

	var username string = "hitesh"
	fmt.Println(username)
	fmt.Printf("Variables is type of: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variables is type of: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variables is type of: %T \n", smallVal)

	var smallFloat float64 = 255.2578964123322255
	fmt.Println(smallFloat)
	fmt.Printf("Variables is type of: %T \n", smallFloat)

	//default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variables is type of: %T \n", anotherVariable)

	//implicit type
	var website = "learncodeonline.in"
	fmt.Println(website)

	//no var style
	numberOfUser := 30000.0
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("Variables is type of: %T \n", LoginToken)
}
