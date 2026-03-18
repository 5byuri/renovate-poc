package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Recommendation struct {
	RecommendedVersion string `json:"recommendedVersion"`
}

var demoRecommendations = map[string]string{
	"next":      "9.6.5",
	"node":      "20.11.1",
	"react":     "18.7.0",
	"react-dom": "18.3.0",
}

func recommendationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	packageName := r.URL.Query().Get("packageName")
	if packageName == "" {
		packageName = r.URL.Query().Get("depName")
	}
	if packageName == "" {
		http.Error(w, "missing packageName", http.StatusBadRequest)
		return
	}

	currentValue := r.URL.Query().Get("packageValue")
	if currentValue == "" {
		currentValue = r.URL.Query().Get("currentValue")
	}
	if currentValue == "" {
		http.Error(w, "missing packageValue/currentValue", http.StatusBadRequest)
		return
	}

	fmt.Printf("custom request: packageName=%s currentValue=%s\n", packageName, currentValue)

	recommendedVersion := demoRecommendations[packageName]
	if recommendedVersion == "" {
		recommendedVersion = currentValue
	}

	resp := Recommendation{
		RecommendedVersion: recommendedVersion,
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(resp)
}

func healthHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("ok\n"))
}

func main() {
	http.HandleFunc("/recommendation", recommendationHandler)
	http.HandleFunc("/healthz", healthHandler)
	log.Println("simple custom datasource demo listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
