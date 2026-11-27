package service

import (
	"fmt"
	"gorm.io/gorm"
	"project/internal/model"
)

func (s *Service) Subscribe(authorID, userID int) error {
	if authorID == userID {
		return gorm.ErrInvalidData
	}
	isSubscribed, err := s.Repository.FindSubscription(authorID, userID)
	if err != nil {
		return err
	}
	if isSubscribed {
		return fmt.Errorf("You are already subscribed to this author")
	}
	var subscription = model.Subscription{AuthorID: authorID, UserID: userID}
	return s.Repository.Subscribe(&subscription)
}

func (s *Service) Unsubscribe(authorID, userID int) error {
	_, err := s.Repository.FindSubscription(authorID, userID)
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	if err == gorm.ErrRecordNotFound {
		return nil
	}
	var subscription = model.Subscription{AuthorID: authorID, UserID: userID}
	return s.Repository.Unsubscribe(&subscription)
}

func (s *Service) ListSubscribers(authorID int) ([]*model.User, error) {
	user, err := s.Repository.GetUserByID(authorID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, fmt.Errorf("User with id %d is not found", authorID)
	}
	return s.Repository.GetSubscribers(authorID)
}
