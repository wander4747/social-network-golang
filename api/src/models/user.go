package models

import (
	"api/src/security"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

// User = user
type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

//Prepare = call methods validate and format the user
func (user *User) Prepare(step string) error {
	if erro := user.validate(step); erro != nil {
		return erro
	}

	if erro := user.format(step); erro != nil {
		return erro
	}
	return nil
}

func (user *User) validate(step string) error {
	if user.Name == "" {
		return errors.New("Name is required and must not be blank")
	}

	if user.Nick == "" {
		return errors.New("Nick is required and must not be blank")
	}

	if user.Email == "" {
		return errors.New("Email is required and must not be blank")
	}

	if erro := checkmail.ValidateFormat(user.Email); erro != nil {
		return errors.New("E-mail invalid")
	}

	if step == "create" && user.Password == "" {
		return errors.New("Password is required and must not be blank")
	}

	return nil
}

func (user *User) format(step string) error {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)

	if step == "create" {
		passwordHash, erro := security.Hash(user.Password)
		if erro != nil {
			return erro
		}
		user.Password = string(passwordHash)
	}

	return nil
}
