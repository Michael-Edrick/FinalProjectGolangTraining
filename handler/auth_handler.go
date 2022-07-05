package handler

import (
	"FinalProject/entity"
	"FinalProject/mapper"
	"FinalProject/utils"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type AuthHandler struct {
	r           *mux.Router
	userService entity.UserServiceInterface
}

func NewAuthHandler(r *mux.Router, userService entity.UserServiceInterface) {
	handler := AuthHandler{
		r:           r,
		userService: userService,
	}
	r.HandleFunc("/users/register", handler.userRegisterHandler)
	r.HandleFunc("/users/login", handler.userLoginHandler)
}

func (h AuthHandler) userRegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		//baca dr body
		decoder := json.NewDecoder(r.Body)
		var newUser entity.User
		err := decoder.Decode(&newUser)
		if err != nil {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write(utils.Response("", err))
			return
		}
		//masuk ke user service
		res, err := h.userService.UserRegisterService(&newUser)
		if err != nil {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write(utils.Response("", err))
			return
		}
		//keluarin response
		response := mapper.RegisterMapper(res)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(utils.Response(response, err))
	}
}

func (h AuthHandler) userLoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)
		var newLogin entity.User
		err := decoder.Decode(&newLogin)
		if err != nil {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write(utils.Response("", err))
			return
		}
		//autentikasi + generate jwt
		jwtToken, err := h.userService.UserLoginService(&newLogin) //(newLogin)
		if err != nil {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write(utils.Response("", err))
			return
		}
		response := map[string]string{
			"token": string(jwtToken),
		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(utils.Response(response, err))
	}
}
