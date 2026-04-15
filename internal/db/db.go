package db

import (
	"context"
	"fmt"
	"os"

	// "github.com/mongodb/mongo-go-driver/mongo/options"
	// "github.com/mongodb/mongo-go-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

func InitDB() (*mongo.Client, error) {
	// Use the SetServerAPIOptions() method to set the version of the Stable API on the client
	var MONGO_URI = os.Getenv("MONGO_URI")
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(MONGO_URI).SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	client, err := mongo.Connect(opts)
	if err != nil {
		// panic(err)
		return nil, fmt.Errorf("error while mongo connect:%w", err)
	}

	// Send a ping to confirm a successful connection
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		// panic(err)
		return nil, fmt.Errorf("error while pinging:%w", err)

	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	return client, nil
}
