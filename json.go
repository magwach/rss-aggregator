package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithJson(w http.ResponseWriter, code int, payload any){
	data, err := json.Marshal(payload)

	if err != nil{
		log.Printf("Error in marshaling JSON response %v", payload)
		w.WriteHeader(500)
		return;
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}