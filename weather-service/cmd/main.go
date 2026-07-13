package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type LocaleOutput struct {
	Localidade string  `json:"city"`
	TempC      float64 `json:"temp_c"`
	TempF      float64 `json:"temp_f"`
	TempK      float64 `json:"temp_k"`
}

func weatherHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	response := LocaleOutput{
		Localidade: "São Paulo",
		TempC:      25.5,
		TempF:      77.9,
		TempK:      298.65,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/weather", weatherHandler)
	log.Println("Weather service running on :8181")
	log.Fatal(http.ListenAndServe(":8181", nil))
}
