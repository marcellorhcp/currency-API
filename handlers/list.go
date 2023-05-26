package handlers

import (
	"currency-API/models"
	"encoding/json"
	"log"
	"net/http"
)

func List(w http.ResponseWriter, r *http.Request) {
	currencies, err := models.GetAll()
	if err != nil {
		log.Printf("Erro ao obter os registros: %v", err)
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(currencies)
}
