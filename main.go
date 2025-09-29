package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
)

type Product struct {
	ProductId    string    `json:"ProductId" xml:"ProductId"`
	Name         string    `json:"Name" xml:"Name"`
	DateSupplied time.Time `json:"DateSupplied" xml:"DateSupplied"`
	Quantity     int       `json:"Quantity" xml:"Quantity"`
	UnitPrice    float64   `json:"UnitPrice" xml:"UnitPrice"`
}

type ProductList struct {
	Products []Product `json:"products" xml:"products"`
}

var productList = []Product{
	{
		ProductId:    "31288741190182539912",
		Name:         "Banana",
		DateSupplied: time.Date(2025, 1, 24, 0, 0, 0, 0, time.UTC),
		Quantity:     124,
		UnitPrice:    0.55,
	},
	{
		ProductId:    "29274582650152771644",
		Name:         "Apple",
		DateSupplied: time.Date(2024, 12, 9, 0, 0, 0, 0, time.UTC),
		Quantity:     18,
		UnitPrice:    1.09,
	},
	{
		ProductId:    "91899274600128155167",
		Name:         "Carrot",
		DateSupplied: time.Date(2025, 3, 31, 0, 0, 0, 0, time.UTC),
		Quantity:     89,
		UnitPrice:    2.99,
	},
	{
		ProductId:    "31288741190182539913",
		Name:         "Banana",
		DateSupplied: time.Date(2025, 2, 13, 0, 0, 0, 0, time.UTC),
		Quantity:     240,
		UnitPrice:    0.65,
	},
}

func sortProducts(products []Product) []Product {
	sortedProducts := make([]Product, len(products))
	copy(sortedProducts, products)

	sort.Slice(sortedProducts, func(i, j int) bool {
		if sortedProducts[i].Name != sortedProducts[j].Name {
			return sortedProducts[i].Name < sortedProducts[j].Name
		}
		return sortedProducts[i].UnitPrice > sortedProducts[j].UnitPrice
	})

	return sortedProducts
}

func printAsJson(products []Product) {
	fmt.Println("Printing in JSON format")
	sortedProducts := sortProducts(products)
	productList := ProductList{Products: sortedProducts}
	jsonData, err := json.MarshalIndent(productList, "", "  ")
	if err != nil {
		fmt.Printf("Error converting to JSON: %v\n", err)
		return
	}
	fmt.Println(string(jsonData))
}

func printAsXml(products []Product) {
	fmt.Println("Printing in XML format")
	sortedProducts := sortProducts(products)
	productList := ProductList{Products: sortedProducts}
	xmlData, err := xml.MarshalIndent(productList, "", "  ")
	if err != nil {
		fmt.Printf("Error converting to XML: %v\n", err)
		return
	}
	fmt.Println(string(xmlData))
}

func printAsCSV(products []Product) {
	fmt.Println("Printing in CSV format")
	sortedProducts := sortProducts(products)
	fmt.Println("ProductId,Name,DateSupplied,Quantity,UnitPrice")
	for _, product := range sortedProducts {
		fmt.Printf("%s,%s,%s,%d,%.2f\n",
			product.ProductId,
			product.Name,
			product.DateSupplied,
			product.Quantity,
			product.UnitPrice,
		)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <format>")
		fmt.Println("Available formats: json, xml, csv")
		return
	}

	format := strings.ToLower(os.Args[1])

	switch format {
	case "json":
		printAsJson(productList)
	case "xml":
		printAsXml(productList)
	case "csv":
		printAsCSV(productList)
	default:
		fmt.Printf("Unknown format: %s\n", format)
		fmt.Println("Available formats: json, xml, csv")
	}
}
