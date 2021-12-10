package main

import (
	"fmt"
	"github.com/ahmedkhaeld/go-web/pkg/handlers"
	"net/http"
)

const portNumber = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting the application %v", portNumber))

	_ = http.ListenAndServe(portNumber, nil)
}
