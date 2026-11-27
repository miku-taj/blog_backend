package repository

import (
	"database/sql"
	"fmt"
	"gorm.io/gorm"
	"project/internal/model"
)

var (
	commentsOnPage = 10
)

func (r *Repository) GetPostComments(postID int, page int) ([]*model.Comment, error) {
	var comments []*model.Comment
	if err := r.db.Offset((page-1)*commentsOnPage).Limit(commentsOnPage).Find(&comments, model.Comment{PostID: postID}).Error; err != nil {
		return nil, err
	} else {
		return comments, nil
	}
}

func (r *Repository) GetRecentPostComments(postID, page int) ([]*model.Comment, error) {
	var comments []*model.Comment
	//if err := r.db.Raw("SELECT * FROM comments WHERE parent_id IS NULL ORDER BY created_at").Scan(&comments).Error; err != nil {
	//	return nil, err
	//} else {
	//	return comments, nil
	//}
	if err := r.db.Offset((page - 1) * commentsOnPage).Limit(commentsOnPage).Order("created_at desc").Find(&comments).Error; err != nil {
		return nil, err
	} else {
		return comments, nil
	}
}

func (r *Repository) GetPopularPostComments(postID, page int) ([]*model.Comment, error) {
	var comments []*model.Comment
	query := r.db.Order("likes_number desc")
	if query.Error != nil {
		return nil, query.Error
	}
	if err := query.Offset((page-1)*commentsOnPage).Limit(commentsOnPage).Find(&comments, model.Comment{PostID: postID}).Error; err != nil {
		return nil, err
	} else {
		return comments, nil
	}
}

func (r *Repository) GetCommentByID(commentID int) (*model.Comment, error) {
	var comment *model.Comment
	if err := r.db.First(&comment, commentID).Error; err != nil {
		return nil, err
	}
	return comment, nil
}

func (r *Repository) GetParentComment(childCommentID int) (*model.Comment, error) {
	var comment *model.Comment
	if comment.ParentID.Valid == false {
		return nil, fmt.Errorf("Comment woth id %d does not have parent comment\n", childCommentID)
	}
	var parentComment *model.Comment
	if err := r.db.First(&parentComment, int(comment.ParentID.Int64)).Error; err != nil {
		return nil, err
	}
	return parentComment, nil
}

func (r *Repository) GetChildrenComments(parentCommentID int) ([]*model.Comment, error) {
	var comments []*model.Comment
	if err := r.db.Find(&comments, model.Comment{ParentID: sql.NullInt64{Int64: int64(parentCommentID), Valid: true}}).Error; err != nil {
		return nil, err
	}
	if len(comments) == 0 {
		return nil, fmt.Errorf("comment with id %d has no cildren comments", parentCommentID)
	}
	return comments, nil
}

func (r *Repository) AddComment(comment *model.Comment) (*model.Comment, error) {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := r.db.Create(comment).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return comment, nil
}

func (r *Repository) EditComment(comment *model.Comment) (*model.Comment, error) {
	result := r.db.Save(comment)
	if result.Error != nil {
		return nil, result.Error
	}
	result = r.db.Find(&comment, comment.CommentID)
	if result.Error != nil {
		return nil, result.Error
	}
	return comment, nil
}

func (r *Repository) DeleteComment(commentID int) error {

	children, err := r.GetChildrenComments(commentID)
	if len(children) != 0 {
		for child := range children {
			if err := r.DeleteComment(children[child].CommentID); err != nil {
				return err
			}
		}
	}
	if err != nil {
		return err
	}
	err = r.db.Transaction(func(tx *gorm.DB) error {
		fmt.Printf("Deleting coment with id %d", commentID)
		if err := r.db.Delete(&model.Comment{CommentID: commentID}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}
