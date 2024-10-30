package like

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var likeService *LikeService

// InitLikeService initializes the Like service
func InitLikeService(service *LikeService) {
	likeService = service
}

// CreateLikeHandler handles the creation of a new like
func CreateLikeHandler(db *mongo.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var like Like
		if err := json.NewDecoder(r.Body).Decode(&like); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := likeService.Create(&like); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(like)
	}
}

// DeleteLikeHandler handles the deletion of a like
func DeleteLikeHandler(db *mongo.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := primitive.ObjectIDFromHex(vars["id"])
		if err != nil {
			http.Error(w, "Invalid ID format", http.StatusBadRequest)
			return
		}

		if err := likeService.Delete(id); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

// GetLikeHandler retrieves a like by ID
func GetLikeHandler(db *mongo.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := primitive.ObjectIDFromHex(vars["id"])
		if err != nil {
			http.Error(w, "Invalid ID format", http.StatusBadRequest)
			return
		}

		like, err := likeService.GetByID(id)
		if err != nil {
			http.Error(w, "Like not found", http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(like)
	}
}
