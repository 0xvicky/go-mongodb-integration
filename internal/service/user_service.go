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

func (svc UserService) GetAllUsers(c *gin.Context) ([]any, error) {

	users, fetchErr := svc.Repository.GetAllUsers(c)
	if fetchErr != nil {
		return nil, fetchErr
	}

	return users, nil
}
func (svc UserService) UpdateUser(c *gin.Context, userId string, user model.User) (any, error) {
	userIdObj, bsonErr := bson.ObjectIDFromHex(userId)
	if bsonErr != nil {
		return model.User{}, fmt.Errorf("id conversion error:%w", bsonErr)
	}
	user.UserId = userIdObj
	users, fetchErr := svc.Repository.UpdateUser(c, user)
	if fetchErr != nil {
		return nil, fetchErr
	}

	return users, nil
}
func (svc UserService) DeleteById(c *gin.Context, userId string) (any, error) {
	userIdObj, bsonErr := bson.ObjectIDFromHex(userId)
	if bsonErr != nil {
		return model.User{}, fmt.Errorf("id conversion error:%w", bsonErr)
	}

	delRes, delErr := svc.Repository.DeleteUserById(c, userIdObj)
	if delErr != nil {
		return nil, delErr
	}

	return delRes, nil
}
func (svc UserService) DeleteAll(c *gin.Context) (any, error) {

	delRes, delErr := svc.Repository.DeleteAllUsers(c)
	if delErr != nil {
		return nil, delErr
	}

	return delRes, nil
}
