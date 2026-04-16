package service

import (
	"fmt"
	"go-mongo/internal/model"
	repository "go-mongo/internal/repository/user"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
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

func (svc UserService) GetUserById(c *gin.Context, userId string) (any, error) {
	fmt.Println("Cool")
	userIdObj, bsonErr := bson.ObjectIDFromHex(userId)
	if bsonErr != nil {
		return model.User{}, fmt.Errorf("id conversion error:%w", bsonErr)
	}
	user, fetchErr := svc.Repository.GetUserById(c, userIdObj)
	if fetchErr != nil {
		return model.User{}, fetchErr
	}

	return user, nil
}
