package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func init() {
	routes.Add("API", "GET", "/",
		"Provides information about the API.", handleAPI)
}

func handleAPI(w http.ResponseWriter, r *http.Request) {
	accessLog(r)

	w.Header().Set("Content-Type", "application/json")

	res, err := json.Marshal(routes)
	if err != nil {
		log.Printf("API Error. Could not marshal JSON: %s", err)
	}

	fmt.Fprint(w, string(res))
}
