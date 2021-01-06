package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"user-crud/models"

	"github.com/go-chi/chi"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	userRepository models.IUserRepository
}

func NewUserController(userRepository models.IUserRepository) UserController {
	return UserController{userRepository: userRepository}
}

func (u *UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	users, err := u.userRepository.Get()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Something went wrong"))
		return
	}

	data, err := json.Marshal(users)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Something went wrong"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (u *UserController) GetUserById(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")

	id, err := strconv.ParseInt(idParam, 8, 10)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Something went wrong"))
		return
	}

	user, err := u.userRepository.Find(id)

	data, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Something went wrong"))
		return
	}

	w.Write(data)
}

//Update ..
func (u *UserController) Update(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	body := &models.User{}
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	password, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	user := &models.User{
		Username:  body.Username,
		Password:  string(password),
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Avatar:    body.Avatar,
	}

	err = u.userRepository.Update(int64(id), user)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Success updating user"))
}

//Delete ..
func (u *UserController) Delete(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	err = u.userRepository.Delete(int64(id))
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Success updating user"))
}
