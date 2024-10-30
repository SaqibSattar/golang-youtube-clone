package video

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// VideoService can be injected here, assume it's defined somewhere
var videoService *VideoService

// Initialize the video service (usually done in main.go or through dependency injection)
func InitVideoService(service *VideoService) {
	videoService = service
}

// CreateVideoHandler handles video creation
func CreateVideoHandler(db *mongo.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var video Video
		if err := json.NewDecoder(r.Body).Decode(&video); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		video.CreatedAt = time.Now()
		video.UpdatedAt = time.Now()

		err := videoService.CreateVideo(&video)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(video)
	}
}

// GetVideoHandler retrieves a video by ID
func GetVideoHandler(db *mongo.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := primitive.ObjectIDFromHex(vars["id"])
		if err != nil {
			http.Error(w, "Invalid ID format", http.StatusBadRequest)
			return
		}

		video, err := videoService.GetVideoByID(id)
		if err != nil {
			http.Error(w, "Video not found", http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(video)
	}
}

// UpdateVideoHandler updates a video
func UpdateVideoHandler(db *mongo.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := primitive.ObjectIDFromHex(vars["id"])
		if err != nil {
			http.Error(w, "Invalid ID format", http.StatusBadRequest)
			return
		}

		var video Video
		if err := json.NewDecoder(r.Body).Decode(&video); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		video.ID = id
		video.UpdatedAt = time.Now()

		err = videoService.UpdateVideo(&video)
		if err != nil {
			http.Error(w, "Failed to update video", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(video)
	}
}

// DeleteVideoHandler deletes a video
func DeleteVideoHandler(db *mongo.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := primitive.ObjectIDFromHex(vars["id"])
		if err != nil {
			http.Error(w, "Invalid ID format", http.StatusBadRequest)
			return
		}

		err = videoService.DeleteVideo(id)
		if err != nil {
			http.Error(w, "Failed to delete video", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

// GetAllVideosHandler retrieves all videos
func GetAllVideosHandler(db *mongo.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		videos, err := videoService.GetAllVideos()
		if err != nil {
			http.Error(w, "Failed to retrieve videos", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(videos)
	}
}
