package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func (h *BlogHandler) ListPostLikes(c *gin.Context) {
	postID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("Failed to ListPostLikes - strconv.Atoi error %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	users, err := h.service.ListPostLikesUsers(postID)
	if err != nil {
		log.Printf("Failed to ListPostLikes - service.ListPostLikesUsers error %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if users == nil {
		c.JSON(http.StatusOK, gin.H{"message": "No one liked the post yet"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": users})
}

func (h *BlogHandler) LikePost(c *gin.Context) {
	postID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("Failed to LikePost - strconv.Atoi error %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	value, exists := c.Get("user_id")
	if !exists || value == nil {
		log.Printf("Failed to UnlikePost - user is not authenticated %v\n", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": fmt.Errorf("You must log in")})
		return
	}
	userID := value.(int)

	_, err = h.service.LikePost(userID, postID)
	if err != nil {
		log.Printf("Failed to LikePost - service.LikePost error %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Post liked"})
}

func (h *BlogHandler) UnlikePost(c *gin.Context) {
	postID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("Failed to UnlikePost - strconv.Atoi error %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	value, exists := c.Get("user_id")
	if !exists || value == nil {
		log.Printf("Failed to UnlikePost - user is not authenticated %v\n", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": fmt.Errorf("You must log in")})
		return
	}
	userID := value.(int)

	err = h.service.UnlikePost(userID, postID)
	if err != nil {
		log.Printf("Failed to UnlikePost - service.UnlikePost error %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Post unliked"})
}

func (h *BlogHandler) ListCommentLikes(c *gin.Context) {
	commentID, err := strconv.Atoi(c.Param("comment-id"))
	if err != nil {
		log.Printf("Failed to ListCommentLikes - strconv.Atoi error %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	users, err := h.service.ListCommentLikesUsers(commentID)
	if users == nil {
		c.JSON(http.StatusOK, gin.H{"message": "No one liked the comment yet"})
		return
	}
	if err != nil {
		log.Printf("Failed to ListCommentLikes - service.ListCommentLikesUsers error %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": users})
}

func (h *BlogHandler) LikeComment(c *gin.Context) {
	postID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("Failed to LikeComment - strconv.Atoi error %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	value, exists := c.Get("user_id")
	if !exists || value == nil {
		log.Printf("Failed to LikeComment - user is not authenticated %v\n", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": fmt.Errorf("You must log in")})
		return
	}
	userID := value.(int)

	_, err = h.service.LikeComment(userID, postID)
	if err != nil {
		log.Printf("Failed to LikeComment - service.LikeComment error %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Comment liked"})
}

func (h *BlogHandler) UnlikeComment(c *gin.Context) {
	commentID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("Failed to UnlikeComment - strconv.Atoi error %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	value, exists := c.Get("user_id")
	if !exists || value == nil {
		log.Printf("Failed to UnlikeComment - user is not authenticated %v\n", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": fmt.Errorf("You must log in")})
		return
	}
	userID := value.(int)

	err = h.service.UnlikeComment(userID, commentID)
	if err != nil {
		log.Printf("Failed to UnlikeComment - service.UnlikeComment error %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Comment unliked"})
}
