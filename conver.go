package main

import (
	"net/http"
	"encoding/json"
	"fmt"
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
// update the instance in table convs
// update the instance in table convs
// update the instance in table convs
func upConvHandler(w http.ResponseWriter, r *http.Request) {
    var convert conv
    // take the first instance in table convs
    DB.First(&convert)

    // Parse request body
    var requestBody struct {
        History string `json:"history"`
    }
    err := json.NewDecoder(r.Body).Decode(&requestBody)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // delete all instance in table convs
	DB.Where("history = ?", convert.History).Delete(&convert)
	// create new instance in table convs
	DB.Create(&conv{History: requestBody.History})

    fmt.Fprintf(w, "History updated: %s", requestBody.History)
}

