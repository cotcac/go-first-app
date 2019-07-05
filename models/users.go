package models

// users type details
type User struct {
	Id   int
	Name string
	Email string `json:",omitempty"`
  }