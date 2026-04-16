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

func (m *UserRepository) GetAllUsers(c *gin.Context) ([]any, error) {
	var users []any
	cursor, findErr := m.MongoCollection.Find(c, bson.M{})

	if findErr != nil {
		return nil, fmt.Errorf("error while fetching cursor:%w", findErr)
	}

	if err := cursor.All(c, &users); err != nil {
		return nil, fmt.Errorf("error while fetching all users:%w", err)
	}

	return users, nil
}

func (m *UserRepository) UpdateUser(c *gin.Context, user model.User) (any, error) {
	updates := bson.M{
		"$set": bson.M{
			"name": user.Name,
		},
	}
	updateRes, findErr := m.MongoCollection.UpdateByID(c, user.UserId, updates)

	if findErr != nil {
		return nil, fmt.Errorf("error while fetching cursor:%w", findErr)
	}
	return updateRes, nil
}
func (m *UserRepository) DeleteUserById(c *gin.Context, userId bson.ObjectID) (any, error) {
	filter := bson.M{"_id": userId}
	delRes, delErr := m.MongoCollection.DeleteOne(c, filter)
	if delErr != nil {
		return nil, fmt.Errorf("delete user failed:%w", delErr)
	}

	return delRes, nil
}
func (m *UserRepository) DeleteAllUsers(c *gin.Context) (any, error) {
	filter := bson.M{}

	delAllRes, delAllErr := m.MongoCollection.DeleteMany(c, filter)
	if delAllErr != nil {
		return nil, fmt.Errorf("delete user failed:%w", delAllErr)
	}

	return delAllRes, nil
}
