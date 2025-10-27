package main

import "net/http"

func HandlerRediness(w http.ResponseWriter, r *http.Request){
	RespondWithJson(w, 200, struct{}{})
}