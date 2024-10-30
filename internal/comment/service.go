package comment

import "go.mongodb.org/mongo-driver/bson/primitive"

type CommentService struct {
	repository *CommentRepository
}

func NewCommentService(repo *CommentRepository) *CommentService {
	return &CommentService{repository: repo}
}

func (s *CommentService) Create(comment *Comment) error {
	return s.repository.Create(comment)
}

func (s *CommentService) Update(comment *Comment) error {
	return s.repository.Update(comment)
}

func (s *CommentService) Delete(id primitive.ObjectID) error {
	return s.repository.Delete(id)
}

func (s *CommentService) GetByID(id primitive.ObjectID) (*Comment, error) {
	return s.repository.GetByID(id)
}

func (s *CommentService) GetAllByVideo(videoID primitive.ObjectID) ([]Comment, error) {
	return s.repository.GetAllByVideo(videoID)
}
