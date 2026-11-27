package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"project/internal/model"
	"strconv"
)

func (h *BlogHandler) SignUp(c *gin.Context) {
	var user model.User

	err := c.BindJSON(&user)
	if err != nil {
		log.Printf("SignUp - c.BindJSON error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.UserID = 0

	createdUser, err := h.service.CreateUser(&user)
	if err != nil {
		log.Printf("SignUp - h.service.CreateUser error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": createdUser})
}

func (h *BlogHandler) SignIn(c *gin.Context) {
	var user model.User

	if err := c.BindJSON(&user); err != nil {
		log.Printf("SignIn - c.BindJSON error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.service.SignIn(&user)
	if err != nil {
		log.Printf("SignIn - h.service.SignIn error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (h *BlogHandler) GetUsers(c *gin.Context) {

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
	users, err := h.service.ListUsers(page)
	if err != nil {
		log.Printf("Failed to Get Users - strconv.Atoi error %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": users})
}

func (h *BlogHandler) GetUserByID(c *gin.Context) {
	// получаем пользователя
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("Failed to GetUserByID - strconv.Atoi error %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := h.service.FindUserByID(id)
	if err != nil {
		log.Printf("Failed to GetUserByID - service.FindUserByID error %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func (h *BlogHandler) GetUserByUsername(c *gin.Context) {

	username := c.Param("username")
	users, err := h.service.UsersByUsername(username)
	if err != nil {
		log.Printf("Failed to search Users by username - service.UsersByUsername error %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": users})

}

func (h *BlogHandler) EditUser(c *gin.Context) {
	var user model.User
	if err := c.BindJSON(&user); err != nil {
		log.Printf("Failed to EditUser - BindJSON error %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	value, exists := c.Get("user_id")
	if !exists {
		log.Printf("Failed to EditUser - Unauthorized error\n")
		c.JSON(http.StatusUnauthorized, gin.H{"error": fmt.Errorf("User is not logged in")})
		return
	}
	userID := value.(int)

	updatedUser, err := h.service.UpdateUser(&user, userID)
	if err != nil {
		log.Printf("Failed to EditUser - service.UpdateUser error %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": updatedUser})
}

func (h *BlogHandler) GetPostsByAuthor(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("Failed to GetPostsByAuthor - strconv.Atoi error %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
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
	posts, err := h.service.FindPostsByAuthorID(id, page)
	if err != nil {
		log.Printf("Failed to GetPostsByAuthor - service.FindPostsByAuthorID error %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": posts})
}

func (h *BlogHandler) GetDrafts(c *gin.Context) {
	authorID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("Failed to GetDrafts - strconv.Atoi error %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	value, exists := c.Get("user_id")
	if !exists {
		log.Printf("Failed to GetDrafts - Unauthorized error\n")
		c.JSON(http.StatusUnauthorized, gin.H{"error": fmt.Errorf("User is not logged in")})
		return
	}
	userID := value.(int)
	page := 1
	if c.Param("page-id") != "" {
		page, err = strconv.Atoi(c.Param("page-id"))
		if err != nil {
			log.Printf("Failed to GetDrafts - strconv.Atoi error %v\n", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}
	posts, err := h.service.ListDrafts(authorID, userID, page)
	if err != nil {
		log.Printf("Failed to GetDrafts - service.ListDrafts error %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": posts})
}

func (h *BlogHandler) DeleteUser(c *gin.Context) {
	//deletedID, err := strconv.Atoi(c.Param("id"))
	//if err != nil {
	//	log.Printf("Failed to DeleteUser - strconv.Atoi error %v\n", err)
	//	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//	return
	//}

	value, exists := c.Get("user_id")
	if !exists {
		log.Printf("Failed to DeleteUser - c.Get error: key does not exist\n")
		c.JSON(http.StatusUnauthorized, gin.H{})
		return
	}
	userID := value.(int)
	//if deletedID != value.(int) {
	//	log.Printf("Failed to DeleteUser - User is unauthorized")
	//	c.JSON(http.StatusUnauthorized, gin.H{"error": fmt.Errorf("User is unauthorized for the action\n")})
	//	return
	//}
	err := h.service.DeleteUser(userID)
	if err != nil {
		log.Printf("Failed to DeleteUser - service.DeleteUser error %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})

}
