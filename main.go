package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"receipt-processor-challenge/handlers"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/receipts/process", handlers.ProcessReceipt).Methods("POST")
	router.HandleFunc("/receipts/{id}/points", handlers.GetPoints).Methods("GET")

	http.ListenAndServe(":10000", router)
}
