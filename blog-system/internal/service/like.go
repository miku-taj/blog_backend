package service

import (
	"fmt"
	"gorm.io/gorm"
	"project/internal/model"
)

func (s *Service) LikePost(userID, postID int) (*model.PostLike, error) {
	_, err := s.Repository.GetPostByID(postID)
	if err == gorm.ErrRecordNotFound {
		return nil, fmt.Errorf("post with id %d not found %v", postID, err)
	} else if err != nil {
		return nil, err
	}
	_, err = s.Repository.GetUserByID(userID)
	if err == gorm.ErrRecordNotFound {
		return nil, fmt.Errorf("User with id %d not found %v\n", userID, err)
	} else if err != nil {
		return nil, err
	}
	var like model.PostLike
	like.UserID = userID
	like.PostID = postID
	l, err := s.Repository.FindPostLike(&like)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	if err == nil && l != nil {
		return nil, fmt.Errorf("you already liked the post")
	}
	return s.Repository.LikePost(&like)
}

func (s *Service) LikeComment(userID, commentID int) (*model.CommentLike, error) {
	_, err := s.Repository.GetCommentByID(commentID)
	if err == gorm.ErrRecordNotFound {
		return nil, fmt.Errorf("Comment with id %d not found %v\n", commentID, err)
	} else if err != nil {
		return nil, err
	}
	_, err = s.Repository.GetUserByID(userID)
	if err == gorm.ErrRecordNotFound {
		return nil, fmt.Errorf("User with id %d not found %v\n", userID, err)
	} else if err != nil {
		return nil, err
	}
	var like model.CommentLike
	like.UserID = userID
	like.CommentID = commentID
	l, err := s.Repository.FindCommentLike(&like)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	if err == nil && l != nil {
		return nil, fmt.Errorf("you already liked the comment")
	}
	return s.Repository.LikeComment(&like)
}

func (s *Service) UnlikePost(userID, postID int) error {
	var like model.PostLike
	like.UserID = userID
	like.PostID = postID
	if _, err := s.Repository.FindPostLike(&like); err == gorm.ErrRecordNotFound {
		return nil
	} else if err != nil {
		return err
	}
	return s.Repository.UnlikePost(&like)
}

func (s *Service) UnlikeComment(userID, commentID int) error {
	var like model.CommentLike
	like.UserID = userID
	like.CommentID = commentID
	if _, err := s.Repository.FindCommentLike(&like); err == gorm.ErrRecordNotFound {
		return nil
	} else if err != nil {
		return err
	}
	return s.Repository.UnlikeComment(&like)
}

func (s *Service) ListPostLikesUsers(postID int) ([]model.User, error) {
	_, err := s.Repository.GetPostByID(postID)
	if err == gorm.ErrRecordNotFound {
		return nil, fmt.Errorf("Post with id %d not found %v\n", postID, err)
	} else if err != nil {
		return nil, err
	}
	return s.Repository.GetPostLikesUsers(postID)
}

func (s *Service) ListCommentLikesUsers(commentID int) ([]model.User, error) {
	_, err := s.Repository.GetCommentByID(commentID)
	if err == gorm.ErrRecordNotFound {
		return nil, fmt.Errorf("Comment with id %d not found %v\n", commentID, err)
	} else if err != nil {
		return nil, err
	}
	return s.Repository.GetCommentLikesUsers(commentID)
}
