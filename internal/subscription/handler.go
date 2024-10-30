package subscription

import (
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/gorilla/mux"
)

var subscriptionService *SubscriptionService

func InitSubscriptionService(service *SubscriptionService) {
	subscriptionService = service
}

func SubscribeHandler(db *mongo.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		subscriberID, err := primitive.ObjectIDFromHex(vars["subscriberID"])
		if err != nil {
			http.Error(w, "Invalid subscriber ID", http.StatusBadRequest)
			return
		}
		channelID, err := primitive.ObjectIDFromHex(vars["channelID"])
		if err != nil {
			http.Error(w, "Invalid channel ID", http.StatusBadRequest)
			return
		}

		err = subscriptionService.Subscribe(subscriberID, channelID)
		if err != nil {
			http.Error(w, "Could not subscribe", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode("Subscribed successfully")
	}
}

func UnsubscribeHandler(db *mongo.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		subscriptionID, err := primitive.ObjectIDFromHex(vars["subscriptionID"])
		if err != nil {
			http.Error(w, "Invalid subscription ID", http.StatusBadRequest)
			return
		}

		err = subscriptionService.Unsubscribe(subscriptionID)
		if err != nil {
			http.Error(w, "Could not unsubscribe", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

// Additional handlers can be added as needed
