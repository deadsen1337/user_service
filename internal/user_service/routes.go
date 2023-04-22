package user_service

import "github.com/gorilla/mux"

func (s *UserSevice) InitRoutes(r *mux.Router) {
	r.HandleFunc("/user/get", s.Get).Methods("GET")
	r.HandleFunc("/user/create", s.Create).Methods("POST")
	r.HandleFunc("/user/update", s.Update).Methods("PUT")
	r.HandleFunc("/user/delete", s.Delete).Methods("DELETE")
}
