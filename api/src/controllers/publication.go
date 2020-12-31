package controllers

import (
	"api/src/authentication"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/response"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// CreatePublication = route to save new publication
func CreatePublication(w http.ResponseWriter, r *http.Request) {
	userID, erro := authentication.ExtractUserID(r)
	if erro != nil {
		response.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	body, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		response.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var publication models.Publication
	if erro = json.Unmarshal(body, &publication); erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	publication.AuthorID = userID

	if erro = publication.Prepare(); erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repository := repositories.NewRepositoriesPublication(db)
	publication.ID, erro = repository.Create(publication)

	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusCreated, publication)
}

// AllPublication = Get all publication
func AllPublication(w http.ResponseWriter, r *http.Request) {
	userID, erro := authentication.ExtractUserID(r)
	if erro != nil {
		response.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repository := repositories.NewRepositoriesPublication(db)
	publications, erro := repository.All(userID)

	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusOK, publications)

}

// FindPublication = Get publication
func FindPublication(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	publicationID, erro := strconv.ParseUint(parameters["id"], 10, 64)

	if erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repository := repositories.NewRepositoriesPublication(db)
	publication, erro := repository.Find(publicationID)

	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusOK, publication)
}

// UpdatePublication = Update publication
func UpdatePublication(w http.ResponseWriter, r *http.Request) {

}

// DeletePublication = Delete publication
func DeletePublication(w http.ResponseWriter, r *http.Request) {

}
