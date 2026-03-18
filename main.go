package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Recommendation struct {
	PackageName        string `json:"packageName"`
	RecommendedVersion string `json:"recommendedVersion"`
	Strategy           string `json:"strategy"`
}

var demoVersions = map[string]string{
	"curl":            "7.88.1-10+deb12u9",
	"git":             "1:2.39.5-0+deb12u3",
	"ca-certificates": "20230311+deb12u1",
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
		packageName = "curl"
	}

	currentValue := r.URL.Query().Get("packageValue")
	if currentValue == "" {
		currentValue = r.URL.Query().Get("currentValue")
	}

	fmt.Printf("custom request: packageName=%s currentValue=%s\n", packageName, currentValue)

	recommended := demoVersions[packageName]
	if recommended == "" {
		if currentValue != "" {
			recommended = currentValue
		} else {
			recommended = "0.0.0-demo"
		}
	}

	resp := Recommendation{
		PackageName:        packageName,
		RecommendedVersion: recommended,
		Strategy:           "custom-http-demo",
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
	http.HandleFunc("/recommendation/debian", recommendationHandler)
	http.HandleFunc("/healthz", healthHandler)
	log.Println("simple custom datasource demo listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
