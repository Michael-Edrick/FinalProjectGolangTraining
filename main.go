package main

import (
	"FinalProject/db/connection"
	"FinalProject/handler"
	"FinalProject/repository"
	"FinalProject/service"
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
	
	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8088",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}





