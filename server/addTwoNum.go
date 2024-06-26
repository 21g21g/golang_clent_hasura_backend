package main

import (
	"encoding/json"
	"net/http"
)

func addTwo(w http.ResponseWriter, r *http.Request) {

	type Numbers struct {
		Number1 int `json:"num1"`
		Number2 int `json:"num2"`
	}

	type Result struct {
		Sum int `json:"the sum is"`
	}

	if r.Method == http.MethodGet {

		var number Numbers
		if err := json.NewDecoder(r.Body).Decode(&number); err != nil {
			http.Error(w, "the number you enter are not intger", http.StatusBadRequest)
			return
		}
		reult := number.Number1 + number.Number2

		totalSum := Result{
			Sum: reult,
		} 
		totalSumJson, err := json.Marshal(totalSum)
		if err != nil {
			http.Error(w, "the number you enter are not intger", http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-type", "application/json")

		w.Write(totalSumJson)
		// log.Printf("Accepted numbers: %d, %d", number.Number1, number.Number2)

	}

}