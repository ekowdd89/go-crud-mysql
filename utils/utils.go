package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Data    any    `json:"data,omitempty"`
	Err     any    `json:"errors,omitempty"`
}

func Responder(w http.ResponseWriter, res Response) {
	b, err := json.Marshal(res)
	if err != nil {
		log.Panic(err)
	}
	w.Write(b)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Accept", "application/json")
}
