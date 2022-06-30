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

type SocialMediaHandler struct {
	r                  *mux.Router
	socialMediaService entity.SocialMediaServiceInterface
}

func NewSocialMediaHandler(r *mux.Router, socialMediaService entity.SocialMediaServiceInterface) {
	handler := SocialMediaHandler{
		r:                  r,
		socialMediaService: socialMediaService,
	}
	s := r.PathPrefix("").Subrouter()
	s.Use(middleware.IsAuthorized())
	s.HandleFunc("/socialmedias", handler.socialMediaPostGetHandler)
	s.HandleFunc("/socialmedias/{Id}", handler.socialMediaUpdateDeleteHandler)
}

func (h SocialMediaHandler) socialMediaPostGetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		//baca dr body
		decoder := json.NewDecoder(r.Body)
		var postSocialMedia entity.SocialMedia
		err := decoder.Decode(&postSocialMedia)
		if err != nil {
			w.Write([]byte("error"))
			return
		}
		//get claims
		header := r.Header.Get("Authorization")
		claims, err := utils.ParseJWT(header)
		if err != nil {
			res, _ := json.Marshal(err.Error())
			w.Header().Add("Content-Type", "application/json")
			w.Write(res)
			return
		}
		postSocialMedia.UserId = int(claims["userid"].(float64))
		//masuk ke social media service
		res, err := h.socialMediaService.SocialMediaPostService(&postSocialMedia)
		if err != nil {
			res, _ := json.Marshal(err.Error())
			w.Header().Add("Content-Type", "application/json")
			w.Write(res)
			return
		}
		//keluarin response
		response := mapper.PostSocialMediaMapper(res)
		jsonData, _ := json.Marshal(&response)
		w.Header().Add("Content-Type", "application/json")
		w.Write(jsonData)
	}
	if r.Method == "GET" {
		//get claims
		header := r.Header.Get("Authorization")
		claims, err := utils.ParseJWT(header)
		if err != nil {
			res, _ := json.Marshal(err.Error())
			w.Header().Add("Content-Type", "application/json")
			w.Write(res)
			return
		}
		//masuk ke social media service
		var getSocialMedias entity.SocialMedia
		getSocialMedias.UserId = int(claims["userid"].(float64))
		res, err := h.socialMediaService.SocialMediaGetService(&getSocialMedias)
		if err != nil {
			res, _ := json.Marshal(err.Error())
			w.Header().Add("Content-Type", "application/json")
			w.Write(res)
			return
		}
		//keluarin response
		response := mapper.GetSocialMediaMapper(res)
		jsonData, _ := json.Marshal(&response)
		w.Header().Add("Content-Type", "application/json")
		w.Write(jsonData)
	}
}

func (h SocialMediaHandler) socialMediaUpdateDeleteHandler(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	id := param["Id"]
	if r.Method == "PUT" {
		//baca dr body
		decoder := json.NewDecoder(r.Body)
		var updateSocialMedia entity.SocialMedia
		err := decoder.Decode(&updateSocialMedia)
		if err != nil {
			w.Write([]byte("error"))
			return
		}
		//convert param
		if id != "" {
			fmt.Println(id)
			idInt, err := strconv.Atoi(id)
			if err != nil {
				w.Write([]byte("error"))
				return
			} else {
				updateSocialMedia.Id = idInt
			}
			//masuk ke social media service
			res, err := h.socialMediaService.SocialMediaUpdateService(&updateSocialMedia)
			if err != nil {
				res, _ := json.Marshal(err.Error())
				w.Header().Add("Content-Type", "application/json")
				w.Write(res)
				return
			}
			//keluarin response
			response := mapper.UpdateSocialMediaMapper(res)
			jsonData, _ := json.Marshal(&response)
			w.Header().Add("Content-Type", "application/json")
			w.Write(jsonData)
		}
	}
	if r.Method == "DELETE" {
		var deleteSocialMedia entity.SocialMedia
		//convert param
		if id != "" {
			fmt.Println(id)
			idInt, err := strconv.Atoi(id)
			if err != nil {
				w.Write([]byte("error"))
				return
			} else {
				deleteSocialMedia.Id = idInt
				//masuk ke social media service
				err = h.socialMediaService.SocialMediaDeleteService(&deleteSocialMedia)
				if err == nil {
					response := map[string]string{
						"message": "Your social media has been successfully deleted",
					}
					res, _ := json.Marshal(response)
					w.Header().Add("Content-Type", "application/json")
					w.Write(res)
					return
				}
			}
		}
	}
}
