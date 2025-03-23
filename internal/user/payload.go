package user

type UserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

func NewUser(email, name, password string) *User_ {
	return &User_{
		Email:    email,
		Name:     name,
		Password: password,
	}
}
