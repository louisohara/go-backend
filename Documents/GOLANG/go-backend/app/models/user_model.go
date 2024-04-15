package models

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	Surname   string `json:"surname"`
	Email     string `json:"email"`
	DOB 	  string `json:"dob"`
	// Add other fields as per your requirements
}