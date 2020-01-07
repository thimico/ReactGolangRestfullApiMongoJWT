package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/gorilla/mux"

	"ReactGolangRestfullApiMongoJWT/src/apis/jwtauth"
	"ReactGolangRestfullApiMongoJWT/src/apis/userapi"
	"ReactGolangRestfullApiMongoJWT/src/apis/profileapi"
)

func main() {
	r := mux.NewRouter()

	// Route handles & endpoints
	r.HandleFunc("/api/users/login", jwtauth.GenerateToken).Methods("POST")
	r.HandleFunc("/api/users/register", userapi.Create).Methods("POST")
	r.HandleFunc("/api/users/test", userapi.Test).Methods("GET")
	r.HandleFunc("/api/users/current", userapi.Current).Methods("GET")
	r.HandleFunc("/api/profile/test", profileapi.Test).Methods("GET")

	r.HandleFunc("/api/profile/all", profileapi.FindAll).Methods("GET")
	r.HandleFunc("/api/profile", profileapi.Current).Methods("GET")
	r.HandleFunc("/api/profile/{id}", profileapi.Find).Methods("GET")
	r.HandleFunc("/api/profile/handle/{handle}", profileapi.Handle).Methods("GET")
	r.HandleFunc("/api/profile", profileapi.Create).Methods("POST")
	r.HandleFunc("/api/profile/{id}", profileapi.Update).Methods("PUT")
	r.HandleFunc("/api/profile/{id}", profileapi.Delete).Methods("DELETE")

	p := os.Getenv("PORT")
	if p == "" {
		p = "8000"
		log.Printf("Defaulting to port %s", p)
	}

	log.Printf("Listening on port %s", p)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", p), r))
}
