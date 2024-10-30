package db

import (
	"context"
	"errors"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoConfig holds the MongoDB connection settings
type MongoConfig struct {
	URI  string
	Name string
}

var Client *mongo.Client

// InitMongoDB initializes the MongoDB client and returns the database instance
func InitMongoDB(config MongoConfig) (*mongo.Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Use mongo.Connect instead of mongo.NewClient and Client.Connect
	var err error
	Client, err = mongo.Connect(ctx, options.Client().ApplyURI(config.URI))
	if err != nil {
		return nil, errors.New("failed to connect to MongoDB: " + err.Error())
	}

	// Verify the connection by pinging the MongoDB server
	err = Client.Ping(ctx, nil)
	if err != nil {
		return nil, errors.New("failed to ping MongoDB: " + err.Error())
	}

	log.Println("Connected to MongoDB!")

	// Return the database instance
	return Client.Database(config.Name), nil
}
