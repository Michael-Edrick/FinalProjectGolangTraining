package handler

import (
	"FinalProject/entity"
	"FinalProject/middleware"
	"net/http"

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
	// s.HandleFunc("/photos/{Id}", handler.photoUpdateDeleteHandler)
}

func (h SocialMediaHandler) socialMediaPostGetHandler(w http.ResponseWriter, r *http.Request) {

}
