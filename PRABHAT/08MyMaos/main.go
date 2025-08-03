package main

import "fmt"

func main() {
	fmt.Println("Maps in golang")

	languages := make(map[string]string)

	languages["JS"] = "JavaScript"
	languages["PY"] = "Python"
	languages["RJS"] = "ReactJavaScript"

	fmt.Println("List of all languages : ", languages)
	fmt.Println("JS shorts for : ", languages["JS"])

	delete(languages, "RJS")
	fmt.Println("List of all languages : ", languages)

	//loops are interesting in golang
	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}

	for _, value := range languages {
		fmt.Printf("For key v, value is %v\n", value)
	}
}
