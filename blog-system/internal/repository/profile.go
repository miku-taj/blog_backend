package repository

import "project/internal/model"

func (r *Repository) GetUserProfile(userID int) (*model.Profile, error) {
	var profile *model.Profile
	if err := r.db.First(&profile, model.Profile{UserID: userID}).Error; err != nil {
		return nil, err
	}
	return profile, nil
}

func (r *Repository) AddProfile(profile *model.Profile) (*model.Profile, error) {
	if err := r.db.Create(profile).Error; err != nil {
		return nil, err
	}
	if err := r.db.Find(&profile, profile.ProfileID).Error; err != nil {
		return nil, err
	}
	return profile, nil
}

func (r *Repository) EditProfile(profile *model.Profile) (*model.Profile, error) {
	if err := r.db.Save(profile).Error; err != nil {
		return nil, err
	}
	if err := r.db.Find(&profile, profile.ProfileID).Error; err != nil {
		return nil, err
	}
	return profile, nil
}

func (r *Repository) DeleteProfile(userID int) error {
	var profile model.Profile
	profile.UserID = userID
	if err := r.db.Delete(&profile).Error; err != nil {
		return err
	}
	return nil
}
