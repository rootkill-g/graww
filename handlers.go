package main

import (
	"encoding/json"
	"net/http"
)

func handleProductInfo(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		GetProduct(w, r)
	case http.MethodPost:
		AddNewProduct(w, r)
	case http.MethodPatch:
		UpdateProductInfo(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	var productId = r.URL.Query().Get("productId")
	productInfo, ok := database[productId]
	if !ok || productId == "" {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	response := ProductInfo{
		Id:          productId,
		Name:        productInfo.Name,
		Price:       productInfo.Price,
		Description: productInfo.Description,
	}
	json.NewEncoder(w).Encode(response)
}

func AddNewProduct(w http.ResponseWriter, r *http.Request) {
	var productId = r.URL.Query().Get("productId")
	// Decode the JSON payload directly into the struct
	var payloadData ProductInfo
	if err := json.NewDecoder(r.Body).Decode(&payloadData); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	database[productId] = payloadData
	w.WriteHeader(http.StatusCreated)
}

func UpdateProductInfo(w http.ResponseWriter, r *http.Request) {
	var productId = r.URL.Query().Get("productId")
	productInfo, ok := database[productId]
	if !ok || productId == "" {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}
	// Decode the JSON payload directly into the struct
	var payloadData ProductInfo
	if err := json.NewDecoder(r.Body).Decode(&payloadData); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	// Update the product information
	productInfo.Id = productId
	productInfo.Name = payloadData.Name
	productInfo.Price = payloadData.Price
	productInfo.Description = payloadData.Description
	database[productInfo.Id] = productInfo
	w.WriteHeader(http.StatusOK)
}
