package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

// Init Posts var as a slice Post struct
var posts []Post

// Post Structs similar
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
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

// Get Single Post
func getPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get params
	// Find id

	for _, item := range posts {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Post{})
}

// Create a New Post
func createPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var post Post
	_ = json.NewDecoder(r.Body).Decode(&post)
	post.ID = strconv.Itoa(rand.Intn(1000000)) // mock id
	posts = append(posts, post)
	json.NewEncoder(w).Encode(post)
}

// Update Post
func updatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range posts {
		if item.ID == params["id"] {
			posts = append(posts[:index], posts[index+1:]...)
			var post Post
			_ = json.NewDecoder(r.Body).Decode(&post)
			post.ID = params["id"]
			posts = append(posts, post)
			json.NewEncoder(w).Encode(post)
			return
		}
	}
	json.NewEncoder(w).Encode(posts)
}

// Delete Post
func deletePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range posts {
		if item.ID == params["id"] {
			posts = append(posts[:index], posts[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(posts)
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