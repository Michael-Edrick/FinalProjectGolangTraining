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

type PhotoHandler struct {
	r            *mux.Router
	photoService entity.PhotoServiceInterface
}

func NewPhotoHandler(r *mux.Router, photoService entity.PhotoServiceInterface) {
	handler := PhotoHandler{
		r:            r,
		photoService: photoService,
	}
	s := r.PathPrefix("").Subrouter()
	s.Use(middleware.IsAuthorized())
	s.HandleFunc("/photos", handler.photoPostGetHandler)
	s.HandleFunc("/photos/{Id}", handler.photoUpdateDeleteHandler)
}

func (h PhotoHandler) photoPostGetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		//baca dr body
		decoder := json.NewDecoder(r.Body)
		var postPhoto entity.Photo
		err := decoder.Decode(&postPhoto)
		if err != nil {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write(utils.Response("", err))
			return
		}
		//get claims
		header := r.Header.Get("Authorization")
		claims, err := utils.ParseJWT(header)
		if err != nil {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write(utils.Response("", err))
			return
		}
		postPhoto.UserId = int(claims["userid"].(float64))
		//masuk ke photo service
		res, err := h.photoService.PhotoPostService(&postPhoto)
		if err != nil {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write(utils.Response("", err))
			return
		}
		//keluarin response
		response := mapper.PostPhotoMapper(res)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(utils.Response(response, err))
	}
	if r.Method == "GET" {
		//get claims
		header := r.Header.Get("Authorization")
		claims, err := utils.ParseJWT(header)
		if err != nil {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write(utils.Response("", err))
			return
		}
		//masuk ke photo service
		var getPhotos entity.Photo
		getPhotos.UserId = int(claims["userid"].(float64))
		res, err := h.photoService.PhotoGetService(&getPhotos)
		if err != nil {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write(utils.Response("", err))
			return
		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(utils.Response(res, err))
	}
}

func (h PhotoHandler) photoUpdateDeleteHandler(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	id := param["Id"]
	if r.Method == "PUT" {
		//baca dr body
		decoder := json.NewDecoder(r.Body)
		var updatePhoto entity.Photo
		err := decoder.Decode(&updatePhoto)
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
				updatePhoto.Id = idInt
				//masuk ke photo service
				res, err := h.photoService.PhotoUpdateService(&updatePhoto)
				if err != nil {
					w.Header().Add("Content-Type", "application/json")
					w.WriteHeader(http.StatusBadRequest)
					w.Write(utils.Response("", err))
					return
				}
				//keluarin response
				response := mapper.UpdatePhotoMapper(res)
				w.Header().Add("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				w.Write(utils.Response(response, err))
			}
		}
	}
	if r.Method == "DELETE" {
		var deletePhoto entity.Photo
		//convert param
		if id != "" {
			idInt, err := strconv.Atoi(id)
			if err != nil {
				w.Header().Add("Content-Type", "application/json")
				w.WriteHeader(http.StatusBadRequest)
				w.Write(utils.Response("", err))
				return
			} else {
				deletePhoto.Id = idInt
				//masuk ke photo service
				err = h.photoService.PhotoDeleteService(&deletePhoto)
				if err == nil {
					response := map[string]string{
						"message": "Your photo has been successfully deleted",
					}
					w.Header().Add("Content-Type", "application/json")
					w.WriteHeader(http.StatusOK)
					w.Write(utils.Response(response, err))
					return
				}
			}
		}
	}
}
