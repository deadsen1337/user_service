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
	conn, err := snowflake.NewConnection()
	if err != nil {
		log.Fatalf("grpc connect ffs error: %s", err)
	}
	client := snowflake.NewClient(conn)
	dbConn, err := db.NewConnection()
	if err != nil {
		log.Fatalf("could not open db connection: %s", err)
	}
	userRepo := user.NewRepository(dbConn)
	service := user_service.NewService(userRepo, client)

	r := mux.NewRouter()

	service.InitRoutes(r)

	log.Fatal(http.ListenAndServe(":7001", r))
}
