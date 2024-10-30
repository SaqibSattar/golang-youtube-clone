package subscription

import "go.mongodb.org/mongo-driver/bson/primitive"

type SubscriptionService struct {
	repository *SubscriptionRepository
}

func NewSubscriptionService(repo *SubscriptionRepository) *SubscriptionService {
	return &SubscriptionService{repository: repo}
}

func (s *SubscriptionService) Subscribe(subscriberID, channelID primitive.ObjectID) error {
	subscription := &Subscription{
		Subscriber: subscriberID,
		Channel:    channelID,
	}

	return s.repository.Create(subscription)
}

func (s *SubscriptionService) Unsubscribe(subscriptionID primitive.ObjectID) error {
	return s.repository.Delete(subscriptionID)
}

// Add more methods as needed (like checking if a user is subscribed, getting all subscriptions, etc.)
