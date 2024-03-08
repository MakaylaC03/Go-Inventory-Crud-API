package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Item struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Price    string `json:"price"`
	Brand    *Brand `json:"brand"`
	Date     *Date  `json:"date"`
	Quantity int    `json:"quantity"`
}

type Brand struct {
	BrandName string `json:"brandname"`
}

type Date struct {
	DateBought string `json:"datebought"`
	DateListed string `json:"datelisted"`
	DateSold   string `json:"datesold"`
}

var Inventory []Item

func getItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Inventory)
}

func deleteItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range Inventory {
		if item.ID == params["id"] {
			Inventory = append(Inventory[:index], Inventory[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(Inventory)
}

func getItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range Inventory {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var item Item
	_ = json.NewDecoder(r.Body).Decode(&item)
	item.ID = strconv.Itoa(rand.Intn(100000000))
	Inventory = append(Inventory, item)
	json.NewEncoder(w).Encode(Inventory)
}

func updateItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range Inventory {
		if item.ID == params["id"] {
			Inventory = append(Inventory[:index], Inventory[index+1:]...)
			var item Item
			_ = json.NewDecoder(r.Body).Decode(&item)
			item.ID = params["id"]
			Inventory = append(Inventory, item)
			json.NewEncoder(w).Encode(Inventory)
			return
		}
	}
}

func main() {
	r := mux.NewRouter()

	Inventory = append(Inventory, Item{ID: "1", Name: "Harlow Coffe Table", Price: "$300", Brand: &Brand{BrandName: "Bernhardt"}, Date: &Date{DateBought: "Unknown", DateListed: "2/9/24", DateSold: "N/A"}, Quantity: 1})
	Inventory = append(Inventory, Item{ID: "2", Name: "STORNAS", Price: "$500", Brand: &Brand{BrandName: "Ikea"}, Date: &Date{DateBought: "Unknown", DateListed: "2/7/24", DateSold: "N/A"}, Quantity: 1})

	r.HandleFunc("/items", getItems).Methods("GET")
	r.HandleFunc("/items/{id}", getItem).Methods("GET")
	r.HandleFunc("/items", createItem).Methods("POST")
	r.HandleFunc("/items/{id}", updateItem).Methods("PUT")
	r.HandleFunc("/items/{id}", deleteItem).Methods("DELETE")

	fmt.Printf("Starting server at port 8000 \n")
	log.Fatal(http.ListenAndServe(":8000", r))

}
