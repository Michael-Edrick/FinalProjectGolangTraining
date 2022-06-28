package handler

import (
	"FinalProject/entity"
	"FinalProject/mapper"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type AuthHandler struct{
	r *mux.Router
	userService entity.UserServiceInterface
}

func NewAuthHandler(r *mux.Router, userService entity.UserServiceInterface){
	handler := AuthHandler{
		r : r,
		userService: userService, 
	}
	r.HandleFunc("/users/register", handler.userRegisterHandler)
	r.HandleFunc("/users/login", handler.userLoginHandler)
	
}

func (h AuthHandler) userRegisterHandler (w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		//baca dr body
		decoder := json.NewDecoder(r.Body)
		var newUser entity.User
		err := decoder.Decode(&newUser)
		if err != nil {
			w.Write([]byte("error"))
			return
		}
		//masuk ke user service
		res, err := h.userService.UserRegisterService(newUser)
		if err != nil {
			res, _ := json.Marshal(err.Error())
			w.Header().Add("Content-Type", "application/json")
			w.Write(res)
			return
		}
		//keluarin response
		response := mapper.RegisterMapper(res)
		jsonData, _ := json.Marshal(&response)
		w.Header().Add("Content-Type", "application/json")
		w.Write(jsonData)
	}
}

func (h AuthHandler)userLoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)
		var newLogin entity.User
		err := decoder.Decode(&newLogin)
		if err != nil {
			w.Write([]byte("error"))
			return
		}
		//autentikasi + generate jwt
		jwtToken, err := h.userService.UserLoginService(newLogin) //(newLogin)
		if err !=nil {
			errorMessage, _ := json.Marshal(err.Error())
			w.WriteHeader(http.StatusUnauthorized)
			w.Write(errorMessage)
			return
		} 
		response := map[string]string{
			"token": string(jwtToken),
		}
		responseJson, _ := json.Marshal(&response)
		w.Header().Add("Content-Type", "application/json")
		w.Write(responseJson)
	}
}