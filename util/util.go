package util

import (
	"encoding/json"
	"log"
	"net/http"
)

func ResponseWithJson(w http.ResponseWriter, code int, payload interface{}){
	response, err := json.Marshal(payload) 

	if err != nil{
		log.Fatal("Unable to marshal server responsd with 5XX")
		w.WriteHeader(code)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response) 
	
}


func ResponseWithError(w http.ResponseWriter, code int, message string){
	if (code > 499){
		log.Fatal("Something went Wrong")
	}

	type ErrorMessage struct{
		Error string `json:"error"`
	}

	ResponseWithJson(w, code, ErrorMessage{
		Error:message,
	})
}