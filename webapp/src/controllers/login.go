package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"webapp/src/response"
)

// DoLogin = uses email and password to access the application
func DoLogin(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.ParseForm())
	user, erro := json.Marshal(map[string]string{
		"email":    r.FormValue("email"),
		"password": r.FormValue("password"),
	})

	if erro != nil {
		response.JSON(w, http.StatusBadRequest, response.ErroAPI{Erro: erro.Error()})
		return
	}

	result, erro := http.Post("http://localhost:5000/login", "application/json", bytes.NewBuffer(user))

	if erro != nil {
		response.JSON(w, http.StatusInternalServerError, response.ErroAPI{Erro: erro.Error()})
		return
	}

	defer result.Body.Close()

	token, _ := ioutil.ReadAll(result.Body)

	fmt.Println(result.StatusCode, string(token))
}
