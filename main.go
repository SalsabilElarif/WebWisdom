package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"wedding-advice/config"
	"wedding-advice/database"
)

func submitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	var advice database.Advice

	// Decode JSON body into advice struct
	err := json.NewDecoder(r.Body).Decode(&advice)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if advice.Message == "" || advice.Name == "" || advice.Relation == "" {
		http.Error(w, "All fields (message, name, relation) are required", http.StatusBadRequest)
		return
	}

	result := database.DB.Create(&advice)
	if result.Error != nil {
		http.Error(w, "failed to save advice", http.StatusInternalServerError)
		return
	}

		
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"message":"Advice received, thank you!"}`))
	
}

func getAllAdviceHandler(w http.ResponseWriter, r *http.Request) {
	var adviceList []database.Advice
	result := database.DB.Order("id DESC").Find(&adviceList)
	if result.Error != nil {
		http.Error(w, "failed to retrieve advice", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(adviceList)
}


func main() {
	// Load up variables
	config.LoadEnv()

	// Connect to database
	database.ConnectDB()

	// Handle requests to the root URL "/"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Wedding Advice API is running")
	})

	http.HandleFunc("/submit", submitHandler)
	http.HandleFunc("/advice", getAllAdviceHandler)

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
