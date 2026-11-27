package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
)

func (h *BlogHandler) Subscribe(c *gin.Context) {
	authorID, err := strconv.Atoi(c.Param("user-id"))
	if err != nil {
		log.Printf("Failed to Subscribe - strconv.Atoi error %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	value, exists := c.Get("user_id")
	if !exists || value == nil {
		log.Printf("Failed to Subscribe - user is not authenticated %v\n", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": fmt.Errorf("You must log in")})
		return
	}
	userID := value.(int)
	err = h.service.Subscribe(authorID, userID)
	if err == gorm.ErrInvalidData {
		log.Printf("Failed to Subscribe - service.Subscribe error %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Errorf("User cannot subscrie to himself")})
		return
	}
	if err != nil {
		log.Printf("Failed to Subscribe - service.Subscribe error %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Subscribed to user"})
}

func (h *BlogHandler) Unsubscribe(c *gin.Context) {
	authorID, err := strconv.Atoi(c.Param("user-id"))
	if err != nil {
		log.Printf("Failed to Unsubscribe - strconv.Atoi error %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	value, exists := c.Get("user_id")
	if !exists || value == nil {
		log.Printf("Failed to Unsubscribe - user is not authenticated %v\n", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": fmt.Errorf("You must log in")})
		return
	}
	userID := value.(int)
	err = h.service.Unsubscribe(authorID, userID)
	if err != nil {
		log.Printf("Failed to Unsubscribe - service.Unsubscribe error %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Unsubscribed from user"})
}

func (h *BlogHandler) ViewSubscribers(c *gin.Context) {
	authorID, err := strconv.Atoi(c.Param("user-id"))
	if err != nil {
		log.Printf("Failed to ViewSubscribers - strconv.Atoi error %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	users, err := h.service.ListSubscribers(authorID)
	if err != nil {
		log.Printf("Failed to ViewSubscribers - service.ListSubscribers error %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if users == nil {
		c.JSON(http.StatusOK, gin.H{"message": "User does not have subscribers yet"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": users})
}
