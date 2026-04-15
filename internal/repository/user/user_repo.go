package repository

import (
	"fmt"
	"go-mongo/internal/model"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type UserRepository struct {
	MongoCollection *mongo.Collection
}

func NewUserRepository(m *mongo.Collection) UserInterface {
	return &UserRepository{
		MongoCollection: m,
	}
}

func (m *UserRepository) CreateUser(c *gin.Context, newUser model.User) (any, error) {
	savedUser, insertErr := m.MongoCollection.InsertOne(c, newUser)
	if insertErr != nil {
		return nil, fmt.Errorf("error occured while inserting user:%w", insertErr)
	}

	return savedUser, nil
}
