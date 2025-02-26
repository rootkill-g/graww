package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/product/info", handleProductInfo)

	log.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
