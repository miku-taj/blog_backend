package service

import (
	"fmt"
	"gorm.io/gorm"
	"project/internal/model"
)

func (s *Service) FindProfile(userID int) (*model.Profile, error) {
	_, err := s.FindUserByID(userID)
	if err != nil {
		return nil, err
	}
	p, err := s.Repository.GetUserProfile(userID)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	} else if err == gorm.ErrRecordNotFound {
		return nil, fmt.Errorf("User with id %d does not have a profile %v\n", err)
	}
	return p, nil
}

func (s *Service) CreateProfile(profile *model.Profile, userID int) (*model.Profile, error) {
	_, err := s.FindProfile(userID)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	if err == nil {
		return nil, fmt.Errorf("User with id %d already as a profile\n", userID)
	}
	err = profile.ValidateProfile()
	if err != nil {
		return nil, err
	}
	profile.UserID = userID
	return s.Repository.AddProfile(profile)
}

func (s *Service) UpdateProfile(profile *model.Profile, userID int) (*model.Profile, error) {
	profile.UserID = userID
	err := profile.ValidateProfile()
	if err != nil {
		return nil, err
	}
	return s.Repository.EditProfile(profile)
}

func (s *Service) DeleteProfile(userID int) error {
	_, err := s.FindProfile(userID)
	if err != nil {
		return err
	}
	return s.Repository.DeleteProfile(userID)
}
