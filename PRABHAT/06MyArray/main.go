package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golangs")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Orange"
	fruitList[3] = "Mango"

	fmt.Println("Fruit list is : ", fruitList)
	fmt.Println("Fruit list is : ", len(fruitList))

	var vegList = [5]string{"Potato", "Beans", "Tomato"}
	fmt.Println("Vegy list is : ", vegList)
	fmt.Println("Vegy list is : ", len(vegList))
}
