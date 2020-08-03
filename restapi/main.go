package main

import(
	"encoding/json"
	"log"
	"net/http"
	"math/rand"
	"strconv"
	"github.com/gorilla/mux"
	)

// Post Structs (Model)
type Post struct {
	ID 			string 	`json:"id"`
	Title 		string 	`json:"title"`
	Body 		string 	`json:"body"`
	User 		*User 	`json:"user"`
}

func main() {
	// Init Router
	router := mux.NewRouter()

	// Route Handlers / Endpoints
	router.HandleFunc("/api/posts", getPosts).Methods("GET")
	router.HandleFunc("/api/posts/{id}", getPost).Methods("GET")
	router.HandleFunc("/api/posts", createPost).Methods("POST")
	router.HandleFunc("/api/posts/{id}", updatePost).Methods("PUT")
	router.HandleFunc("/api/posts/{id}", deletePost).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))

}