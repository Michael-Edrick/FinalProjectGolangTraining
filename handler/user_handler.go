package handler

import (
	"FinalProject/entity"
	"FinalProject/mapper"
	"FinalProject/middleware"
	"FinalProject/utils"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type UserHandler struct {
	r           *mux.Router
	userService entity.UserServiceInterface
}

func NewUserHandler(r *mux.Router, userService entity.UserServiceInterface) {
	handler := UserHandler{
		r:           r,
		userService: userService,
	}
	s := r.PathPrefix("").Subrouter()
	s.Use(middleware.IsAuthorized())
	s.HandleFunc("/users/{Id}", handler.userUpdateHandler)
	s.HandleFunc("/users", handler.userDeleteHandler)
}

func (h UserHandler) userUpdateHandler(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	id := param["Id"]
	if r.Method == "PUT" {
		//baca dr body
		decoder := json.NewDecoder(r.Body)
		var updateUser entity.User
		err := decoder.Decode(&updateUser)
		if err != nil {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write(utils.Response("", err))
			return
		}
		//convert param
		if id != "" {
			idInt, err := strconv.Atoi(id)
			if err != nil {
				w.Header().Add("Content-Type", "application/json")
				w.WriteHeader(http.StatusBadRequest)
				w.Write(utils.Response("", err))
				return
			} else {
				updateUser.Id = idInt
				//masuk ke user service
				res, err := h.userService.UserUpdateService(&updateUser)
				if err != nil {
					w.Header().Add("Content-Type", "application/json")
					w.WriteHeader(http.StatusBadRequest)
					w.Write(utils.Response("", err))
					return
				}
				//keluarin response
				response := mapper.UpdateMapper(res)
				w.Header().Add("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				w.Write(utils.Response(response, err))
			}
		}
	}
}

func (h UserHandler) userDeleteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "DELETE" {
		//get claims
		header := r.Header.Get("Authorization")
		claims, err := utils.ParseJWT(header)
		if err != nil {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write(utils.Response("", err))
			return
		}
		//masuk ke user service
		var loginUser entity.User
		loginUser.Id = int(claims["userid"].(float64))
		err = h.userService.UserDeleteService(&loginUser)
		if err == nil {
			response := map[string]string{
				"message": "Your account has been successfully deleted",
			}
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(utils.Response(response, err))
			return
		}
	}
}
