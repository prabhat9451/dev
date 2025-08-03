package main

import "fmt"

func main() {
	fmt.Println("Loop and break in golang")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}

	fmt.Println(days)

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for index, day := range days {
	// 	fmt.Printf("Index is %v and value is %v\n", index, day)
	// }

	// for _, day := range days {
	// 	fmt.Printf("Index is and value is %v\n", day)
	// }

	rougueValue := 1
	for rougueValue < 10 {

		if rougueValue == 2 {
			goto loc
		}

		if rougueValue == 5 {
			rougueValue++
			// break
			continue
		}
		fmt.Println("Value is: ", rougueValue)
		rougueValue++
	}

loc:
	fmt.Println("Jumping at LearnCodenonline.in")
}
