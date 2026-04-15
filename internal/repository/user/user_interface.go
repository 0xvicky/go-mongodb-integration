package repository

import (
	"go-mongo/internal/model"

	"github.com/gin-gonic/gin"
)

type UserInterface interface {
	CreateUser(c *gin.Context, newUser model.User) (any, error)
}
