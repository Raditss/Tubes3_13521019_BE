package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
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
	// Parse request body
	var requestBody struct {
		Expr string `json:"expr"`
		Alg  string `json:"alg"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{err.Error()})
		return
	}
	input := requestBody.Expr
	parts := strings.Split(input, "?")
	Alg := requestBody.Alg
	var err error
	result := ""
	var resultarr = make([]string, len(parts))
	for i := 0; i < len(parts)-1; i++ {
		input = strings.ToLower(parts[i])
		inputType := defineType(input)
		var numres float64
		fmt.Println(inputType)
		switch inputType {
		case "calculator":
			numres, err = calculate(input)
			resultarr[i] = strconv.FormatFloat(numres, 'f', 2, 64)
			fmt.Println(resultarr[i])
		case "date":
			resultarr[i] = getDayOfWeek(input)
			fmt.Println(resultarr[i])
		case "textQuestion":
			qtype := defineQuestionType(input)
			switch {
			case qtype == "add":
				resultarr[i] = addQuestionToDB(input)
				fmt.Println(resultarr[i])
			case qtype == "del":
				resultarr[i] = delQuestion(input)
				fmt.Println(resultarr[i])
			case qtype == "question":
				if Alg == "KMP" {
					resultarr[i] = getAnswerKMP(input)
					fmt.Println(resultarr[i])
				} else if Alg == "BM" {
					resultarr[i] = getAnswerBM(input)
					fmt.Println(resultarr[i])
				}
			default:
				resultarr[i] = ""
			}
		}
	}

	if err != nil {
		// Return error if input expression is invalid
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{err.Error()})
		return
	}
	result= strings.Join(resultarr, " ; ")

	// Return successful result
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ResultResponse{result})

}

func defineType(input string) string {
	// define what type is the input
	// textQuestionRegex := regexp.MustCompile(`^[A-Za-z0-9\s\?\.,!]+$`)
	calculatorRegex := regexp.MustCompile(`^[\d\+\-\*\/\^\(\)]+$`)
	dateRegex := regexp.MustCompile(`^hari apa tanggal (\d{1,2})/(\d{1,2})/(\d{4})$`)

	var questionType string
	switch {
	case calculatorRegex.MatchString(input):
		questionType = "calculator"
	case dateRegex.MatchString(input):
		questionType = "date"
	// case textQuestionRegex.MatchString(input):
	// 	questionType = "textQuestion"
	default:
		questionType = "textQuestion"
	}
	return questionType
}

func NewAPIServer(port string) *APIServer {
	return &APIServer{port: port}
}

func (s *APIServer) Start() {
	// Create router
	router := mux.NewRouter()

	// Add handler for arithmetic expressions
	router.HandleFunc("/api/question", FuncHandler).Methods("POST")


	// Add CORS middleware to router
	handler := corsMiddleware(router)

	// Start server
	log.Fatal(http.ListenAndServe(s.port, handler))
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Max-Age", "86400") // 24 hours
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}
