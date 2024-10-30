package like

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LikeService struct {
	repository *LikeRepository
}

func NewLikeService(repo *LikeRepository) *LikeService {
	return &LikeService{repository: repo}
}

func (s *LikeService) Create(like *Like) error {
	return s.repository.Create(like)
}

func (s *LikeService) Delete(id primitive.ObjectID) error {
	return s.repository.Delete(id)
}

func (s *LikeService) GetByID(id primitive.ObjectID) (*Like, error) {
	return s.repository.GetByID(id)
}
