package main

import (
	"encoding/json"
	"net/http"
)

func divTWoNum(w http.ResponseWriter, r *http.Request) {
	type Number struct {
		Number1 float64 `json:"num1"`
		Number2 float64 `json:"num2"`
	}

	type Result struct {
		Div float64 `json:"the division is"`
	}

	if r.Method == http.MethodPost {
		var number Number

		if err := json.NewDecoder(r.Body).Decode(&number); err != nil {
			http.Error(w, "the number you enter are not intger", http.StatusBadRequest)
			return
		}
		if number.Number2 == 0 {
			w.Write([]byte("you cannot devide number by zero")) //used to write textual responses directly to the client's HTTP response body.
			w.Header()                                          //to manipulate the HTTP headers before sending them to the client. You can add, delete, or modify headers using methods like Add, Set, Del, etc.
			//w.WriteHeader()// used to set the HTTP status code (e.g., 200 OK, 404 Not Found, etc.) based on the result of processing a request.

		} else {
			result := number.Number1 / number.Number2
			response := Result{
				Div: result,
			}
			responseJson, err := json.Marshal(response)

			if err != nil {
				http.Error(w, "the number you enter are not intger", http.StatusBadRequest)
				return
			}
			w.Header().Set("Content-type", "application/json")
			w.Write(responseJson)

		}

	}
}