package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"github.com/gorilla/mux"
	"regexp"
	"strconv"
)

func JSONWriter(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

type APIServer struct {
	port string
}

// ErrorResponse defines the structure of error responses
type ErrorResponse struct {
	Error string `json:"error"`
}

// ResultResponse defines the structure of successful result responses
type ResultResponse struct {
	Result string `json:"result"`
}

func FuncHandler(w http.ResponseWriter, r *http.Request) {
	// Get input expression from URL query parameter
	input := r.URL.Query().Get("expr")
	input = strings.ReplaceAll(input, " ", "+")
	inputType := defineType(input)
	result:= ""
	var numres float64
	var err error
	if input == "" {
		// Return error if input is missing
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{"Missing input expression"})
		return
	}

	// Evaluate input expression
	switch inputType {
	case "calculator":
		numres,err = calculate(input)
		result = strconv.FormatFloat(numres, 'f', 2, 64)
	case "date":
		result = "date"
	case "addQuestion":
		result = "addQuestion"
	case "deleteQuestion":
		result = "deleteQuestion"
	case "textQuestion":
		result = "textQuestion"
	}

	if err != nil {
		// Return error if input expression is invalid
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{err.Error()})
		return
	}

	// Return successful result
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ResultResponse{result})
}


func defineType(input string) string {
	// define what type is the input
	textQuestionRegex := regexp.MustCompile(`^[A-Za-z0-9\s\?\.,!]+$`)
	calculatorRegex := regexp.MustCompile(`^[\d\+\-\*\/\^\(\)]+$`)
	dateRegex := regexp.MustCompile(`^hariapatanggal(\d{1,2})/(\d{1,2})/(\d{4})$`)

	var questionType string
	switch {
	case calculatorRegex.MatchString(input):
		questionType = "calculator"
	case dateRegex.MatchString(input):
		questionType = "date"
	case textQuestionRegex.MatchString(input):
		questionType = "textQuestion"
	default:
		questionType = "unknown"
	}
	return questionType
}

func NewAPIServer(port string) *APIServer {
	return &APIServer{port: port}
}
func (s *APIServer) Start() {
	// Create router
	router := mux.NewRouter()

	// define what type is the input

	// Add handler for arithmetic expressions
	router.HandleFunc("/input", FuncHandler).Methods("GET")

	// Start server
	log.Fatal(http.ListenAndServe(s.port, router))
}
