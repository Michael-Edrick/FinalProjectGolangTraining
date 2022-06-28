package main

import (
	"FinalProject/db/connection"
	"FinalProject/handler"
	"FinalProject/repository"
	"FinalProject/service"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

var PORT = ":8088"


func main(){

	db, err := connection.InitDatabase()
	if err != nil{
		log.Fatalf("%v\n",err)
	}
	r := mux.NewRouter()
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	handler.NewAuthHandler(r, userService)
	handler.NewUserHandler(r, userService)

	photoRepository := repository.NewPhotoRepository(db)
	photoService := service.NewPhotoService(photoRepository)
	handler.NewPhotoHandler(r, photoService)

	
	srv := &http.Server{
		Handler: r,
		Addr:    fmt.Sprintf("127.0.0.1%s", PORT),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}





