package subscription

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Subscription represents a subscription relationship between two users
type Subscription struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Subscriber primitive.ObjectID `bson:"subscriber"` // The subscriber (who is subscribing)
	Channel    primitive.ObjectID `bson:"channel"`    // The channel (to whom the subscriber is subscribing)
	CreatedAt  primitive.DateTime `bson:"createdAt"`
	UpdatedAt  primitive.DateTime `bson:"updatedAt"`
}
