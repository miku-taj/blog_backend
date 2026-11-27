package service

import (
	"fmt"
	"gorm.io/gorm"
	"project/internal/model"
	"time"
)

func (s *Service) ListPostComments(postID int, page int) ([]*model.Comment, error) {
	_, err := s.Repository.GetPostByID(postID)
	if err == gorm.ErrRecordNotFound {
		return nil, fmt.Errorf("Post with id %d not found %v\n", postID, err)
	} else if err != nil {
		return nil, err
	}
	return s.Repository.GetPostComments(postID, page)
}

func (s *Service) ListRecentPostComments(postID, page int) ([]*model.Comment, error) {
	_, err := s.Repository.GetPostByID(postID)
	if err == gorm.ErrRecordNotFound {
		return nil, fmt.Errorf("Post with id %d not found %v\n", postID, err)
	} else if err != nil {
		return nil, err
	}
	return s.Repository.GetRecentPostComments(postID, page)
}

func (s *Service) ListPopularPostComments(postID, page int) ([]*model.Comment, error) {
	_, err := s.Repository.GetPostByID(postID)
	if err == gorm.ErrRecordNotFound {
		return nil, fmt.Errorf("Post with id %d not found %v\n", postID, err)
	} else if err != nil {
		return nil, err
	}
	return s.Repository.GetPopularPostComments(postID, page)
}

func (s *Service) CreateComment(comment *model.Comment, userID, postID int) (*model.Comment, error) {
	post, err := s.Repository.GetPostByID(postID)
	if err == gorm.ErrRecordNotFound || post.PostID == 0 {
		return nil, fmt.Errorf("Post with id %d not found %v\n", comment.PostID, err)
	} else if err != nil {
		return nil, err
	}
	comment.UserID = userID
	comment.PostID = postID
	currentTime := time.Now()
	comment.CreatedAt = &currentTime
	return s.Repository.AddComment(comment)
}

func (s *Service) UpdateComment(comment *model.Comment, userID int) (*model.Comment, error) {
	oldComment, err := s.Repository.GetCommentByID(comment.CommentID)
	if err != nil || oldComment.CommentID == 0 {
		return nil, err
	}
	if oldComment.UserID != userID {
		return nil, fmt.Errorf("User is unauthorized\n")
	}
	comment.UserID = userID
	currentTime := time.Now()
	comment.CreatedAt = oldComment.CreatedAt
	comment.UpdatedAt = &currentTime
	return s.Repository.EditComment(comment)
}

func (s *Service) DeleteComment(commentID int, userID int, postID int) error {
	comment, err := s.Repository.GetCommentByID(commentID)
	if err == gorm.ErrRecordNotFound {
		return fmt.Errorf("Comment with id %d not found %v\n", commentID, err)
	}
	if err != nil {
		return err
	}
	post, err := s.Repository.GetPostByID(postID)
	if err == gorm.ErrRecordNotFound {
		return fmt.Errorf("Post with id %d not found %v\n", postID, err)
	}
	if err != nil {
		return err
	}
	if comment.UserID != userID && post.UserID != userID {
		return fmt.Errorf("User is unauthorized\n")
	}
	return s.Repository.DeleteComment(commentID)
}

func (s *Service) FindParentComment(commentID int) (*model.Comment, error) {
	_, err := s.Repository.GetCommentByID(commentID)
	if err == gorm.ErrRecordNotFound {
		return nil, fmt.Errorf("Comment with id %d not found %v\n", commentID, err)
	}
	comment, err := s.Repository.GetParentComment(commentID)
	if err != nil {
		return nil, err
	}
	return comment, nil
}
func (s *Service) FindChildrenComments(commentID int) ([]*model.Comment, error) {
	_, err := s.Repository.GetCommentByID(commentID)
	if err == gorm.ErrRecordNotFound {
		return nil, fmt.Errorf("Comment with id %d not found %v\n", commentID, err)
	}
	return s.Repository.GetChildrenComments(commentID)
}
