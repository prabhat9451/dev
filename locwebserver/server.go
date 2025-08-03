package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Handler for GET /
func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Welcome to learncodeonline server")
}

// Handler for GET /get
func getHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Hello from learncodeonline.in",
	})
}

// Handler for POST /post (expects JSON)
func postJSONHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method allowed", http.StatusMethodNotAllowed)
		return
	}
	var data map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

// Handler for POST /postform (expects form data)
func postFormHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method allowed", http.StatusMethodNotAllowed)
		return
	}
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}
	formData := make(map[string]string)
	for key, values := range r.Form {
		formData[key] = values[0] // only first value for each key
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(formData)
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/get", getHandler)
	http.HandleFunc("/post", postJSONHandler)
	http.HandleFunc("/postform", postFormHandler)

	port := ":3000"
	fmt.Println("Server listening at http://localhost" + port)
	log.Fatal(http.ListenAndServe(port, nil))
}
