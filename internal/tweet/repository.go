package tweet

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TweetRepository struct {
	collection *mongo.Collection
}

func NewTweetRepository(db *mongo.Database) *TweetRepository {
	return &TweetRepository{
		collection: db.Collection("tweets"),
	}
}

func (r *TweetRepository) Create(tweet *Tweet) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := r.collection.InsertOne(ctx, tweet)
	return err
}

func (r *TweetRepository) FindByID(id primitive.ObjectID) (*Tweet, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var tweet Tweet
	err := r.collection.FindOne(ctx, Tweet{ID: id}).Decode(&tweet)
	if err != nil {
		return nil, err
	}
	return &tweet, nil
}

func (r *TweetRepository) Update(tweet *Tweet) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Create an update document
	update := bson.M{
		"$set": bson.M{
			"content":   tweet.Content,
			"updatedAt": time.Now(), // You might want to update the timestamp here
		},
	}

	// Update the tweet in the collection
	_, err := r.collection.UpdateOne(ctx, bson.M{"_id": tweet.ID}, update)
	return err
}

func (r *TweetRepository) Delete(id primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := r.collection.DeleteOne(ctx, Tweet{ID: id})
	return err
}

// Add more methods as needed (like GetAll, FindByUser, etc.)
