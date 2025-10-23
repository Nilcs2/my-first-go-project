package main

import "fmt"

type Product struct {
	ID    int
	Name  string
	Price float64
}

func main() {
	products := []Product{
		{ID: 1, Name: "Laptop", Price: 1499.99},
		{ID: 2, Name: "Smartphone", Price: 1099.49},
		{ID: 3, Name: "Tablet", Price: 699.00},
	}
	fmt.Println("Product Catalogue: ")
	for _, p := range products {
		fmt.Printf("ID: %d, Name: %s, Price: %.2f\n", p.ID, p.Name, p.Price)
	}
	var id int
	var name string
	var price float64
	fmt.Println("\n Add a new product to the catalogue: ")
	fmt.Println("Enter Product ID: ")
	fmt.Scanln(&id)
	fmt.Println("Enter Product Name: ")
	fmt.Scanln(&name)
	fmt.Println("Enter Product Price: ")
	fmt.Scanln(&price)
	newProduct := Product{ID: id, Name: name, Price: price}
	products = append(products, newProduct)
	fmt.Println("\nUpdated Product Catalogue: ")
	for _, p := range products {
		fmt.Printf("ID: %d, Name: %s, Price: $%.2f\n", p.ID, p.Name, p.Price)
	}
}
