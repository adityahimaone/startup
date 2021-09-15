package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"startup/auth"
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
	authService := auth.NewService()

	token, err := authService.ValidateToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo1fQ.4G2N7tjtshRsA3ry4yrCCl5nMyfEftvdVwmFEOc2Vvg")
	if err != nil {
		fmt.Println("ERROR")
	}
	if token.Valid {
		fmt.Println("VALID")
	} else {
		fmt.Println("INVALID")
	}
	userHandler := handler.NewUserHandler(userService, authService)

	router := gin.Default()
	//api versioning
	api := router.Group("/api/v1")

	/*END POINT*/
	//akan dialihkan ke register user
	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvailable)
	api.POST("/avatars", userHandler.UploudAvatar)
	router.Run()
}

//komentar

//fill input
/*	userInput := user.RegisterUserInput{}
	userInput.Name = "Test"
	userInput.Email = "aa@mail.com"
	userInput.Occupation = "soft dev"
	userInput.Password = "secret"
	userService.RegisterUser(userInput)*/

/*input := user.LoginInput{
	Email:    "aa@mail.com",
	Password: "secret",
}
user, err := userService.Login(input)
if err != nil {
	fmt.Println(err.Error())
	fmt.Println("Salah")
}
fmt.Println(user.Email)
fmt.Println(user.Name)*/
/*userByEmail, err := userRepository.FindByEmail("adiyt@mail.com")
if err != nil {
	fmt.Println(err.Error())
}
if userByEmail.ID == 0 {
	fmt.Println("user not found")
} else {
	fmt.Println(userByEmail.Name)
}*/
