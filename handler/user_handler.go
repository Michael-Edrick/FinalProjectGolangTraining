package handler

import (
	"FinalProject/entity"
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
	r.HandleFunc("/users/{Id}", handler.userUpdateHandler)
	r.HandleFunc("/users/", handler.userDeleteHandler)
	// r.Handle("/users/{Id}", middleware.IsAuthorized())

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
		if id != "" {
			fmt.Println(id)
			idInt, err := strconv.Atoi(id)
			if err != nil{
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
				jsonData, _ :=json.Marshal(&res)
				w.Header().Add("Content-Type", "application/json")
				w.Write(jsonData)
			}
		}
	}
}

func (h UserHandler)userDeleteHandler(w http.ResponseWriter, r *http.Request){
	if r.Method == "DELETE" {

		header:= r.Header.Get("Authorization")
		claims, err := utils.ParseJWT(header)
		if err != nil {
			res, _ := json.Marshal(err.Error())
			w.Header().Add("Content-Type", "application/json")
			w.Write(res)
			return
		}
		
		//masuk ke user service
		var loginEmail entity.User
		loginEmail.Email = claims["email"].(string)
		err = h.userService.UserDeleteService(loginEmail)
		if err != nil {
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

