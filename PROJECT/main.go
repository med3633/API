package main

import (
	"encoding/json" // Import the package for JSON encoding and decoding
	"fmt"           // Import the package for formatted I/O
	"log"           // Import the package for logging
	"net/http"      // Import the package for HTTP client and server implementations
)

// Product represents a product with ID, Name, Quantity, and Price.
type Product struct {
	Id       string  // Product ID (as a string)
	Name     string  // Product Name
	Quantity int     // Quantity available
	Price    float64 // Price of the product
}

var Products []Product // Slice to hold a list of products

// homepage handles the root URL and returns a welcome message.
func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to this project of product") // Write a welcome message to the response writer
	fmt.Println("Server running on port 8080")           // Log a message indicating the server is running
}

// getProduct returns a specific product based on its ID.
func getProduct(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint hit: getProduct")             // Log that the function is called
	key := r.URL.Path[len("/product/"):]                // Extract the ID from the URL path
	for _, product := range Products {                  // Iterate through the list of products
		if product.Id == key {                           // If the product ID matches the key
			json.NewEncoder(w).Encode(product)          // Encode the product as JSON and write it to the response
			return                                      // Return to avoid further processing
		}
	}
	http.NotFound(w, r)                                 // Return a 404 Not Found if no product matches
}

// returnAllProducts returns all the products as a JSON response.
func returnAllProducts(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint hit: returnAllProducts")      // Log that the function is returning all products
	json.NewEncoder(w).Encode(Products)                 // Encode the Products slice as JSON and write it to the response writer
}

// handleRequest sets up the routes and starts the server.
func handleRequest() {
	http.HandleFunc("/product/", getProduct)            // Map the /product/ endpoint to the getProduct handler
	http.HandleFunc("/products", returnAllProducts)     // Map the /products endpoint to the returnAllProducts handler
	http.HandleFunc("/", homepage)                      // Map the root endpoint to the homepage handler
	http.ListenAndServe(":8080", nil)                   // Start the server on port 8080
}

func main() {
	Products = []Product{                               // Initialize the Products slice with sample data
		{Id: "1", Name: "Chair", Quantity: 100, Price: 100.00}, // First product
		{Id: "2", Name: "Desk", Quantity: 200, Price: 200.00},  // Second product
	}
	handleRequest()                                     // Call the function to handle requests and start the server
}
