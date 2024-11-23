package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"receipt-processor-challenge/models"
	"receipt-processor-challenge/services"
)

var receipts = make(map[string]models.Receipt)

func ProcessReceipt(w http.ResponseWriter, r *http.Request) {
	var receipt models.Receipt

	// Decode the JSON body
	if err := json.NewDecoder(r.Body).Decode(&receipt); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Generate a unique ID for the receipt
	id := services.GenerateID()

	// Evaluate points for the receipt
	_, err := services.CalculatePoints(receipt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Store the receipt with its ID
	receipts[id] = receipt

	// Respond with the generated ID and calculated points
	response := map[string]string{
		"id": id,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func GetPoints(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	// Get the receipt by ID
	receipt, exists := receipts[id]
	if !exists {
		http.Error(w, "Receipt not found", http.StatusNotFound)
		return
	}

	// Calculate points for the retrieved receipt
	points, err := services.CalculatePoints(receipt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with the calculated points
	response := map[string]int{
		"points": points,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
