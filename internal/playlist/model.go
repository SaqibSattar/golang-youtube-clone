// internal/playlist/model.go

package playlist

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Playlist struct {
	ID          primitive.ObjectID   `bson:"_id,omitempty" json:"id"`
	Name        string               `bson:"name" json:"name" binding:"required"`
	Description string               `bson:"description" json:"description" binding:"required"`
	Videos      []primitive.ObjectID `bson:"videos" json:"videos"`
	Owner       primitive.ObjectID   `bson:"owner" json:"owner"`
	CreatedAt   time.Time            `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time            `bson:"updated_at" json:"updated_at"`
}
