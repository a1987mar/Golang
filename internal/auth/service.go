package auth

import (
	"errors"
	"fmt"
	"go/adv-demo/internal/user"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UserRepository *user.UserRepository
}

func NewAuthService(userRepository *user.UserRepository) *AuthService {
	return &AuthService{
		UserRepository: userRepository,
	}
}

func (service *AuthService) Regester(email, password, name string) (string, error) {

	existedUser, _ := service.UserRepository.FindByEmail(email)
	if existedUser != nil {
		returnString := fmt.Sprintf("вже існує, password %s", existedUser.Password)
		return returnString, nil
	}

	bcryptoPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	us := user.NewUser(email, name, string(bcryptoPassword))

	res_, err := service.UserRepository.CreateUser(us)
	if err != nil {
		return "", err
	}
	return res_.Email, nil

}

func (service *AuthService) Logining(email, passwork string) (string, error) {
	existedUser, _ := service.UserRepository.FindByEmail(email)

	if existedUser == nil {
		return "record not found", errors.New(ErrWrongCredetials)
	}
	err := bcrypt.CompareHashAndPassword([]byte(existedUser.Password), []byte(passwork))
	if err != nil {
		return "record not found", errors.New(ErrWrongCredetials)
	}
	return email, nil
}
