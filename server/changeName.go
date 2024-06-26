package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func changeName(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "methode not allowed", http.StatusMethodNotAllowed)
		return
	}

	type requestData struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}

	var data requestData

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "the name you enterd is not comfort with its format", http.StatusBadRequest)
	}
	graphqlEndpoint := "http://localhost:8083/v1/graphql"
	mutation := `mutation updateChar($id:Int!,$name:String!){
		update_results_character_by_pk(pk_columns:{id:$id}_set:{name:$name}){
		  id
		  name
		  
		}
	  }`
	names := map[string]interface{}{
		"id":   data.ID,
		"name": data.Name,
	}
	//the purpose of the above declaration is inorder to declar key value pairs in go the key type is string and its value is interface that means it holds any data types.

	reqBodys := map[string]interface{}{
		"query":     mutation,
		"variables": names,
	}
	reqMarshal, err := json.Marshal(reqBodys)

	if err != nil {
		http.Error(w, "error when marshaling the bdy", http.StatusInternalServerError)
		return

	}

	req, err := http.NewRequest("POST", graphqlEndpoint, bytes.NewBuffer(reqMarshal))
	if err != nil {
		http.Error(w, "request is not handled correctly", http.StatusInternalServerError)
		return
	}
	req.Header.Set("Content-type", "application/json")

	//create http client
	client := &http.Client{}
	//send http request
	res, err := client.Do(req)

	if err != nil {
		http.Error(w, "the request is not send", http.StatusInternalServerError)
		return
	}
	defer res.Body.Close()
	var response map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		http.Error(w, "Failed to decode response body", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Successfully updated character: %+v", response)

}
