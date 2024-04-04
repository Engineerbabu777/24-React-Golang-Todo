package main

import (
	"fmt"
	"log"
	"net/http"
	"react-go-todo/router"
)

func main() {

	r := router.Router()

	fmt.Println("Starting server on port 9000...")

	log.Fatal(http.ListenAndServe(":8080", r))
}
