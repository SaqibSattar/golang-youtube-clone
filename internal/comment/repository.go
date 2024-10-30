package comment

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CommentRepository struct {
	collection *mongo.Collection
}

func NewCommentRepository(db *mongo.Database) *CommentRepository {
	return &CommentRepository{
		collection: db.Collection("comments"),
	}
}

func (r *CommentRepository) Create(comment *Comment) error {
	comment.CreatedAt = time.Now()
	comment.UpdatedAt = time.Now()
	_, err := r.collection.InsertOne(context.TODO(), comment)
	return err
}

func (r *CommentRepository) Update(comment *Comment) error {
	comment.UpdatedAt = time.Now()
	_, err := r.collection.UpdateOne(context.TODO(),
		map[string]interface{}{"_id": comment.ID},
		map[string]interface{}{"$set": comment})
	return err
}

func (r *CommentRepository) Delete(id primitive.ObjectID) error {
	_, err := r.collection.DeleteOne(context.TODO(), map[string]interface{}{"_id": id})
	return err
}

func (r *CommentRepository) GetByID(id primitive.ObjectID) (*Comment, error) {
	var comment Comment
	err := r.collection.FindOne(context.TODO(), map[string]interface{}{"_id": id}).Decode(&comment)
	return &comment, err
}

func (r *CommentRepository) GetAllByVideo(videoID primitive.ObjectID) ([]Comment, error) {
	var comments []Comment
	cursor, err := r.collection.Find(context.TODO(), map[string]interface{}{"video": videoID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var comment Comment
		if err = cursor.Decode(&comment); err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}

	return comments, cursor.Err()
}
