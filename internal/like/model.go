package like

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Like struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Video     primitive.ObjectID `bson:"video" json:"video,omitempty"`
	Comment   primitive.ObjectID `bson:"comment" json:"comment,omitempty"`
	Tweet     primitive.ObjectID `bson:"tweet" json:"tweet,omitempty"`
	LikedBy   primitive.ObjectID `bson:"likedBy" json:"likedBy"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}
