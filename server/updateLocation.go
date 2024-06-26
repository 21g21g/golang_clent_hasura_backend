package main

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func updateLocation(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "the request you ask is incorrect", http.StatusMethodNotAllowed)
		return
	}

	type locationType struct {
		ID   int    `json:"id"`
		Type string `json:"type"`
	}

	var location locationType
	if err := json.NewDecoder(r.Body).Decode(&location); err != nil {
		http.Error(w, "the coming request is not accepted", http.StatusBadRequest)
		return

	}
	graphqlEndpoint := "http://localhost:8083/v1/graphql"
	mutation := `mutation changeLocationType($id:Int!,$type:String!){
		update_results_location_by_pk(pk_columns:{id:$id},_set:{type:$type}){
		  id
		  type
		}
		
	  }`
	// map is used for storing key-value pair.
	types := map[string]interface{}{
		"id":   location.ID,
		"type": location.Type,
	}
	typeRequest := map[string]interface{}{
		"query":     mutation,
		"variables": types,
	}

	typeMarshal, err := json.Marshal(typeRequest)
	if err != nil {
		http.Error(w, "casting into json is not succesful", http.StatusNotAcceptable)
		return
	}

	reqtoHasura, err := http.NewRequest("POST", graphqlEndpoint, bytes.NewBuffer(typeMarshal))
	
	if err != nil {
		http.Error(w, "your request not send to hasura", http.StatusServiceUnavailable)
		return
	}
	reqtoHasura.Header.Set("Content-type", "application-json")
	client := &http.Client{}

	res, err := client.Do(reqtoHasura)
	if err != nil {
		http.Error(w, "your res is not send correctly", http.StatusInternalServerError)
		return
	}
	defer res.Body.Close()

}
