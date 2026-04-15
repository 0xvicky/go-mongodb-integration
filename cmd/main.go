package main

import (
	"fmt"
	"go-mongo/internal/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	fmt.Println("Gin with MongoDB")
	//routes
	r := gin.Default()

	r.GET("/", controller.Hello)
	r.POST("/newuser", controller.CreateUser)
	r.Run()
	//connect db

	//start
	fmt.Println("server listening at 8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		fmt.Println("[ERROR] error while starting server:", err)
	}

}
