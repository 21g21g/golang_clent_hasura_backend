package main

import (
	"encoding/json"
	"net/http"
)

func subTwoNum(w http.ResponseWriter, r *http.Request) {

	type Numbers struct {
		Number1 int `json:"num1"`
		Number2 int `json:"num2"`
	}

	type Result struct {
		Sub int `json:"the difference is"`
	}

	if r.Method == http.MethodGet {

		var number Numbers
		if err := json.NewDecoder(r.Body).Decode(&number); err != nil {
			http.Error(w, "the number you enter are not intger", http.StatusBadRequest)
			return
		}

		result := number.Number1 - number.Number2

		response := Result{
			Sub: result,
		}
		responseJson, err := json.Marshal(response)
		if err != nil {
			http.Error(w, "the number you enter are not intger", http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(responseJson)

	}
}
