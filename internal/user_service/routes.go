package user_service

import "github.com/gorilla/mux"

func (s *UserSevice) InitRoutes(r *mux.Router) {
	r.HandleFunc("/users/{id}", s.Get).Methods("GET")
	r.HandleFunc("/users", s.Create).Methods("POST")
	r.HandleFunc("/users/{id}", s.Update).Methods("PUT")
	r.HandleFunc("/users/{id}", s.Delete).Methods("DELETE")
}
