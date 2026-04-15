package main

import (
	"context"
	"fmt"
	"go-mongo/internal/controller"
	"go-mongo/internal/db"
	repository "go-mongo/internal/repository/user"
	"go-mongo/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	envErr := godotenv.Load()
	if envErr != nil {
		fmt.Println("[ERROR] error while loading env", envErr)
	}

	//load db
	client, dbErr := db.InitDB()
	if dbErr != nil {
		fmt.Println("[ERROR] ", dbErr)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
			// return fmt.Errorf("error while disconnection:%w", err)
		}
	}()

	db := client.Database("users")
	collection := db.Collection("users-collection")

	fmt.Println("Gin with MongoDB")
	//routes
	r := gin.Default()

	//instances
	repo := repository.NewUserRepository(collection)
	svc := service.NewUserService(repo)
	h := controller.NewUserController(svc)

	r.GET("/", h.Hello)
	r.POST("/newuser", h.CreateUser)
	r.Run()
	//connect db

	//start
	fmt.Println("server listening at 8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		fmt.Println("[ERROR] error while starting server:", err)
	}

}
