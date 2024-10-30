package user

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"youtube-backend/configs"
	"youtube-backend/internal/common"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// UserService can be injected here, assume it's defined somewhere
var userService *UserService

// Initialize the user service (usually done in main.go or through dependency injection)
// Initialize the user service (usually done in main.go or through dependency injection)
func InitUserService(service *UserService) {
	userService = service
}

// RegisterHandler handles user registration
func RegisterHandler(db *mongo.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		println("RegisterHandler 1")

		// Parse the multipart form data
		if err := r.ParseMultipartForm(10 << 20); err != nil { // 10 MB limit
			common.RespondError(w, common.ApiError{
				StatusCode: http.StatusBadRequest,
				Message:    "Unable to parse form",
				Errors:     nil,
				Success:    false,
			})
			return
		}
		println("RegisterHandler 2")

		// Retrieve user details from form
		fullName := r.FormValue("fullName")
		email := r.FormValue("email")
		username := r.FormValue("username")
		password := r.FormValue("password")

		// Validate required fields
		if fullName == "" || email == "" || username == "" || password == "" {
			common.RespondError(w, common.ApiError{
				StatusCode: http.StatusBadRequest,
				Message:    "All fields are required",
				Errors:     nil,
				Success:    false,
			})
			return
		}

		println("RegisterHandler 3")

		// Check for existing user
		existedUser, err := userService.repository.FindByEmail(email)
		if err == nil && existedUser != nil {
			common.RespondError(w, common.ApiError{
				StatusCode: http.StatusConflict,
				Message:    "User with email already exists",
				Errors:     nil,
				Success:    false,
			})
			return
		}
		println("RegisterHandler 4")

		existedUser, err = userService.repository.FindByUsername(username)
		if err == nil && existedUser != nil {
			common.RespondError(w, common.ApiError{
				StatusCode: http.StatusConflict,
				Message:    "User with username already exists",
				Errors:     nil,
				Success:    false,
			})
			return
		}
		println("RegisterHandler 5")

		// Handle file uploads for avatar and cover image
		var avatarURL, coverImageURL string

		// Upload avatar
		if avatarFile, _, err := r.FormFile("avatar"); err == nil {
			defer avatarFile.Close() // Close file after reading
			avatarURL, err = uploadToCloudinary(avatarFile)
			if err != nil {
				common.RespondError(w, common.ApiError{
					StatusCode: http.StatusInternalServerError,
					Message:    "Failed to upload avatar",
					Errors:     nil,
					Success:    false,
				})
				return
			}
		}
		println("RegisterHandler 6")

		// Upload cover image
		if coverImageFile, _, err := r.FormFile("coverImage"); err == nil {
			defer coverImageFile.Close() // Close file after reading
			coverImageURL, err = uploadToCloudinary(coverImageFile)
			if err != nil {
				common.RespondError(w, common.ApiError{
					StatusCode: http.StatusInternalServerError,
					Message:    "Failed to upload cover image",
					Errors:     nil,
					Success:    false,
				})
				return
			}
		}
		println("RegisterHandler 7")

		// Create the user
		user := User{
			FullName:   fullName,
			Email:      email,
			Username:   username,
			Password:   password,
			Avatar:     avatarURL,
			CoverImage: coverImageURL,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		}

		err = userService.Register(&user)
		if err != nil {
			common.RespondError(w, common.ApiError{
				StatusCode: http.StatusInternalServerError,
				Message:    "Failed to register user",
				Errors:     nil,
				Success:    false,
			})
			return
		}
		println("RegisterHandler 8")

		// Send successful response
		common.Respond(w, http.StatusCreated, user, "User registered successfully")
	}
}

// LoginHandler handles user login
func LoginHandler(db *mongo.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var loginInfo struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		if err := json.NewDecoder(r.Body).Decode(&loginInfo); err != nil {
			common.RespondError(w, common.ApiError{
				StatusCode: http.StatusBadRequest,
				Message:    err.Error(),
				Errors:     nil,
				Success:    false,
			})
			return
		}

		user, err := userService.Login(loginInfo.Email, loginInfo.Password)
		if err != nil {
			common.RespondError(w, common.ApiError{
				StatusCode: http.StatusUnauthorized,
				Message:    "Invalid email or password",
				Errors:     nil,
				Success:    false,
			})
			return
		}

		// Generate JWT token
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"id":       user.ID,
			"username": user.Username,
			"exp":      time.Now().Add(time.Duration(configs.JWTExpiration)).Unix(),
		})

		// Sign the token with your secret
		tokenString, err := token.SignedString([]byte(configs.JWTSecret))
		if err != nil {
			common.RespondError(w, common.ApiError{
				StatusCode: http.StatusInternalServerError,
				Message:    "Could not create token",
				Errors:     nil,
				Success:    false,
			})
			return
		}

		// Example response
		response := map[string]interface{}{
			"user":  user,
			"token": tokenString,
		}

		common.Respond(w, http.StatusOK, response, "Login successful")
	}
}

