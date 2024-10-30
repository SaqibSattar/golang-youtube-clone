package user

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repository *UserRepository
}

// NewUserService creates a new UserService
func NewUserService(repo *UserRepository) *UserService {
	return &UserService{repository: repo}
}

// Register creates a new user after validating and hashing the password
func (service *UserService) Register(user *User) error {
	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	// Save user in the database
	return service.repository.Create(user)
}

// Login verifies the user credentials
func (service *UserService) Login(email, password string) (*User, error) {
	user, err := service.repository.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	// Compare hashed password with provided password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, err // Invalid password
	}

	return user, nil // Return user if credentials are valid
}

// GetUserByID retrieves a user by their ID
func (service *UserService) GetUserByID(id primitive.ObjectID) (*User, error) {
	return service.repository.FindByID(id)
}

// UpdateUser updates user information
func (service *UserService) UpdateUser(user *User) error {
	return service.repository.Update(user)
}

// DeleteUser deletes a user by their ID
func (service *UserService) DeleteUser(id primitive.ObjectID) error {
	return service.repository.Delete(id)
}
