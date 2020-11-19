package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

// Collection of Users
type User struct {
	FullName string `json:"fullname"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// Collection of Post
type Post struct {
	Title  string `json:"title"`
	Body   string `json:"body"`
	Author User   `json:"author"`
}

var post []Post = []Post{}

func addItem(w http.ResponseWriter, r *http.Request) {
	var newPost Post

	json.NewDecoder(r.Body).Decode(&newPost)
	post = append(post, newPost)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(post)
}

func JsonTest(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(struct {
		ID string
	}{"454-454-454"})
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/test", JsonTest)
	router.HandleFunc("/add", addItem).Methods("POST")

	http.ListenAndServe(":5000", router)

}
