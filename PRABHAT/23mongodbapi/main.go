package main

import (
	"fmt"
	"net/http"

	"github.com/hiteshchoudhary/mongodbapi/router"
)

func main() {
	fmt.Println("MongoDB API")
	r := router.Router()
	fmt.Println("Server is getting started.....")
	http.ListenAndServe(":6000", r)
	fmt.Println("Listening at port 6000")
}
