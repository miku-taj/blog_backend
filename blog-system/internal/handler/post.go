package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"project/internal/model"
	"strconv"
)

func (h *BlogHandler) GetPosts(c *gin.Context) {
	page := 1
	if c.Param("page-id") != "" {
		page2, err := strconv.Atoi(c.Param("page-id"))
		page = page2
		if err != nil {
			log.Printf("Failed to GetPosts - strconv.Atoi error %v\n", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	posts, err := h.service.ListPosts(page)
	if err != nil {
		log.Printf("Failed to Get Posts - service.ListPosts error %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": posts})
}

func (h *BlogHandler) GetPostByID(c *gin.Context) {
	// получаем пользователей
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("Failed to Get Post By ID - strconv.Atoi error %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	value, exists := c.Get("user_id")
	if !exists {
		log.Printf("Failed to AddPost - user unauthorized error")
		c.JSON(http.StatusUnauthorized, gin.H{"error": fmt.Errorf("Cannot post without login")})
		return
	}
	user_id := value.(int)
	post, err := h.service.FindPostByID(id, user_id)
	if err != nil {
		log.Printf("Failed to Get Post by ID - service.FindPostByID error %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": post})
}

func (h *BlogHandler) AddPost(c *gin.Context) {
	var post model.Post

	if err := c.BindJSON(&post); err != nil {
		log.Printf("Failed to AddPost - c.BindJSON error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	value, exists := c.Get("user_id")
	if !exists || value == nil {
		log.Printf("Failed to AddPost - user unauthorized error")
		c.JSON(http.StatusUnauthorized, gin.H{"error": fmt.Errorf("Cannot post without login")})
		return
	}

	UserID := value.(int)

	createdPost, err := h.service.CreatePost(&post, UserID)
	if err != nil {
		log.Printf("Failed to AddPost - h.service.CreatePost error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": createdPost})
}

func (h *BlogHandler) EditPost(c *gin.Context) {
	var post model.Post
	if err := c.BindJSON(&post); err != nil {
		log.Printf("Failed to EditPost - c.BindJSON error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	value, exists := c.Get("user_id")
	if !exists {
		log.Printf("Failed to AddPost - user unauthorized error")
		c.JSON(http.StatusUnauthorized, gin.H{"error": fmt.Errorf("Cannot post without login")})
		return
	}

	UserID := value.(int)

	updatedPost, err := h.service.UpdatePost(&post, UserID)
	if err != nil {
		log.Printf("Failed to EditPost - h.service.UpdatePost error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": updatedPost})
}

func (h *BlogHandler) DeletePost(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("Failed to Delete Post By ID - strconv.Atoi error %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	value, exists := c.Get("user_id")
	if !exists {
		log.Printf("Failed to AddPost - user unauthorized error")
		c.JSON(http.StatusUnauthorized, gin.H{"error": fmt.Errorf("Cannot post without login")})
		return
	}

	oldPost, err := h.service.FindPostByID(id, value.(int))
	if err != nil {
		log.Printf("Failed to DeletePost - h.service.FindPostByID error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if oldPost.UserID != value.(int) {
		log.Printf("Failed to EditPost - User is not authorized error")
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	err = h.service.DeletePost(id, value.(int))
	if err != nil {
		log.Printf("Failed to Delete Post By ID - service.DeletePost error %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post deleted"})

}
