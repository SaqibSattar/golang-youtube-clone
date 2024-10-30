package configs

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	MongoDBURI          string
	MongoDBName         string
	JWTSecret           string
	ServerAddress       string
	JWTExpiration       int64 // Change this to int64
	CORSAllowOrigin     string
	LogLevel            string
	CloudinaryCloudName string
	CloudinaryAPIKey    string
	CloudinaryAPISecret string
)

func LoadConfig() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Set configuration variables
	MongoDBURI = os.Getenv("MONGO_DB_URI")
	MongoDBName = os.Getenv("MONGO_DB_NAME")
	JWTSecret = os.Getenv("JWT_SECRET")
	ServerAddress = os.Getenv("SERVER_ADDRESS")
	CloudinaryCloudName = os.Getenv("CLOUDINARY_CLOUD_NAME")
	CloudinaryAPIKey = os.Getenv("CLOUDINARY_API_KEY")
	CloudinaryAPISecret = os.Getenv("CLOUDINARY_API_SECRET")

	// Log loaded configurations for debugging
	log.Printf("Loaded configurations: MongoDBURI: %s, ServerAddress: %s", MongoDBURI, ServerAddress)

	// Parse JWT_EXPIRATION as int64
	expiration, err := strconv.ParseInt(os.Getenv("JWT_EXPIRATION"), 10, 64)
	if err != nil {
		JWTExpiration = 3600 // Default to 1 hour if not set
	} else {
		JWTExpiration = expiration
	}

	CORSAllowOrigin = os.Getenv("CORS_ALLOW_ORIGIN")
	LogLevel = os.Getenv("LOG_LEVEL")
}
