package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type CalculationRequest struct {
	Operation string    `json:"operation"`
	Numbers   []float64 `json:"numbers"`
}

type CalculationResponse struct {
	Result float64 `json:"result"`
	Error  string  `json:"error,omitempty"`
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/calculate", calculateHandler).Methods("POST", "OPTIONS")
	
	// Create a new CORS handler
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"}, // Allow requests from React app
		AllowedMethods: []string{"POST", "GET", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Origin", "Accept", "*"},
	})

	// Wrap the router with the CORS handler
	handler := c.Handler(r)

	fmt.Println("Server is running on :8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}

func calculateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		return
	}

	var req CalculationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := performCalculation(req.Operation, req.Numbers)
	
	response := CalculationResponse{Result: result}
	if err != nil {
		response.Error = err.Error()
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func performCalculation(operation string, numbers []float64) (float64, error) {
	switch operation {
	case "add":
		return numbers[0] + numbers[1], nil
	case "subtract":
		return numbers[0] - numbers[1], nil
	case "multiply":
		return numbers[0] * numbers[1], nil
	case "divide":
		if numbers[1] == 0 {
			return 0, fmt.Errorf("division by zero")
		}
		return numbers[0] / numbers[1], nil
	case "power":
		return math.Pow(numbers[0], numbers[1]), nil
	case "sqrt":
		return math.Sqrt(numbers[0]), nil
	case "sin":
		return math.Sin(numbers[0]), nil
	case "cos":
		return math.Cos(numbers[0]), nil
	case "tan":
		return math.Tan(numbers[0]), nil
	default:
		return 0, fmt.Errorf("unknown operation: %s", operation)
	}
}