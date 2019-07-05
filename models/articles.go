package models

// users type details
type Article struct {
	Id   int
	Title string
	UserId int
	User User `json:",omitempty"`
  }