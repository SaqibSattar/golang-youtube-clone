package tweet

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Tweet represents a tweet created by a user
type Tweet struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Content   string             `bson:"content"` // The content of the tweet
	Owner     primitive.ObjectID `bson:"owner"`   // The user who owns the tweet
	CreatedAt primitive.DateTime `bson:"createdAt"`
	UpdatedAt primitive.DateTime `bson:"updatedAt"`
}
