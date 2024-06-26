package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func training() {
	var num1 int = 10
	var num2 int = 20
	if num1 > num2 {
		fmt.Println("this is the true print of the out put")
	} else {
		fmt.Println("why you greater one number from other number are you creasy.")
	}

}

func multTwoNum(w http.ResponseWriter, r *http.Request) {

	type Number struct {
		Number1 int `json:"num1"`
		Number2 int `json:"num2"`
	}

	type Result struct {
		Mult int `json:"the multiplication is"`
	}

	if r.Method == http.MethodPost {
		var numbers Number
		if err := json.NewDecoder(r.Body).Decode(&numbers); err != nil {
			http.Error(w, "the number you enter are not intger", http.StatusBadRequest)
			return
		}
		result := numbers.Number1 * numbers.Number2

		response := Result{
			Mult: result,
		}
		multresponse, err := json.Marshal(response)
		if err != nil {
			http.Error(w, "the number you enter are not intger", http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-type", "application/json")
		w.Write(multresponse)

	}
}