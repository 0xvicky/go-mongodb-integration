package model

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type User struct {
	UserId    bson.ObjectID `bson:"_id" json:"userId"`
	Email     string        `bson:"email" json:"email,omitempty"`
	Name      string        `bson:"name" json:"name,omitempty"`
	Password  string        `bson:"password" json:"password,omitempty"`
	CreatedAt time.Time     `bson:"created_at" json:"created_at,omitempty"`
}
