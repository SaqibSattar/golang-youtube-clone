package subscription

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type SubscriptionRepository struct {
	collection *mongo.Collection
}

func NewSubscriptionRepository(db *mongo.Database) *SubscriptionRepository {
	return &SubscriptionRepository{
		collection: db.Collection("subscriptions"),
	}
}

func (r *SubscriptionRepository) Create(subscription *Subscription) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := r.collection.InsertOne(ctx, subscription)
	return err
}

func (r *SubscriptionRepository) FindBySubscriberAndChannel(subscriberID, channelID primitive.ObjectID) (*Subscription, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var subscription Subscription
	err := r.collection.FindOne(ctx, Subscription{Subscriber: subscriberID, Channel: channelID}).Decode(&subscription)
	if err != nil {
		return nil, err
	}
	return &subscription, nil
}

func (r *SubscriptionRepository) Delete(subscriptionID primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := r.collection.DeleteOne(ctx, Subscription{ID: subscriptionID})
	return err
}

// Add more methods as needed (like GetAll, FindByUser, etc.)
