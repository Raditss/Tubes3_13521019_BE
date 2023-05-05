package main

import (
	// "encoding/json"
	"net/http"
)



type conv struct {
	History string `json:"history"`
}


// create a new handler for the /conv endpoint to get all instance in table conv

func convHandler(w http.ResponseWriter, r *http.Request) {
	var convert conv
	// take all instance in table conv
	DB.Find(&convert)
	// return all instance in table conv
	JSONWriter(w, convert)
}