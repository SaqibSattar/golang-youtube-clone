// internal/playlist/service.go

package playlist

import "go.mongodb.org/mongo-driver/bson/primitive"

type PlaylistService struct {
	repository *PlaylistRepository
}

func NewPlaylistService(repo *PlaylistRepository) *PlaylistService {
	return &PlaylistService{repository: repo}
}

func (s *PlaylistService) Create(playlist *Playlist) error {
	return s.repository.Create(playlist)
}

func (s *PlaylistService) Update(playlist *Playlist) error {
	return s.repository.Update(playlist)
}

func (s *PlaylistService) Delete(id primitive.ObjectID) error {
	return s.repository.Delete(id)
}

func (s *PlaylistService) GetByID(id primitive.ObjectID) (*Playlist, error) {
	return s.repository.GetByID(id)
}

func (s *PlaylistService) GetAll() ([]Playlist, error) {
	return s.repository.GetAll()
}
