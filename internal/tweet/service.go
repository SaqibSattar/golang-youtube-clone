package tweet

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TweetService struct {
	repository *TweetRepository
}

func NewTweetService(repo *TweetRepository) *TweetService {
	return &TweetService{repository: repo}
}

func (s *TweetService) CreateTweet(owner primitive.ObjectID, content string) error {
	tweet := &Tweet{
		Owner:   owner,
		Content: content,
	}

	return s.repository.Create(tweet)
}

func (s *TweetService) GetTweetByID(id primitive.ObjectID) (*Tweet, error) {
	return s.repository.FindByID(id)
}

func (s *TweetService) DeleteTweet(id primitive.ObjectID) error {
	return s.repository.Delete(id)
}

func (s *TweetService) EditTweet(id primitive.ObjectID, content string) error {
	tweet := &Tweet{
		ID:        id,
		Content:   content,
		UpdatedAt: primitive.NewDateTimeFromTime(time.Now()),
	}

	return s.repository.Update(tweet)
}

// Add more methods as needed (like GetAll, FindByUser, etc.)
