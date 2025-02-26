package main

// Schema for product information
type ProductInfo struct {
	Id          string
	Name        string
	Price       string
	Description string
}

// In-memory database
var database = map[string]ProductInfo{
	"Light Phone III": {
		Id:          "LP1",
		Name:        "Light Phone III",
		Price:       "$599",
		Description: "A phone designed to give you the tools to flourish as the most thoughtful & intentional version of yourself.",
	},
	"Fairphone 5": {
		Id:          "F5",
		Name:        "Fairphone 5",
		Price:       "$599",
		Description: "The smartphone that lasts a long time",
	},
}
