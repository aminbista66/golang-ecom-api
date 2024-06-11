package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/aminbista66/golang-ecom-api/service/user"
	"github.com/gorilla/mux"
)

type APIServer struct {
	address string
	db *sql.DB
}

func NewAPIServer(address string, db *sql.DB) *APIServer {
	return &APIServer{
		address: address,
		db: db,
	}
}

func Ping(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("pong"))
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	subrouter.HandleFunc("/ping", Ping).Methods("GET")

	userHandler := user.NewHandler()
	userHandler.RegisterRoutes(subrouter)

	log.Println("Listening on", s.address)
	
	return http.ListenAndServe(s.address, router)
}