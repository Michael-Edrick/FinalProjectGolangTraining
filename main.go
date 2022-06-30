package main

import (
	"FinalProject/db/connection"
	"FinalProject/handler"
	"FinalProject/repository"
	"FinalProject/service"
	"FinalProject/utils"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	config := utils.InitConfig()
	db, err := connection.InitDatabase(config)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	r := mux.NewRouter()
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	handler.NewAuthHandler(r, userService)
	handler.NewUserHandler(r, userService)

	photoRepository := repository.NewPhotoRepository(db)
	photoService := service.NewPhotoService(photoRepository)
	handler.NewPhotoHandler(r, photoService)

	commentRepository := repository.NewCommentRepository(db)
	commentService := service.NewCommentService(commentRepository)
	handler.NewCommentHandler(r, commentService)

	socialMediaRepository := repository.NewSocialMediaRepository(db)
	socialMediaService := service.NewSocialMediaService(socialMediaRepository)
	handler.NewSocialMediaHandler(r, socialMediaService)

	srv := &http.Server{
		Handler:      r,
		Addr:         fmt.Sprintf("127.0.0.1:%s", config.Port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
