package repository

import (
	"gorm.io/gorm"
	"project/internal/model"
)

func (r *Repository) FindPostLike(like *model.PostLike) (*model.PostLike, error) {
	if err := r.db.First(&like).Error; err != nil {
		return nil, err
	}
	return like, nil
}

func (r *Repository) FindCommentLike(like *model.CommentLike) (*model.CommentLike, error) {
	if err := r.db.First(&like).Error; err != nil {
		return nil, err
	}
	return like, nil
}

func (r *Repository) LikePost(like *model.PostLike) (*model.PostLike, error) {
	var post *model.Post
	if err := r.db.First(&post, like.PostID).Error; err != nil {
		return nil, err
	}
	post.LikesNumber = post.LikesNumber + 1
	if err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := r.db.Create(&like).Error; err != nil {
			return err
		}
		if err := r.db.Save(&post).Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}
	if err := r.db.Find(&like).Error; err != nil {
		return nil, err
	}
	return like, nil
}

func (r *Repository) LikeComment(like *model.CommentLike) (*model.CommentLike, error) {
	var comment *model.Comment
	if err := r.db.First(&comment, like.CommentID).Error; err != nil {
		return nil, err
	}
	comment.LikesNumber = comment.LikesNumber + 1
	if err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := r.db.Create(&like).Error; err != nil {
			return err
		}
		if err := r.db.Save(&comment).Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}
	if err := r.db.Find(&like).Error; err != nil {
		return nil, err
	}
	return like, nil
}

func (r *Repository) UnlikePost(like *model.PostLike) error {

	if err := r.db.Delete(like).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repository) UnlikeComment(like *model.CommentLike) error {
	if err := r.db.Delete(like).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetPostLikesUsers(postID int) ([]model.User, error) {
	var users []model.User
	var likes []model.PostLike
	var userIDs []int
	if err := r.db.Select("user_id").Find(&likes, model.PostLike{PostID: postID}).Error; err != nil {
		return nil, err
	}
	for l := range likes {
		userIDs = append(userIDs, likes[l].UserID)
	}
	if len(userIDs) == 0 {
		return nil, nil
	}
	if err := r.db.Select("user_id, username, created_at, updated_at").Find(&users, userIDs).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *Repository) GetCommentLikesUsers(commentID int) ([]model.User, error) {
	var users []model.User
	var likes []model.CommentLike
	var userIDs []int
	if err := r.db.Select("user_id").Find(&likes, model.CommentLike{CommentID: commentID}).Error; err != nil {
		return nil, err
	}
	for l := range likes {
		userIDs = append(userIDs, likes[l].UserID)
	}
	if len(userIDs) == 0 {
		return nil, nil
	}
	if err := r.db.Select("user_id, username, created_at, updated_at").Find(&users, userIDs).Error; err != nil {
		return nil, err
	}
	return users, nil
}
