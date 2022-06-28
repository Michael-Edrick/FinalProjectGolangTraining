package handler

import (
	"FinalProject/entity"
	"FinalProject/mapper"
	"FinalProject/middleware"
	"FinalProject/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)



type UserHandler struct {
	r *mux.Router
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

func (h UserHandler) userUpdateHandler (w http.ResponseWriter, r *http.Request){
	param := mux.Vars(r)
	id := param["Id"]
	fmt.Printf("%v\n", r.Method)
	if r.Method == "PUT" {
		//baca dr body
		decoder := json.NewDecoder(r.Body)
		var updateUser entity.User
		err := decoder.Decode(&updateUser)
		if err != nil {
			w.Write([]byte("error"))
			return
		}
		//convert param
		if id != "" {
			fmt.Println(id)
			idInt, err := strconv.Atoi(id)
			if err != nil{
				w.Write([]byte("error"))
				return
			}else{
				updateUser.Id = idInt
				//masuk ke user service
				res, err := h.userService.UserUpdateService(updateUser)
				if err!= nil{
					res, _ := json.Marshal(err.Error())
					w.Header().Add("Content-Type", "application/json")
					w.Write(res)
					return
				}
				//keluarin response
				response := mapper.UpdateMapper(res)
				jsonData, _ :=json.Marshal(&response)
				w.Header().Add("Content-Type", "application/json")
				w.Write(jsonData)
			}
		}
	}
}

func (h UserHandler)userDeleteHandler(w http.ResponseWriter, r *http.Request){
	if r.Method == "DELETE" {
		//get claims
		header:= r.Header.Get("Authorization")
		claims, err := utils.ParseJWT(header)
		if err != nil {
			res, _ := json.Marshal(err.Error())
			w.Header().Add("Content-Type", "application/json")
			w.Write(res)
			return
		}
		
		//masuk ke user service
		var loginUser entity.User
		loginUser.Id= int(claims["userid"].(float64))
		err = h.userService.UserDeleteService(loginUser)
		if err == nil {
			response := map[string]string{
				"message": "Your account has been successfully deleted",
			}
			res, _ := json.Marshal(response)
			w.Header().Add("Content-Type", "application/json")
			w.Write(res)
			return
		}
	}
}

