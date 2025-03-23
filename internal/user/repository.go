package user

import (
	"fmt"
	"go/adv-demo/pkg/db"
)

type UserRepository struct {
	Database *db.Db
}

func NewUserRepository(database *db.Db) *UserRepository {
	return &UserRepository{
		Database: database,
	}
}

func (u *UserRepository) CreateUser(us *User_) (*User_, error) {
	result := u.Database.DB.Create(us)
	if result.Error != nil {
		return nil, result.Error
	}
	return us, nil
}

func (u *UserRepository) FindByEmail(email string) (*User_, error) {
	var user User_
	resultFind := u.Database.DB.First(&user, "email = ?", email)
	if resultFind.Error != nil {
		fmt.Println("error", email)
		return nil, resultFind.Error
	}

	return &user, nil
}

func (u UserRepository) DeleteByID(id int) error {
	var user User_
	result := u.Database.DB.Delete(&user, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
