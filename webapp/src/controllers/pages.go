package controllers

import (
	"net/http"
	"webapp/src/utils"
)

// LoadViewLogin = render view login
func LoadViewLogin(w http.ResponseWriter, r *http.Request) {
	utils.RunTemplate(w, "login.html", nil)
}

// LoadViewCreateUser = render view create user
func LoadViewCreateUser(w http.ResponseWriter, r *http.Request) {
	utils.RunTemplate(w, "register.html", nil)
}
