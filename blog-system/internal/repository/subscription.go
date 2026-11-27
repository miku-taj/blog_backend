package repository

import (
	"project/internal/model"
)

func (r *Repository) FindSubscription(authorID, userID int) (bool, error) {
	var subscription model.Subscription
	if err := r.db.First(&subscription, &model.Subscription{AuthorID: authorID, UserID: userID}).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (r *Repository) Subscribe(subscription *model.Subscription) error {
	if err := r.db.Create(subscription).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repository) Unsubscribe(subscription *model.Subscription) error {
	if err := r.db.Delete(subscription).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetSubscribers(authorID int) ([]*model.User, error) {
	var subscription []model.Subscription
	var userIDs []int
	var users []*model.User
	if err := r.db.Find(&subscription, model.Subscription{AuthorID: authorID}).Error; err != nil {
		return nil, err
	}
	if len(subscription) == 0 {
		return nil, nil
	}
	for s := range subscription {
		userIDs = append(userIDs, subscription[s].UserID)
	}
	if err := r.db.Select("user_id, username").Find(&users, userIDs).Error; err != nil {
		return nil, err
	}
	return users, nil
}
