package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"project/internal/model"
	"strconv"
)

func (h *BlogHandler) AddProfile(c *gin.Context) {
	var profile model.Profile

	if err := c.BindJSON(&profile); err != nil {
		log.Printf("Failed to AddProfile - BindJSON error %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	value, exists := c.Get("user_id")
	if value == nil || !exists {
		log.Printf("Failed to AddProfile - Unauthorized error\n")
		c.JSON(http.StatusUnauthorized, gin.H{"error": fmt.Errorf("User is not logged in")})
		return
	}
	userID := value.(int)

	p, err := h.service.CreateProfile(&profile, userID)
	if err != nil {
		log.Printf("Failed to AddProfile - service.CreateProfile error\n")
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Errorf("User is not logged in")})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": p})
}

func (h *BlogHandler) UpdateProfile(c *gin.Context) {
	var profile model.Profile

	if err := c.BindJSON(&profile); err != nil {
		log.Printf("Failed to UpdateProfile - BindJSON error %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	value, exists := c.Get("user_id")
	if value == nil || !exists {
		log.Printf("Failed to UpdateProfile - Unauthorized error\n")
		c.JSON(http.StatusUnauthorized, gin.H{"error": fmt.Errorf("User is not logged in")})
		return
	}
	userID := value.(int)

	p, err := h.service.UpdateProfile(&profile, userID)
	if err != nil {
		log.Printf("Failed to AddProfile - service.UpdateProfile error\n")
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Errorf("User is not logged in")})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"data": p})
}

func (h *BlogHandler) GetProfile(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("Failed to GetProfile - strconv.Atoi error\n")
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	p, err := h.service.FindProfile(userID)
	if err != nil {
		log.Printf("Failed to GetProfile - service.FindProfile error\n")
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Errorf("User is not logged in")})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"data": p})
}

func (h *BlogHandler) DeleteProfile(c *gin.Context) {
	value, exists := c.Get("user_id")
	if value == nil || !exists {
		log.Printf("Failed to AddProfile - Unauthorized error\n")
		c.JSON(http.StatusUnauthorized, gin.H{"error": fmt.Errorf("User is not logged in")})
		return
	}
	userID := value.(int)
	err := h.service.DeleteProfile(userID)
	if err != nil {
		log.Printf("Failed to DeleteProfile - service.DeleteProfile error\n")
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Errorf("User is not logged in")})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"message": "Profile Deleted"})
}
