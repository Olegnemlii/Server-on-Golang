package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Founder struct {
	Product     string `json:"product`
	Price       int    `json:"price"`
	Quantity    int    `json:"quantity"`
	TotalAmount int    `json:"total_amount"`
}

var founders []Founder

func greetingsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Сервер запущен")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var founder Founder
	json.NewDecoder(r.Body).Decode(&founder)

	founder.TotalAmount = founder.Price * founder.Quantity
	founders = append(founders, founder)

	json.NewEncoder(w).Encode(founders)

}

func main() {
	r := mux.NewRouter()

	founders = append(founders, Founder{Product: "Sour candies", Price: 11, Quantity: 3, TotalAmount: 0})

	r.HandleFunc("/", greetingsHandler).Methods("GET")
	r.HandleFunc("/form", formHandler).Methods("POST")

	fmt.Println("Welcome to the candy shop!")
	fmt.Println("Server start, port 8080\n")
	log.Fatal(http.ListenAndServe(":8080", r))
}
