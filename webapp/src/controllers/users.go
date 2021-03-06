package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"webapp/src/response"
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
		response.JSON(w, http.StatusBadRequest, response.ErroAPI{Erro: erro.Error()})
		return
	}

	result, erro := http.Post("http://localhost:5000/users", "application/json", bytes.NewBuffer(user))

	if erro != nil {
		response.JSON(w, http.StatusInternalServerError, response.ErroAPI{Erro: erro.Error()})
		return
	}

	defer result.Body.Close()

	if result.StatusCode >= 400 {
		response.ProcessStatusCodeErro(w, result)
		return
	}

	response.JSON(w, result.StatusCode, nil)
}
