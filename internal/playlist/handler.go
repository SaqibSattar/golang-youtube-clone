// internal/playlist/handler.go

package playlist

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var playlistService *PlaylistService

// InitPlaylistService initializes the Playlist service
func InitPlaylistService(service *PlaylistService) {
	playlistService = service
}

// CreatePlaylistHandler handles the creation of a new playlist
func CreatePlaylistHandler(db *mongo.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var playlist Playlist
		if err := json.NewDecoder(r.Body).Decode(&playlist); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := playlistService.Create(&playlist); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(playlist)
	}
}

// UpdatePlaylistHandler handles the updating of a playlist
func UpdatePlaylistHandler(db *mongo.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := primitive.ObjectIDFromHex(vars["id"])
		if err != nil {
			http.Error(w, "Invalid ID format", http.StatusBadRequest)
			return
		}

		var playlist Playlist
		if err := json.NewDecoder(r.Body).Decode(&playlist); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		playlist.ID = id

		if err := playlistService.Update(&playlist); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(playlist)
	}
}

// DeletePlaylistHandler handles the deletion of a playlist
func DeletePlaylistHandler(db *mongo.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := primitive.ObjectIDFromHex(vars["id"])
		if err != nil {
			http.Error(w, "Invalid ID format", http.StatusBadRequest)
			return
		}

		if err := playlistService.Delete(id); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

// GetPlaylistHandler retrieves a playlist by ID
func GetPlaylistHandler(db *mongo.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := primitive.ObjectIDFromHex(vars["id"])
		if err != nil {
			http.Error(w, "Invalid ID format", http.StatusBadRequest)
			return
		}

		playlist, err := playlistService.GetByID(id)
		if err != nil {
			http.Error(w, "Playlist not found", http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(playlist)
	}
}

// GetAllPlaylistsHandler retrieves all playlists
func GetAllPlaylistsHandler(db *mongo.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		playlists, err := playlistService.GetAll()
		if err != nil {
			http.Error(w, "Failed to retrieve playlists", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(playlists)
	}
}
