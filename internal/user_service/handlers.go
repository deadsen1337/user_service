package user_service

import (
	"context"
	"encoding/json"
	"google.golang.org/grpc"
	"io"
	"log"
	"net/http"
	"strconv"
	"user_service/internal/model"
	snowflake_api_service "user_service/pkg/snowflake-api"
)

func (s *UserSevice) Get(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	user, err := s.userRepo.Get([]byte(id))
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(user)
}

func (s *UserSevice) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	err := s.userRepo.Delete([]byte(id))
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (s *UserSevice) Create(w http.ResponseWriter, r *http.Request) {
	userData, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var check model.User

	err = json.Unmarshal(userData, &check)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp, err := s.snowflakeClient.NewID(context.Background(), &snowflake_api_service.IdRequest{}, []grpc.CallOption{}...)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if resp.GetId() == 0 {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	id := []byte(strconv.FormatUint(resp.GetId(), 10))

	err = s.userRepo.Create(id, userData)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(id)
}

func (s *UserSevice) Update(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	userData, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var check model.User

	err = json.Unmarshal(userData, &check)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = s.userRepo.Update([]byte(id), userData)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}