package service

import (
	"fmt"
	"go-mongo/internal/model"
	repository "go-mongo/internal/repository/user"

	"github.com/gin-gonic/gin"
)

type UserService struct {
	Repository repository.UserInterface
}

func NewUserService(r repository.UserInterface) *UserService {
	return &UserService{
		Repository: r,
	}
}

func (svc UserService) CreateUser(c *gin.Context, newUser model.User) (any, error) {
	savedUser, createErr := svc.Repository.CreateUser(c, newUser)
	if createErr != nil {
		return nil, fmt.Errorf("%w", createErr)
	}

	return savedUser, nil
}
