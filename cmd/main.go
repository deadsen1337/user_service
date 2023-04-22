package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"user_service/internal/db"
	"user_service/internal/db/user"
	"user_service/internal/snowflake"
	"user_service/internal/user_service"
)

func main() {
	conn := snowflake.NewConnection()
	client := snowflake.NewClient(conn)
	dbConn := db.NewConnection()
	userRepo := user.NewRepository(dbConn)
	service := user_service.NewService(userRepo, client)

	r := mux.NewRouter()

	service.InitRoutes(r)

	log.Fatal(http.ListenAndServe(":7001", r))
}
