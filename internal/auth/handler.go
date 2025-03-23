package auth

import (
	"fmt"
	"go/adv-demo/configs"
	"go/adv-demo/internal/user"
	"go/adv-demo/pkg/jwt"
	"go/adv-demo/pkg/req"
	res "go/adv-demo/pkg/respose"
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

type AuthHandler struct {
	*configs.Config
	*AuthService
	*RegRepository
}

type AuthHandlerDeps struct {
	*configs.Config
	*AuthService
	*RegRepository
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		Config:        deps.Config,
		AuthService:   deps.AuthService,
		RegRepository: deps.RegRepository,
	}
	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /auth/register", handler.RegUser())
	router.HandleFunc("POST /auth/newUser", handler.Create())
	router.HandleFunc("PATCH /auth/update/{id}", handler.Update())
	router.HandleFunc("DELETE /auth/{id}", handler.Delete())
}

func (c *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[LoginRequest](&w, r)
		fmt.Println(body)
		if err != nil {
			return
		}

		res_, err := c.AuthService.Logining(body.Email, body.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		nt, err := jwt.NewJWT(c.Auth.Secret).Create(jwt.JWTData{Email: res_})
		if err != nil {
			http.Error(w, "bad secret", http.StatusBadRequest)
			return
		}
		// res.Json(w, nt, 200)

		data := LoginResponse{
			Token: nt,
		}
		res.Json(w, data, 201)

	}
}

func (c *AuthHandler) RegUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[user.User_](&w, r)
		if err != nil {
			panic(err.Error())
		}
		res_, err := c.AuthService.Regester(body.Email, body.Password, body.Name)
		if err != nil {
			http.Error(w, "", http.StatusBadRequest)
		}
		nt, err := jwt.NewJWT(c.Auth.Secret).Create(jwt.JWTData{Email: res_})
		if err != nil {
			http.Error(w, "bad secret", http.StatusBadRequest)
			return
		}
		data := RegistResponse{
			Token: nt,
		}
		res.Json(w, data, 201)
	}
}

func (c *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		body, err := req.HandleBody[RegistRequest](&w, r)
		if err != nil {
			return
		}
		nt, err := jwt.NewJWT(c.Auth.Secret).Create(jwt.JWTData{Email: body.Email})
		if err != nil {
			http.Error(w, "bad secret", http.StatusBadRequest)
			return
		}
		data := RegistResponse{
			Token: nt,
		}
		res.Json(w, data, 201)
	}

}

func (c *AuthHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[RegistRequest](&w, r)
		if err != nil {
			panic(err.Error())
		}
		reg := NewReg(body.Email, body.Name, body.Password)
		createReg, err := c.CreateUser(reg)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		res.Json(w, createReg, 201)
	}
}

func (c *AuthHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[RegistRequest](&w, r)
		if err != nil {
			panic(err.Error())
		}
		idInt, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		retUp, err := c.UpdateById(&Reg{
			Model: gorm.Model{
				ID: uint(idInt),
			},
			Name:     body.Name,
			Password: body.Password,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		res.Json(w, retUp, 201)

	}
}

func (c *AuthHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		idInt, err := strconv.Atoi(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = c.DeleteById(idInt)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		res.Json(w, err, http.StatusOK)
	}
}
