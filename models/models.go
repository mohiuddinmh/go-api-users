package models

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	DoB       string `json:"dob"` // Assuming DoB is stored as a string in the database
}
