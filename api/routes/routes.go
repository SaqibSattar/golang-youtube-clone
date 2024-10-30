package routes

import (
	"youtube-backend/internal/comment"
	"youtube-backend/internal/like"
	"youtube-backend/internal/playlist"
	"youtube-backend/internal/subscription"
	"youtube-backend/internal/tweet"
	"youtube-backend/internal/user"
	"youtube-backend/internal/video"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

// InitRoutes initializes all routes and injects dependencies
func InitRoutes(db *mongo.Database) *mux.Router {
	router := mux.NewRouter()

	// Initialize services for each module
	userRepo := user.NewUserRepository(db)       // Initialize UserRepository
	userService := user.NewUserService(userRepo) // Initialize UserService with the repository
	user.InitUserService(userService)            // Initialize the service globally

	// Video module initialization
	videoRepo := video.NewVideoRepository(db)        // Create a new VideoRepository
	videoService := video.NewVideoService(videoRepo) // Create a new VideoService using the repository
	video.InitVideoService(videoService)             // Initialize VideoService

	// Playlist module initialization
	playlistRepo := playlist.NewPlaylistRepository(db)
	playlistService := playlist.NewPlaylistService(playlistRepo)
	playlist.InitPlaylistService(playlistService) // Pass the service to the playlist module

	// Comment module initialization
	commentRepo := comment.NewCommentRepository(db)
	commentService := comment.NewCommentService(commentRepo)
	comment.InitCommentService(commentService)

	// Like module initialization
	likeRepo := like.NewLikeRepository(db)
	likeService := like.NewLikeService(likeRepo)
	like.InitLikeService(likeService) // Pass the service to the like module

	// Subscription module initialization
	subscriptionRepo := subscription.NewSubscriptionRepository(db)
	subscriptionService := subscription.NewSubscriptionService(subscriptionRepo)
	subscription.InitSubscriptionService(subscriptionService) // Pass the service to the subscription module

	// Tweet module initialization
	tweetRepo := tweet.NewTweetRepository(db)
	tweetService := tweet.NewTweetService(tweetRepo)
	tweet.InitTweetService(tweetService) // Pass the service to the tweet module

	// User routes
	userRoutes := router.PathPrefix("/api/v1/users").Subrouter()
	userRoutes.HandleFunc("/register", user.RegisterHandler(db)).Methods("POST")
	userRoutes.HandleFunc("/login", user.LoginHandler(db)).Methods("POST")
	userRoutes.HandleFunc("/{id}", user.GetUserHandler(db)).Methods("GET")
	userRoutes.HandleFunc("/{id}", user.UpdateUserHandler(db)).Methods("PUT")
	userRoutes.HandleFunc("/{id}", user.DeleteUserHandler(db)).Methods("DELETE")
	userRoutes.HandleFunc("/", user.GetAllUsersHandler(db)).Methods("GET")

	// Video routes
	videoRoutes := router.PathPrefix("/api/v1/videos").Subrouter()
	videoRoutes.HandleFunc("/", video.CreateVideoHandler(db)).Methods("POST")
	videoRoutes.HandleFunc("/{id}", video.GetVideoHandler(db)).Methods("GET")
	videoRoutes.HandleFunc("/{id}", video.UpdateVideoHandler(db)).Methods("PUT")
	videoRoutes.HandleFunc("/{id}", video.DeleteVideoHandler(db)).Methods("DELETE")
	videoRoutes.HandleFunc("/", video.GetAllVideosHandler(db)).Methods("GET")

	// Playlist routes
	playlistRoutes := router.PathPrefix("/api/v1/playlists").Subrouter()
	playlistRoutes.HandleFunc("/", playlist.CreatePlaylistHandler(db)).Methods("POST")
	playlistRoutes.HandleFunc("/{id}", playlist.GetPlaylistHandler(db)).Methods("GET")
	playlistRoutes.HandleFunc("/{id}", playlist.UpdatePlaylistHandler(db)).Methods("PUT")
	playlistRoutes.HandleFunc("/{id}", playlist.DeletePlaylistHandler(db)).Methods("DELETE")
	playlistRoutes.HandleFunc("/", playlist.GetAllPlaylistsHandler(db)).Methods("GET")

	// Comment routes
	commentRoutes := router.PathPrefix("/api/v1/comments").Subrouter()
	commentRoutes.HandleFunc("/", comment.CreateCommentHandler(db)).Methods("POST")
	commentRoutes.HandleFunc("/{id}", comment.GetCommentHandler(db)).Methods("GET")
	commentRoutes.HandleFunc("/{id}", comment.UpdateCommentHandler(db)).Methods("PUT")
	commentRoutes.HandleFunc("/{id}", comment.DeleteCommentHandler(db)).Methods("DELETE")
	commentRoutes.HandleFunc("/video/{videoId}", comment.GetAllCommentsByVideoHandler(db)).Methods("GET")

	// Like routes
	likeRoutes := router.PathPrefix("/api/v1/likes").Subrouter()
	likeRoutes.HandleFunc("/", like.CreateLikeHandler(db)).Methods("POST")
	likeRoutes.HandleFunc("/{id}", like.GetLikeHandler(db)).Methods("GET")
	likeRoutes.HandleFunc("/{id}", like.DeleteLikeHandler(db)).Methods("DELETE")

	// // Subscription routes
	subscriptionRoutes := router.PathPrefix("/api/v1/subscriptions").Subrouter()
	subscriptionRoutes.HandleFunc("/{subscriberID}/{channelID}", subscription.SubscribeHandler(db)).Methods("POST")
	subscriptionRoutes.HandleFunc("/{subscriptionID}", subscription.UnsubscribeHandler(db)).Methods("DELETE")

	// // Tweet routes
	tweetRoutes := router.PathPrefix("/api/v1/tweets").Subrouter()
	tweetRoutes.HandleFunc("/", tweet.CreateTweetHandler(db)).Methods("POST")
	tweetRoutes.HandleFunc("/{id}", tweet.GetTweetHandler(db)).Methods("GET")
	tweetRoutes.HandleFunc("/{id}", tweet.DeleteTweetHandler(db)).Methods("DELETE")
	tweetRoutes.HandleFunc("/{id}", tweet.EditTweetHandler(db)).Methods("PUT")

	return router
}
