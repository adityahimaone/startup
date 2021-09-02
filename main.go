package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"startup/handler"
	"startup/user"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/startup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("Connection OK")

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)

	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()
	//api versioning
	api := router.Group("/api/v1")

	//akan dialihkan ke register user
	api.POST("/users", userHandler.RegisterUser)
	router.Run()
	//fill input
	/*	userInput := user.RegisterUserInput{}
		userInput.Name = "Test"
		userInput.Email = "aa@mail.com"
		userInput.Occupation = "soft dev"
		userInput.Password = "secret"
		userService.RegisterUser(userInput)*/

}
