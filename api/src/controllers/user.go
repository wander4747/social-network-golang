package controllers

import (
	"api/src/authentication"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/response"
	"api/src/security"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// Create = route to save new user
func Create(w http.ResponseWriter, r *http.Request) {
	bodyResquest, erro := ioutil.ReadAll(r.Body)

	if erro != nil {
		response.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var user models.User

	if erro = json.Unmarshal(bodyResquest, &user); erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = user.Prepare("create"); erro != nil {
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
	userID, erro := repository.Create(user)

	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	user.ID = userID
	response.JSON(w, http.StatusCreated, user)
}

// All = Get all user
func All(w http.ResponseWriter, r *http.Request) {

	nameOrNick := strings.ToLower(r.URL.Query().Get("user"))

	db, erro := database.Connect()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repository := repositories.NewRepositoriesUser(db)

	users, erro := repository.Search(nameOrNick)

	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusOK, users)
}

// Find = Get user
func Find(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	userID, erro := strconv.ParseUint(parameters["id"], 10, 64)

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

	repository := repositories.NewRepositoriesUser(db)

	user, erro := repository.Find(userID)

	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusOK, user)
}

// Update = Update user
func Update(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	userID, erro := strconv.ParseUint(parameters["id"], 10, 64)

	if erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	userIDToken, erro := authentication.ExtractUserID(r)

	if erro != nil {
		response.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	if userID != userIDToken {
		response.Erro(w, http.StatusForbidden, errors.New("You cannot update a user other than your own"))
		return
	}

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

	if erro = user.Prepare("update"); erro != nil {
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
	if erro = repository.Update(userID, user); erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}

// Delete = Delete user
func Delete(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	userID, erro := strconv.ParseUint(parameters["id"], 10, 64)

	if erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	userIDToken, erro := authentication.ExtractUserID(r)

	if erro != nil {
		response.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	if userID != userIDToken {
		response.Erro(w, http.StatusForbidden, errors.New("You cannot delete a user other than your own"))
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repository := repositories.NewRepositoriesUser(db)
	if erro = repository.Delete(userID); erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}

// FollowUser = Follow user
func FollowUser(w http.ResponseWriter, r *http.Request) {
	followerID, erro := authentication.ExtractUserID(r)

	if erro != nil {
		response.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	parameters := mux.Vars(r)

	userID, erro := strconv.ParseUint(parameters["id"], 10, 64)

	if erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if followerID == userID {
		response.Erro(w, http.StatusForbidden, errors.New("You can't follow yourself"))
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repository := repositories.NewRepositoriesUser(db)

	if erro = repository.Follow(userID, followerID); erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}

// UnFollowUser = Unfollow User
func UnFollowUser(w http.ResponseWriter, r *http.Request) {
	followerID, erro := authentication.ExtractUserID(r)

	if erro != nil {
		response.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	parameters := mux.Vars(r)

	userID, erro := strconv.ParseUint(parameters["id"], 10, 64)

	if erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if followerID == userID {
		response.Erro(w, http.StatusForbidden, errors.New("You cannot stop following yourself"))
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repository := repositories.NewRepositoriesUser(db)

	if erro = repository.UnFollow(userID, followerID); erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}

// Followers = Followers User
func Followers(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	userID, erro := strconv.ParseUint(parameters["id"], 10, 64)

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

	repository := repositories.NewRepositoriesUser(db)
	followers, erro := repository.Followers(userID)

	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusOK, followers)
}

// Following = Following User
func Following(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	userID, erro := strconv.ParseUint(parameters["id"], 10, 64)

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

	repository := repositories.NewRepositoriesUser(db)
	followers, erro := repository.Following(userID)

	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusOK, followers)
}

// UpdatePassword = Update password user
func UpdatePassword(w http.ResponseWriter, r *http.Request) {
	userIDNoToken, erro := authentication.ExtractUserID(r)

	if erro != nil {
		response.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	parameters := mux.Vars(r)
	userID, erro := strconv.ParseUint(parameters["id"], 10, 64)

	if erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if userIDNoToken != userID {
		response.Erro(w, http.StatusForbidden, errors.New("You cannot update the password for a user other than your own"))
		return
	}

	body, erro := ioutil.ReadAll(r.Body)

	var password models.Password
	if erro = json.Unmarshal(body, &password); erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
	}

	db, erro := database.Connect()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repository := repositories.NewRepositoriesUser(db)
	passwordDatabase, erro := repository.SearchPassword(userID)

	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if erro = security.CheckPassword(passwordDatabase, password.Actual); erro != nil {
		response.Erro(w, http.StatusUnauthorized, errors.New("The current password does not match the one saved in the database"))
		return
	}

	passwordHash, erro := security.Hash(password.New)
	if erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = repository.UpdatePassword(userID, string(passwordHash)); erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}
