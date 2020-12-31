package models

// Password = password
type Password struct {
	New    string `json:new`
	Actual string `json:actual`
}
