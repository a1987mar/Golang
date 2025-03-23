package auth

import "gorm.io/gorm"

type Reg struct {
	gorm.Model
	Email    string `json:"email" gorm:"uniqueIndex"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

func NewReg(email, name, password string) *Reg {
	return &Reg{
		Email:    email,
		Name:     name,
		Password: password,
	}
}
