package repository

import (
	"go-mongo/internal/model"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type UserInterface interface {
	CreateUser(c *gin.Context, newUser model.User) (any, error)
	GetUserById(c *gin.Context, userId bson.ObjectID) (any, error)
	GetAllUsers(c *gin.Context) ([]any, error)
	UpdateUser(c *gin.Context, user model.User) (any, error)
	DeleteUserById(c *gin.Context, userId bson.ObjectID) (any, error)
	DeleteAllUsers(c *gin.Context) (any, error)
}
