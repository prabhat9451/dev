package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang")
	content := "This needs to go in a files - LearnCodeOnline.in"

	file, err := os.Create("./mylcogofile.txt")

	// if err != nil {
	// 	panic(err)
	// }
	checkNilErr(err)

	length, err := io.WriteString(file, content)
	checkNilErr(err)

	fmt.Println("Length is : ", length)
	defer file.Close()
	readFile("./mylcogofile.txt")
}

func readFile(fileName string) {
	dataByte, err := ioutil.ReadFile(fileName)
	checkNilErr(err)
	fmt.Println("Text data inside the file is \n", dataByte)
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
