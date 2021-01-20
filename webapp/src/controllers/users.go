package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// CreateUser = Create new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	user, erro := json.Marshal(map[string]string{
		"name":     r.FormValue("name"),
		"email":    r.FormValue("email"),
		"nick":     r.FormValue("nick"),
		"password": r.FormValue("password"),
	})

	if erro != nil {
		log.Fatal(erro)
	}

	response, erro := http.Post("http://localhost:5000/users", "application/json", bytes.NewBuffer(user))

	if erro != nil {
		log.Fatal(erro)
	}

	defer response.Body.Close()

	fmt.Println(response.Body)

}
