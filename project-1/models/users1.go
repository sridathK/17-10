package models

type User struct {
	Name        string //structure
	Email       string
	password    string
	Permissions map[string]bool
}
