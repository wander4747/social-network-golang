package controllers

import (
	"api/src/authentication"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/response"
	"api/src/security"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Login = Login user
func Login(w http.ResponseWriter, r *http.Request) {
	body, erro := ioutil.ReadAll(r.Body)

	if erro != nil {
		response.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var user models.User

	if erro = json.Unmarshal(body, &user); erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Connect()

	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repository := repositories.NewRepositoriesUser(db)
	userDatabase, erro := repository.GetEmail(user.Email)

	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if erro = security.CheckPassword(userDatabase.Password, user.Password); erro != nil {
		response.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	token, erro := authentication.CreateToken(userDatabase.ID)

	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	w.Write([]byte(token))

}
