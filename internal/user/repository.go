package user

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database) *UserRepository {
	return &UserRepository{
		collection: db.Collection("users"),
	}
}

// Create inserts a new user into the database
func (repo *UserRepository) Create(user *User) error {
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	_, err := repo.collection.InsertOne(context.Background(), user)
	return err
}

// FindByUsername retrieves a user by username
func (repo *UserRepository) FindByUsername(username string) (*User, error) {
	var user User
	err := repo.collection.FindOne(context.Background(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// FindByEmail retrieves a user by email
func (repo *UserRepository) FindByEmail(email string) (*User, error) {
	var user User
	err := repo.collection.FindOne(context.Background(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// FindByID retrieves a user by their ID
func (repo *UserRepository) FindByID(id primitive.ObjectID) (*User, error) {
	var user User
	err := repo.collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Update updates user information in the database
func (repo *UserRepository) Update(user *User) error {
	user.UpdatedAt = time.Now()
	_, err := repo.collection.UpdateOne(
		context.Background(),
		bson.M{"_id": user.ID},
		bson.M{"$set": user},
	)
	return err
}

// Delete removes a user from the database
func (repo *UserRepository) Delete(id primitive.ObjectID) error {
	_, err := repo.collection.DeleteOne(context.Background(), bson.M{"_id": id})
	return err
}

// GetAll retrieves all users from the database
func (repo *UserRepository) GetAll() ([]User, error) {
	var users []User
	cursor, err := repo.collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var user User
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
