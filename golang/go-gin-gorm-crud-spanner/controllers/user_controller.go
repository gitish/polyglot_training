package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gitish/polyglot_training/models"
	"github.com/gitish/polyglot_training/services"
)

// UserController defines the interface for user-related API operations.
type UserController interface {
	CreateUser(c *gin.Context)
	GetUsers(c *gin.Context)
	GetUserByID(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
}

// userControllerImpl implements the UserController interface.
type userControllerImpl struct {
	userService services.UserService
}

// NewUserController creates a new instance of userControllerImpl.
func NewUserController(userService services.UserService) UserController {
	return &userControllerImpl{userService: userService}
}

// CreateUser handles the POST request to create a user.
func (ctrl *userControllerImpl) CreateUser(c *gin.Context) {
	var newUser models.User
	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	createdUser, err := ctrl.userService.CreateUser(newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}
	c.JSON(http.StatusCreated, createdUser)
}

// GetUsers handles the GET request to retrieve all users.
func (ctrl *userControllerImpl) GetUsers(c *gin.Context) {
	users, err := ctrl.userService.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user(s)"})
		return
	}
	c.JSON(http.StatusOK, users)
}

// GetUserByID handles the GET request to retrieve a user by ID.
func (ctrl *userControllerImpl) GetUserByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	user, err := ctrl.userService.GetUserByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Failed to get user"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// UpdateUser handles the PUT request to update a user by ID.
func (ctrl *userControllerImpl) UpdateUser(c *gin.Context) {
	var updateUser models.User
	if err := c.BindJSON(&updateUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	updatedUser, err := ctrl.userService.UpdateUser(updateUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}
	c.JSON(http.StatusOK, updatedUser)
}

// DeleteUser handles the DELETE request to remove a user by ID.
func (ctrl *userControllerImpl) DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	err = ctrl.userService.DeleteUser(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Failed to delete user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
