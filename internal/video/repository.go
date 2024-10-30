package video

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// VideoRepository defines the repository interface for videos
type VideoRepository struct {
	collection *mongo.Collection
}

// NewVideoRepository creates a new VideoRepository
func NewVideoRepository(db *mongo.Database) *VideoRepository {
	return &VideoRepository{
		collection: db.Collection("videos"),
	}
}

// Create adds a new video to the database
func (repo *VideoRepository) Create(video *Video) error {
	_, err := repo.collection.InsertOne(context.TODO(), video)
	return err
}

// FindByID retrieves a video by its ID
func (repo *VideoRepository) FindByID(id primitive.ObjectID) (*Video, error) {
	var video Video
	err := repo.collection.FindOne(context.TODO(), Video{ID: id}).Decode(&video)
	return &video, err
}

// Update updates an existing video
func (repo *VideoRepository) Update(video *Video) error {
	_, err := repo.collection.UpdateOne(context.TODO(), Video{ID: video.ID}, video)
	return err
}

// Delete removes a video from the database
func (repo *VideoRepository) Delete(id primitive.ObjectID) error {
	_, err := repo.collection.DeleteOne(context.TODO(), Video{ID: id})
	return err
}

// GetAll retrieves all videos from the database
func (repo *VideoRepository) GetAll() ([]Video, error) {
	var videos []Video
	cursor, err := repo.collection.Find(context.TODO(), nil)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var video Video
		if err := cursor.Decode(&video); err != nil {
			return nil, err
		}
		videos = append(videos, video)
	}

	return videos, nil
}
