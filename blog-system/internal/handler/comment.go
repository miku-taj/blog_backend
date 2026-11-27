package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"project/internal/model"
	"strconv"
)

func (h *BlogHandler) GetComments(c *gin.Context) {
	postID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("Failed to GetComments - strconv.Atoi error %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	page := 1
	if c.Param("page-id") != "" {
		page, err = strconv.Atoi(c.Param("page-id"))
		if err != nil {
			log.Printf("Failed to GetComments - strconv.Atoi error %v\n", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}
	comments, err := h.service.ListPostComments(postID, page)
	if err != nil {
		log.Printf("Failed to GetComments - service.ListPostComments error %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": comments})
}

func (h *BlogHandler) GetRecentPostComments(c *gin.Context) {
	postID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("Failed to GetRecentPostComments - strconv.Atoi error %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	page := 1
	if c.Param("page-id") != "" {
		page, err = strconv.Atoi(c.Param("page-id"))
		if err != nil {
			log.Printf("Failed to GetRecentPostComments - strconv.Atoi error %v\n", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}
	comments, err := h.service.ListRecentPostComments(postID, page)
	if err != nil {
		log.Printf("Failed to GetRecentPostComments - service.ListRecentPostComments error %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": comments})
}

func (h *BlogHandler) GetPopularPostComments(c *gin.Context) {
	postID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("Failed to GetRecentPostComments - strconv.Atoi error %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	page := 1
	if c.Param("page-id") != "" {
		page, err = strconv.Atoi(c.Param("page-id"))
		if err != nil {
			log.Printf("Failed to GetRecentPostComments - strconv.Atoi error %v\n", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}
	comments, err := h.service.ListPopularPostComments(postID, page)
	if err != nil {
		log.Printf("Failed to GetRecentPostComments - service.ListRecentPostComments error %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": comments})
}

func (h *BlogHandler) AddComment(c *gin.Context) {
	postID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("Failed to AddComment - strconv.Atoi error %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	value, exists := c.Get("user_id")
	if !exists || value == nil {
		log.Printf("Failed to AddComment - user is not authenticated %v\n", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": fmt.Errorf("You must log in")})
		return
	}
	userID := value.(int)
	var comment *model.Comment
	if err := c.BindJSON(&comment); err != nil {
		log.Printf("Failed to AddPost - c.BindJSON error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	comment, err = h.service.CreateComment(comment, userID, postID)
	if err != nil {
		log.Printf("Failed to AddComment - service.CreateComment error %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": comment})
}

func (h *BlogHandler) UpdateComment(c *gin.Context) {
	postID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("Failed to AddComment - strconv.Atoi error %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	value, exists := c.Get("user_id")
	if !exists || value == nil {
		log.Printf("Failed to AddComment - user is not authenticated %v\n", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": fmt.Errorf("You must log in\n")})
		return
	}
	userID := value.(int)
	var comment *model.Comment
	if err := c.BindJSON(&comment); err != nil {
		log.Printf("Failed to AddPost - c.BindJSON error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	comment.PostID = postID
	comment, err = h.service.UpdateComment(comment, userID)
	if err != nil {
		log.Printf("Failed to UpdateComment - service.UpdateComment %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": comment})
}

func (h *BlogHandler) DeleteComment(c *gin.Context) {
	postID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("Failed to DeleteComment - strconv.Atoi error %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	value, exists := c.Get("user_id")
	if !exists || value == nil {
		log.Printf("Failed to AddComment - user is not authenticated %v\n")
		c.JSON(http.StatusUnauthorized, gin.H{"error": fmt.Errorf("You must log in")})
		return
	}
	userID := value.(int)
	commentID, err := strconv.Atoi(c.Param("comment-id"))
	if err != nil {
		log.Printf("Failed to DeleteComment - strconv.Atoi error %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = h.service.DeleteComment(commentID, userID, postID)
	if err != nil {
		log.Printf("Failed to UpdateComment - service.UpdateComment %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Comment deleted"})
}

func (h *BlogHandler) GetParentComment(c *gin.Context) {
	commentID, err := strconv.Atoi(c.Param("comment-id"))
	if err != nil {
		log.Printf("Failed to GetParentComment - strconv.Atoi error %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	comment, err := h.service.FindParentComment(commentID)
	if err != nil {
		log.Printf("Failed to GetParentComment - service.FindParentComment %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": comment})
}

func (h *BlogHandler) GetChildrenComments(c *gin.Context) {
	commentID, err := strconv.Atoi(c.Param("comment-id"))
	if err != nil {
		log.Printf("Failed to GetChildrenComments - strconv.Atoi error %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	comment, err := h.service.FindChildrenComments(commentID)
	if err != nil {
		log.Printf("Failed to GetChildrenComments - service.FindChildrenComments %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": comment})
}
