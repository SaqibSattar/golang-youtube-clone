package video

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Video represents the video model
type Video struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	VideoFile   string             `bson:"videoFile" json:"videoFile"`
	Thumbnail   string             `bson:"thumbnail" json:"thumbnail"`
	Title       string             `bson:"title" json:"title"`
	Description string             `bson:"description" json:"description"`
	Duration    int                `bson:"duration" json:"duration"`
	Views       int                `bson:"views" json:"views"`
	IsPublished bool               `bson:"isPublished" json:"isPublished"`
	Owner       primitive.ObjectID `bson:"owner" json:"owner"`
	CreatedAt   time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt   time.Time          `bson:"updatedAt" json:"updatedAt"`
}
