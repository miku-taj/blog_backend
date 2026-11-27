package repository

import (
	"fmt"
	"gorm.io/gorm"
	"project/internal/model"
)

var UserOnPage = 20

func (r *Repository) GetUsers(page int) ([]*model.User, error) {
	var users []*model.User
	result := r.db.Offset((page-1)*UserOnPage).Limit(UserOnPage).Select("user_id", "username", "created_at", "updated_at").Find(&users)
	if result.Error != nil {
		return nil, fmt.Errorf("Failed to Get Users - db.Find error %v\n", result.Error)
	}
	return users, nil
}

func (r *Repository) GetUserByID(userID int) (*model.User, error) {
	var user *model.User
	result := r.db.First(&user, userID)
	if result.Error != nil {
		return nil, fmt.Errorf("Failed to Get User By ID - db.Find error %v\n", result.Error)
	}
	return user, nil
}

func (r *Repository) GetUsersByUsernameLike(username string) ([]*model.User, error) {
	var users []*model.User
	result := r.db.Select("user_id", "username", "created_at", "updated_at").Where("username LIKE ?", "%"+username+"%").Find(&users)
	if result.Error != nil {
		return nil, fmt.Errorf("Failed to Get Users By Username - db.Find error %v\n", result.Error)
	}
	return users, nil
}

func (r *Repository) GetUserByExactUsername(username string) (*model.User, error) {
	var user *model.User
	if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, fmt.Errorf("Failed to Get User By Exact Username - db.Find error %v\n", err)
	}
	return user, nil
}

func (r *Repository) AddUser(user *model.User) (*model.User, error) {
	result := r.db.Create(user)
	if result.Error != nil {
		return nil, fmt.Errorf("Failed to Add User - db.Find error %v\n", result.Error)
	}
	return user, nil
}

func (r *Repository) EditUser(user *model.User) (*model.User, error) {
	result := r.db.Save(user)
	if result.Error != nil {
		return nil, fmt.Errorf("Failed to Edit User - db.Save error %v\n", result.Error)
	}
	var newUser *model.User
	result = r.db.Find(&newUser, user.UserID)
	if result.Error != nil {
		return nil, fmt.Errorf("Failed to Edit User - db.Find error %v\n", result.Error)
	}
	return user, nil
}

func (r *Repository) DeleteUser(id int) error {

	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := r.db.Delete(&model.Post{UserID: id}).Error; err != nil {
			return err
		}
		if err := r.db.Delete(&model.Subscription{AuthorID: id}).Error; err != nil {
			return err
		}
		return nil
	})

	return err
}
