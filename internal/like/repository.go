package like

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type LikeRepository struct {
	collection *mongo.Collection
}

func NewLikeRepository(db *mongo.Database) *LikeRepository {
	return &LikeRepository{
		collection: db.Collection("likes"),
	}
}

func (r *LikeRepository) Create(like *Like) error {
	like.CreatedAt = time.Now()
	like.UpdatedAt = time.Now()
	_, err := r.collection.InsertOne(context.TODO(), like)
	return err
}

func (r *LikeRepository) Delete(id primitive.ObjectID) error {
	_, err := r.collection.DeleteOne(context.TODO(), map[string]interface{}{"_id": id})
	return err
}

func (r *LikeRepository) GetByID(id primitive.ObjectID) (*Like, error) {
	var like Like
	err := r.collection.FindOne(context.TODO(), map[string]interface{}{"_id": id}).Decode(&like)
	return &like, err
}
