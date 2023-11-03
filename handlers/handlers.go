package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/benrobinson-ms/go-blog-api/models"
	"github.com/benrobinson-ms/go-blog-api/storage"
	"github.com/gorilla/mux"
)

// GetArticles handles the retrieval of all articles.
func GetArticles(w http.ResponseWriter, r *http.Request) {
	articles, err := storage.GetAllArticles()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(articles)
}

// CreateArticle handles the creation of a new article.
func CreateArticle(w http.ResponseWriter, r *http.Request) {
	var article models.Article
	err := json.NewDecoder(r.Body).Decode(&article)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = storage.CreateArticle(&article)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Article created successfully", "articleID": strconv.Itoa(int(article.ID))})
}

// GetArticle handles the retrieval of a specific article by ID.
func GetArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid article ID", http.StatusBadRequest)
		return
	}
	article, err := storage.GetArticleByID(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(article)
}

// UpdateArticle handles updating a specific article by ID.
func UpdateArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid article ID", http.StatusBadRequest)
		return
	}

	var updatedArticle models.Article
	err = json.NewDecoder(r.Body).Decode(&updatedArticle)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	updatedArticle.ID = uint(id) // Set the ID for the updatedArticle
	err = storage.UpdateArticle(&updatedArticle)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Article updated successfully"})
}

// DeleteArticle handles the deletion of a specific article by ID.
func DeleteArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid article ID", http.StatusBadRequest)
		return
	}
	err = storage.DeleteArticleByID(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Article deleted successfully"})
}
