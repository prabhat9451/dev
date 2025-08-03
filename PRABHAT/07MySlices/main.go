package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to video on slices")

	var fruitList = []string{"Apple", "Orange", "Banana", "Mango"}
	fmt.Printf("Type of fruitlist is %T\n", fruitList)

	fruitList = append(fruitList, "Grapes", "Lime")
	fmt.Println(fruitList)

	// fruitList = append(fruitList[1:3])
	// fmt.Println(fruitList)

	fruitList = append(fruitList[:3])
	fmt.Println(fruitList)

	highScore := make([]int, 4)
	highScore[0] = 241
	highScore[1] = 342
	highScore[2] = 834
	highScore[3] = 293
	// highScore[4] = 777

	highScore = append(highScore, 555, 666, 777)
	fmt.Println(highScore)
	fmt.Println(sort.IntsAreSorted(highScore))

	sort.Ints(highScore)
	fmt.Println(highScore)
	fmt.Println(sort.IntsAreSorted(highScore))

	//how to remove a value from slices based on index
	var courses = []string{"reactjs", "javascript", "java", "python", "c++"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
