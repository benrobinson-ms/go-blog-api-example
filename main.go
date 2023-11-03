package main

import (
	"net/http"

	"github.com/benrobinson-ms/go-blog-api/database"
	"github.com/benrobinson-ms/go-blog-api/handlers"
	"github.com/gorilla/mux"
)

func main() {
	// Initialize the database
	database.InitDB()
	defer database.CloseDB()

	r := mux.NewRouter()

	// Retrieving all articles
	r.HandleFunc("/articles", handlers.GetArticles).Methods("GET")

	// Creating a new article
	r.HandleFunc("/articles", handlers.CreateArticle).Methods("POST")

	// Retrieving a specific article by ID
	r.HandleFunc("/articles/{id:[0-9]+}", handlers.GetArticle).Methods("GET")

	// Updating a specific article by ID
	r.HandleFunc("/articles/{id:[0-9]+}", handlers.UpdateArticle).Methods("PUT")

	// Deleting a specific article by ID
	r.HandleFunc("/articles/{id:[0-9]+}", handlers.DeleteArticle).Methods("DELETE")

	http.ListenAndServe(":8080", r)
}
