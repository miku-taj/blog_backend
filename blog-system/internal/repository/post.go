package repository

import (
	"fmt"
	"gorm.io/gorm"
	"project/internal/model"
	"time"
)

var postsOnPage = 10

func (r *Repository) GetPosts(page int) ([]*model.Post, error) {
	var posts []*model.Post
	result := r.db.Offset((page-1)*postsOnPage).Limit(postsOnPage).Find(&posts, model.Post{IsDraft: false})
	if result.Error != nil {
		return nil, result.Error
	}
	return posts, nil
}

func (r *Repository) GetPostByID(postID int) (*model.Post, error) {
	var post model.Post
	if err := r.db.First(&post, postID).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

// SearchPostByTitle - implement full-text indexing here?
func (r *Repository) SearchPostsByTitle(title string) ([]*model.Post, error) {
	var posts []*model.Post
	if err := r.db.Where("title = ?", title).Find(&posts).Error; err != nil {
		return nil, fmt.Errorf("Failed to Get User By Exact Username - db.Find error %v\n", err)
	}
	if posts[0].PostID == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return posts, nil
	return nil, nil
}

func (r *Repository) GetPostsByAuthor(authorID, page int) ([]*model.Post, error) {
	var posts []*model.Post
	result := r.db.Offset((page-1)*postsOnPage).Limit(postsOnPage).Where("author_id=? AND is_draft = ?", authorID, false).Find(&posts)
	if result.Error != nil {
		return nil, result.Error
	}
	return posts, nil
}

func (r *Repository) GetDraftsByAuthor(authorID, page int) ([]*model.Post, error) {
	var posts []*model.Post
	result := r.db.Offset((page-1)*postsOnPage).Limit(postsOnPage).Where("author_id=? AND is_draft = ?", authorID, true).Find(&posts)
	if result.Error != nil {
		return nil, result.Error
	}
	return posts, nil
}

func (r *Repository) GetPopularPosts(page int) ([]*model.Post, error) {
	var posts []*model.Post
	result := r.db.Where("created_at >= ?", time.Now().Add(-time.Hour*24*7))
	if result.Error != nil {
		return nil, result.Error
	}
	result = result.Order("likes_number desc").Offset((page - 1) * 10).Limit(10).Find(&posts)
	if result.Error != nil {
		return nil, result.Error
	}
	return posts, nil
}

func (r *Repository) AddPost(post *model.Post) (*model.Post, error) {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := r.db.Create(post).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (r *Repository) EditPost(post *model.Post) (*model.Post, error) {
	result := r.db.Save(post)
	if result.Error != nil {
		return nil, result.Error
	}
	result = r.db.Find(&post, post.PostID)
	if result.Error != nil {
		return nil, result.Error
	}
	return post, nil
}

func (r *Repository) DeletePost(postID int) error {

	err := r.db.Transaction(func(tx *gorm.DB) error {

		if err := r.db.Delete(&model.Comment{PostID: postID}).Error; err != nil {
			return err
		}
		if err := r.db.Delete(&model.PostLike{PostID: postID}).Error; err != nil {
			return err
		}
		if err := r.db.Delete(&model.Post{PostID: postID}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}
