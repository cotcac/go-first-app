package models

// users type details
type Article struct {
	Id   int
	Title string
	UserId int
	User User `json:",omitempty"`
  }
// display 

type Articles struct {
	Id   int
	Title string
	User struct {
		Id int 
		Name string
	}
	Categories struct {
		Id int 
		Title string
	}
}
