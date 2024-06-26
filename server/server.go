package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/add", addTwo)
	http.HandleFunc("/sub", subTwoNum)
	http.HandleFunc("/mult", multTwoNum)
	http.HandleFunc("/div", divTWoNum)
	http.HandleFunc("/changename", changeName)
	http.HandleFunc("/locationtype", updateLocation)
	http.ListenAndServe("localhost:4000", nil)

}
