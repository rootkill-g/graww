package main

import (
	"log"
	"net/http"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

var middlewares = []Middleware{
	TokenAuthMiddleware,
}

func main() {
	var handler http.HandlerFunc = handleProductInfo

	for _, middleware := range middlewares {
		handler = middleware(handler)
	}

	http.HandleFunc("/product/info", handler)

	log.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
