package repository

import (
	"fmt"
	"go-mongo/internal/model"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
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

func (m *UserRepository) GetUserById(c *gin.Context, userId bson.ObjectID) (any, error) {
	filter := bson.M{"_id": userId}
	var user model.User
	err := m.MongoCollection.FindOne(c, filter).Decode(&user)
	if err != nil {
		return nil, fmt.Errorf("user fetch err:%w", err)
	}

	return user, nil
}
