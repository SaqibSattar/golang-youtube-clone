package user

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User represents the user model
type User struct {
	ID           primitive.ObjectID   `bson:"_id,omitempty"`
	Username     string               `bson:"username" validate:"required,lowercase,unique"`
	Email        string               `bson:"email" validate:"required,email,lowercase,unique"`
	FullName     string               `bson:"full_name" validate:"required"`
	Avatar       string               `bson:"avatar" validate:"required,url"` // Cloudinary URL
	CoverImage   string               `bson:"cover_image,omitempty"`          // Optional Cloudinary URL
	WatchHistory []primitive.ObjectID `bson:"watch_history,omitempty"`        // References to videos
	Password     string               `bson:"password" validate:"required,min=8"`
	RefreshToken string               `bson:"refresh_token,omitempty"` // Token for JWT refresh flow
	CreatedAt    time.Time            `bson:"created_at"`
	UpdatedAt    time.Time            `bson:"updated_at"`
}
