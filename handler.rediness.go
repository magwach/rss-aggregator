package main

import "net/http"

func HandlerRediness(w http.ResponseWriter, r *http.Request){
	respondWithJson(w, 200, struct{}{})
}