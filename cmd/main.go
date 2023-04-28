package main

import (
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"user_service/internal/db"
	"user_service/internal/db/user"
	"user_service/internal/snowflake"
	"user_service/internal/user_service"
)

func main() {
	err := godotenv.Load("config/config.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	addr := os.Getenv("SNOWFLAKE_SERVICE_ADDR")
	if addr == "" {
		log.Fatalf("empty SNOWFLAKE_SERVICE_ADDR")
	}

	conn, err := snowflake.NewConnection(addr)
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
