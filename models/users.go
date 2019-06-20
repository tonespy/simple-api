package models

type missingFieldError string

func (m missingFieldError) Error() string {
	return string(m) + " is required"
}

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

// OK represents types capable of validating themselves.
// This would be utilized when decoding from request body.
// Here, we can decide to check for required fields as needed.
func (u *User) OK() error {
	if len(u.Firstname) == 0 {
		return missingFieldError("first_name")
	}

	if len(u.Lastname) == 0 {
		return missingFieldError("last_name")
	}

	if len(u.Email) == 0 {
		return missingFieldError("email")
	}

	if len(u.Password) == 0 {
		return missingFieldError("password")
	}
	return nil
}

// UserStore :- This is a dictionary of users
var UserStore = make(map[string]User)
