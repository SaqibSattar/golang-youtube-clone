package video

import "go.mongodb.org/mongo-driver/bson/primitive"

// VideoService provides methods to work with videos
type VideoService struct {
	repository *VideoRepository
}

// NewVideoService creates a new VideoService
func NewVideoService(repo *VideoRepository) *VideoService {
	return &VideoService{repository: repo}
}

// CreateVideo creates a new video
func (svc *VideoService) CreateVideo(video *Video) error {
	return svc.repository.Create(video)
}

// GetVideoByID retrieves a video by its ID
func (svc *VideoService) GetVideoByID(id primitive.ObjectID) (*Video, error) {
	return svc.repository.FindByID(id)
}

// UpdateVideo updates an existing video
func (svc *VideoService) UpdateVideo(video *Video) error {
	return svc.repository.Update(video)
}

// DeleteVideo deletes a video by its ID
func (svc *VideoService) DeleteVideo(id primitive.ObjectID) error {
	return svc.repository.Delete(id)
}

// GetAllVideos retrieves all videos
func (svc *VideoService) GetAllVideos() ([]Video, error) {
	return svc.repository.GetAll()
}
