package models

// User :- Use for constructing the user's information
type User struct {
	ID        int    `json:"user_id"`
	Firstname string `json:"first_name"`
	Lastname  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

// UserStore :- This is a dictionary of users
var UserStore = make(map[string]User)
