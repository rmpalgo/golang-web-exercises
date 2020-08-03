package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// Init Posts var as a slice Post struct
var posts []Post

// Post Structs
type Post struct {
	ID 			string 	`json:"id"`
	Title 		string 	`json:"title"`
	Body 		string 	`json:"body"`
	User 		*User 	`json:"user"`
}

// User Struct
type User struct {
	Firstname 	string 	`json:"firstname"`
	Lastname 	string 	`json:"lastname"`
}

// Get All Posts
func getPosts(w http.ResponseWriter, r *http.Request) {

}

// Get Single Post
func getPost(w http.ResponseWriter, r *http.Request) {

}

// Create a New Post
func createPost(w http.ResponseWriter, r *http.Request) {

}

// Update Book
func updatePost(w http.ResponseWriter, r *http.Request) {

}

// Delete Book
func deletePost(w http.ResponseWriter, r *http.Request) {

}

func main() {
	// Init Router
	router := mux.NewRouter()

	// Mock Posts
	posts = append(posts, Post{ID: "1", Title: "Post One", User: &User{Firstname: "Ron", Lastname: "Pal"}})
	posts = append(posts, Post{ID: "2", Title: "Post Two", User: &User{Firstname: "Milo", Lastname: "Puppy"}})

	// Route Handlers / Endpoints
	router.HandleFunc("/api/posts", getPosts).Methods("GET")
	router.HandleFunc("/api/posts/{id}", getPost).Methods("GET")
	router.HandleFunc("/api/posts", createPost).Methods("POST")
	router.HandleFunc("/api/posts/{id}", updatePost).Methods("PUT")
	router.HandleFunc("/api/posts/{id}", deletePost).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))

}