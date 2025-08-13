package main

import (
	"encoding/json"
	"fmt"
	"net/http"

)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "I am Mahabub Ahmed.")
}


type Product struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Price float64 `json:"price"`
	ImgURL string `json:"imageUrl"`
}

var (
	productList   []Product
)

func getProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Ahmed")

	if r.Method == "OPTIONS" {
		w.WriteHeader(200)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Please give a GET request", 400)
		return
	}

	encoder := json.NewEncoder(w)
	encoder.Encode(productList)
}


func addProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "OPTIONS" {
		w.WriteHeader(200)
		return
	}

	if r.Method != "POST" {
		http.Error(w, "Method not allowed", 400)
		return
	}

	var newProduct Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "please give a valid json", 400)
		return
	}

	newProduct.ID = len(productList) + 1
	productList = append(productList, newProduct)

	w.WriteHeader(201)
	encoder := json.NewEncoder(w)
	encoder.Encode(newProduct)
}



func main() {
	mux := http.NewServeMux() // router

	mux.HandleFunc("/hello", helloHandler) // route 

	mux.HandleFunc("/about", aboutHandler) // route

	mux.HandleFunc("/products", getProducts)

	mux.HandleFunc("/add-product", addProduct)

	fmt.Println("Server running on : 8080")

	err := http.ListenAndServe(":8080", mux)

	if err != nil {
		fmt.Println("Error staritng the server", err)
	}
}

func init() {
	prd1 := Product{
		ID: 1,
		Title: "Orange",
		Description: "Orange is Orange color. I love orange",
		Price: 100,
		ImgURL: "https://png.pngtree.com/element_our/png/20180903/orange-png-png_75700.jpg",
	}

	prd2 := Product{
		ID: 2,
		Title: "Apple",
		Description: "Apple is red color. I love Apple",
		Price: 100,
		ImgURL: "https://png.pngtree.com/png-vector/20231017/ourmid/pngtree-3d-red-apple-png-with-green-leaf-png-image_10201408.png",
	}

	prd3 := Product{
		ID: 3,
		Title: "Banana",
		Description: "Banana is yellow color. I love Banana",
		Price: 100,
		ImgURL: "https://e7.pngegg.com/pngimages/796/636/png-clipart-banana-banana-thumbnail.png",
	}

/* 	prd4 := Product{
		ID: 4,
		Title: "Mango",
		Description: "Mango is red color. I love mango",
		Price: 100,
		ImgURL: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRUdVvDDvqHyzkbjIC8RlgcNT1d4YPC7mKh7Q&s",
	}

	prd5 := Product{
		ID: 5,
		Title: "cherry",
		Description: "Cherry is red color. I love cherry",
		Price: 100,
		ImgURL: "https://w7.pngwing.com/pngs/689/514/png-transparent-two-cherries-sweet-cherry-black-cherry-sour-cherry-maraschino-cherry-cherry-natural-foods-frutti-di-bosco-food-thumbnail.png",
	}

	prd6 := Product{
		ID: 6,
		Title: "Grapes",
		Description: "Grapes is green color. I love Grapes",
		Price: 100,
		ImgURL: "https://w7.pngwing.com/pngs/997/412/png-transparent-bunch-of-white-grapes-muscat-wine-juice-concord-grape-grape-natural-foods-food-wine-thumbnail.png",
	} */

	productList = append(productList, prd1)
	productList = append(productList, prd2)
	productList = append(productList, prd3)
/* 	productList = append(productList, prd4)
	productList = append(productList, prd5)
	productList = append(productList, prd6) */
}

/* 
		[ ] = list
		{ } = object
		JSON = JavaScript Object Notation
*/