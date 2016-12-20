package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

func init() {
	routes.Add("Encrypt Password", "GET", "/encrypt/{password}",
		"Encrypt the given password and return the hash.", handleEncrypt)
}

func handleEncrypt(w http.ResponseWriter, r *http.Request) {
	accessLog(r)

	vars := mux.Vars(r)
	passHash := passCrypt(vars["password"])

	w.Header().Set("Content-Type", "application/json")

	fmt.Fprint(w, passHash)
}

func passCrypt(pass string) string {
	p, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
	}
	return string(p)
}
