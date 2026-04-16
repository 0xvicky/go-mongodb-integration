package controller

import (
	"go-mongo/internal/model"
	"go-mongo/internal/service"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type UserController struct {
	Service *service.UserService
}

func NewUserController(svc *service.UserService) *UserController {
	return &UserController{
		Service: svc,
	}
}

func (h *UserController) Hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Hello World",
	})
}

func (h *UserController) CreateUser(c *gin.Context) {
	var req model.User
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "failed to create user",
			"error":   err.Error(),
		})
		return
	}
	userId := bson.NewObjectID()
	req.CreatedAt = time.Now()
	req.UserId = userId
	savedUser, createErr := h.Service.CreateUser(c, req)
	if createErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "internal server error",
			"error":   createErr.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    savedUser,
	})
}

func (h *UserController) GetUserById(c *gin.Context) {
	userId := c.Param("id")

	user, err := h.Service.GetUserById(c, userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "internal server error",
			"error":   err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    user,
	})
}
func (h *UserController) GetAllUsers(c *gin.Context) {

	users, err := h.Service.GetAllUsers(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "internal server error",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    users,
	})
}

func (h *UserController) UpdateUser(c *gin.Context) {
	var req model.User
	userId := c.Param("id")
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "internal server error",
			"error":   err.Error(),
		})
	}
	users, err := h.Service.UpdateUser(c, userId, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "internal server error",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    users,
	})
}

func (h *UserController) DeleteById(c *gin.Context) {
	userId := c.Param("id")

	delRes, delErr := h.Service.DeleteById(c, userId)
	if delErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "internal server error",
			"error":   delErr.Error(),
		})

	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    delRes,
	})

}

func (h *UserController) DeleteAll(c *gin.Context) {

	delRes, delErr := h.Service.DeleteAll(c)
	if delErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "internal server error",
			"error":   delErr.Error(),
		})

	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    delRes,
	})

}
