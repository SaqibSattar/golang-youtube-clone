// internal/playlist/repository.go

package playlist

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PlaylistRepository struct {
	collection *mongo.Collection
}

func NewPlaylistRepository(db *mongo.Database) *PlaylistRepository {
	return &PlaylistRepository{
		collection: db.Collection("playlists"),
	}
}

func (r *PlaylistRepository) Create(playlist *Playlist) error {
	playlist.CreatedAt = time.Now()
	playlist.UpdatedAt = time.Now()
	_, err := r.collection.InsertOne(context.TODO(), playlist)
	return err
}

func (r *PlaylistRepository) Update(playlist *Playlist) error {
	playlist.UpdatedAt = time.Now()
	_, err := r.collection.UpdateOne(context.TODO(),
		map[string]interface{}{"_id": playlist.ID},
		map[string]interface{}{"$set": playlist})
	return err
}

func (r *PlaylistRepository) Delete(id primitive.ObjectID) error {
	_, err := r.collection.DeleteOne(context.TODO(), map[string]interface{}{"_id": id})
	return err
}

func (r *PlaylistRepository) GetByID(id primitive.ObjectID) (*Playlist, error) {
	var playlist Playlist
	err := r.collection.FindOne(context.TODO(), map[string]interface{}{"_id": id}).Decode(&playlist)
	return &playlist, err
}

func (r *PlaylistRepository) GetAll() ([]Playlist, error) {
	var playlists []Playlist
	cursor, err := r.collection.Find(context.TODO(), map[string]interface{}{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var playlist Playlist
		if err = cursor.Decode(&playlist); err != nil {
			return nil, err
		}
		playlists = append(playlists, playlist)
	}

	return playlists, cursor.Err()
}