// GetUserHandler retrieves a user by their ID
func GetUserHandler(db *mongo.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		params := mux.Vars(r)
		userID, err := primitive.ObjectIDFromHex(params["id"])
		if err != nil {
			common.RespondError(w, common.ApiError{
				StatusCode: http.StatusBadRequest,
				Message:    "Invalid user ID",
				Errors:     nil,
				Success:    false,
			})
			return
		}

		user, err := userService.GetUserByID(userID)
		if err != nil {
			common.RespondError(w, common.ApiError{
				StatusCode: http.StatusNotFound,
				Message:    "User not found",
				Errors:     nil,
				Success:    false,
			})
			return
		}

		common.Respond(w, http.StatusOK, user, "User retrieved successfully")
	}
}

// UpdateUserHandler updates user information
func UpdateUserHandler(db *mongo.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		userID, err := primitive.ObjectIDFromHex(params["id"])
		if err != nil {
			common.RespondError(w, common.ApiError{
				StatusCode: http.StatusBadRequest,
				Message:    "Invalid user ID",
				Errors:     nil,
				Success:    false,
			})
			return
		}

		var userUpdates User
		if err := json.NewDecoder(r.Body).Decode(&userUpdates); err != nil {
			common.RespondError(w, common.ApiError{
				StatusCode: http.StatusBadRequest,
				Message:    "Invalid request body",
				Errors:     nil,
				Success:    false,
			})
			return
		}

		userUpdates.ID = userID // Set the ID of the user being updated

		err = userService.UpdateUser(&userUpdates)
		if err != nil {
			common.RespondError(w, common.ApiError{
				StatusCode: http.StatusInternalServerError,
				Message:    "Failed to update user",
				Errors:     nil,
				Success:    false,
			})
			return
		}

		common.Respond(w, http.StatusOK, userUpdates, "User updated successfully")
	}
}

// DeleteUserHandler deletes a user by their ID
func DeleteUserHandler(db *mongo.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		userID, err := primitive.ObjectIDFromHex(params["id"])
		if err != nil {
			common.RespondError(w, common.ApiError{
				StatusCode: http.StatusBadRequest,
				Message:    "Invalid user ID",
				Errors:     nil,
				Success:    false,
			})
			return
		}

		err = userService.DeleteUser(userID)
		if err != nil {
			common.RespondError(w, common.ApiError{
				StatusCode: http.StatusInternalServerError,
				Message:    "Failed to delete user",
				Errors:     nil,
				Success:    false,
			})
			return
		}

		common.Respond(w, http.StatusOK, nil, "User deleted successfully")
	}
}

// GetAllUsersHandler retrieves all users
func GetAllUsersHandler(db *mongo.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := userService.repository.GetAll()
		if err != nil {
			common.RespondError(w, common.ApiError{
				StatusCode: http.StatusInternalServerError,
				Message:    "Failed to retrieve users",
				Errors:     nil,
				Success:    false,
			})
			return
		}

		common.Respond(w, http.StatusOK, users, "Users retrieved successfully")
	}
}

// uploadToCloudinary handles file upload to Cloudinary (Assuming this function is defined)
func uploadToCloudinary(file io.Reader) (string, error) {
	// Initialize Cloudinary client
	cld, err := cloudinary.NewFromParams(configs.CloudinaryCloudName, configs.CloudinaryAPIKey, configs.CloudinaryAPISecret)
	if err != nil {
		log.Printf("Error initializing Cloudinary: %v", err)
		return "", err
	}

	// Upload file
	resp, err := cld.Upload.Upload(context.Background(), file, uploader.UploadParams{Folder: "your_folder"})
	if err != nil {
		log.Printf("Error uploading file to Cloudinary: %v", err)
		return "", err
	}

	log.Printf("File uploaded to Cloudinary: %s", resp.SecureURL)
	return resp.SecureURL, nil
}
