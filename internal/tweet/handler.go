package tweet

import (
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/gorilla/mux"
)

var tweetService *TweetService

func InitTweetService(service *TweetService) {
	tweetService = service
}

func CreateTweetHandler(db *mongo.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var tweetData struct {
			Content string `json:"content"`
		}

		if err := json.NewDecoder(r.Body).Decode(&tweetData); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Assume we extract user ID from context or JWT token
		// For example:
		// ownerID := getUserIDFromContext(r.Context())
		ownerID := primitive.NewObjectID() // Replace with actual user ID extraction logic

		err := tweetService.CreateTweet(ownerID, tweetData.Content)
		if err != nil {
			http.Error(w, "Could not create tweet", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode("Tweet created successfully")
	}
}

func GetTweetHandler(db *mongo.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := primitive.ObjectIDFromHex(vars["id"])
		if err != nil {
			http.Error(w, "Invalid tweet ID", http.StatusBadRequest)
			return
		}

		tweet, err := tweetService.GetTweetByID(id)
		if err != nil {
			http.Error(w, "Tweet not found", http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(tweet)
	}
}

func EditTweetHandler(db *mongo.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := primitive.ObjectIDFromHex(vars["id"])
		if err != nil {
			http.Error(w, "Invalid tweet ID", http.StatusBadRequest)
			return
		}

		var tweetData struct {
			Content string `json:"content"`
		}

		if err := json.NewDecoder(r.Body).Decode(&tweetData); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = tweetService.EditTweet(id, tweetData.Content)
		if err != nil {
			http.Error(w, "Could not update tweet", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("Tweet updated successfully")
	}
}

func DeleteTweetHandler(db *mongo.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := primitive.ObjectIDFromHex(vars["id"])
		if err != nil {
			http.Error(w, "Invalid tweet ID", http.StatusBadRequest)
			return
		}

		err = tweetService.DeleteTweet(id)
		if err != nil {
			http.Error(w, "Could not delete tweet", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

// Additional handlers can be added as needed
