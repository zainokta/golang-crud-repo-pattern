package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"user-crud/models"
	"user-crud/utils"

	"golang.org/x/crypto/bcrypt"
)

type Credentials struct {
	Username string
	Password string
}

type AuthController struct {
	userRepository models.IUserRepository
}

func NewAuthController(userRepository models.IUserRepository) AuthController {
	return AuthController{userRepository: userRepository}
}

func (ac *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	creds := &Credentials{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&creds)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}
	input := &models.User{Username: creds.Username}

	result, err := ac.userRepository.FindUserByUsername(input.Username)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	if (models.User{}) != result {
		inputPassword := []byte(creds.Password)
		comparePassword := []byte(result.Password)

		err = bcrypt.CompareHashAndPassword(comparePassword, inputPassword)
		if err != nil {
			log.Println(err)
			log.Println(err)
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Invalid Username/Password"))
			return
		}

		token, err := utils.CreateToken(result.ID)
		if err != nil {
			log.Println(err)
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal Server Error"))
			return
		}

		log.Println(err)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(token))
		return
	}
}

type RegisterCredentials struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Avatar    string `json:"avatar"`
}

func (ac *AuthController) Register(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	register := &RegisterCredentials{}
	err := decoder.Decode(register)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	password, err := bcrypt.GenerateFromPassword([]byte(register.Password), 10)

	user := &models.User{
		Username:  register.Username,
		Password:  string(password),
		FirstName: register.FirstName,
		LastName:  register.LastName,
		Avatar:    register.Avatar,
	}

	userId, err := ac.userRepository.Create(user)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	token, _ := utils.CreateToken(userId)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(token))
	return
}
