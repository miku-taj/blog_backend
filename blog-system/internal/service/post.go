package service

import (
	"fmt"
	"gorm.io/gorm"
	"project/internal/model"
	"time"
)

func (s *Service) ListPosts(page int) ([]*model.Post, error) {
	return s.Repository.GetPosts(page)
}

//func (s *Service) ListPopularPosts(page int) ([]*model.Post, error) {
//	return s.Repository.GetPopularPosts(page)
//}

func (s *Service) FindPostByID(postID, userID int) (*model.Post, error) {
	post, err := s.Repository.GetPostByID(postID)
	if err == gorm.ErrRecordNotFound {
		return nil, fmt.Errorf("Post with id %d not found\n", postID)
	}
	if err != nil {
		return nil, err
	}
	if post.IsDraft == true && userID != post.UserID {
		return nil, fmt.Errorf("User is unauthorized to access the post with id %d", postID)
	}
	return post, nil

}

func (s *Service) FindPostsByTitle(title string) ([]*model.Post, error) {
	return s.Repository.SearchPostsByTitle(title)
}

func (s *Service) FindPostsByAuthorID(authorID, page int) ([]*model.Post, error) {
	_, err := s.Repository.GetUserByID(authorID)
	if err == gorm.ErrRecordNotFound {
		return nil, fmt.Errorf("User with id %d not found\n", authorID)
	}
	posts, err := s.Repository.GetPostsByAuthor(authorID, page)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (s *Service) ListDrafts(authorID, userID, page int) ([]*model.Post, error) {
	_, err := s.Repository.GetUserByID(authorID)
	if err == gorm.ErrRecordNotFound {
		return nil, fmt.Errorf("User with id %d not found\n", authorID)
	}
	if authorID != userID {
		return nil, fmt.Errorf("User is unauthorized to view drafts of author with id %d\n", authorID)
	}
	drafts, err := s.Repository.GetDraftsByAuthor(authorID, page)
	if err != nil {
		return nil, err
	}
	return drafts, nil
}

func (s *Service) CreatePost(post *model.Post, authorID int) (*model.Post, error) {
	post.UserID = authorID
	err := post.ValidatePost()
	if err != nil {
		return nil, err
	}

	return s.Repository.AddPost(post)
}

func (s *Service) UpdatePost(post *model.Post, userID int) (*model.Post, error) {
	oldPost, err := s.Repository.GetPostByID(post.PostID)
	if err == gorm.ErrRecordNotFound {
		return nil, fmt.Errorf("Post with id %d not found", post.PostID)
	}
	if err != nil {
		return nil, err
	}

	if oldPost.UserID != userID {
		return nil, fmt.Errorf("User is unauthorized\n")
	}

	err = post.ValidatePost()
	if err != nil {
		return nil, err
	}

	currentTime := time.Now()
	post.UserID = userID
	post.CreatedAt = oldPost.CreatedAt
	post.UpdatedAt = &currentTime
	post.LikesNumber = oldPost.LikesNumber
	return s.Repository.EditPost(post)
}

func (s *Service) DeletePost(postID, userID int) error {
	post, err := s.FindPostByID(postID, userID)
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	if post == nil || post.PostID == 0 {
		return fmt.Errorf("Post with id %d not found %v", postID, err)
	}
	return s.Repository.DeletePost(postID)
}
