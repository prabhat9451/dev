package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://loc.dev"

func main() {
	fmt.Println("LCO web request")
	response, err := http.Get(url)
	checkNilErr(err)

	fmt.Printf("Response is of type: %T\n", response)
	defer response.Body.Close() //caller's responsibility to close the connection

	dataByte, err := ioutil.ReadAll(response.Body)
	checkNilErr(err)

	content := string(dataByte)
	fmt.Println(content)
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
