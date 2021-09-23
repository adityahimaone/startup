package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"startup/auth"
	"startup/handler"
	"startup/helper"
	"startup/user"
	"strings"
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

	userHandler := handler.NewUserHandler(userService, authService)

	router := gin.Default()
	//api versioning
	api := router.Group("/api/v1")

	/*END POINT*/
	//akan dialihkan ke register user
	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvailable)
	api.POST("/avatars", authMiddleware(authService, userService), userHandler.UploudAvatar)
	router.Run()
}

func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeaders := c.GetHeader("Authorization")

		if !strings.Contains(authHeaders, "Bearer") {
			response := helper.APIResponse("Unathorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		//Bearer token
		var tokenString string = ""
		arrayToken := strings.Split(authHeaders, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		//validasi token
		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.APIResponse("Unathorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		payload, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			response := helper.APIResponse("Unathorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userID := int(payload["user_id"].(float64))

		user, err := userService.GetUserByID(userID)

		if err != nil {
			response := helper.APIResponse("Unathorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("currentUser", user)
	}
}

//komentar

// ambil nilai headr authorization: Bearer Token JWT
// dari header authorization, kita ambil nilai tokennya
// validasi token
// ambil user_id
// ambil user dari db berdasarkan user_id lewat service
// if user ada set context isinya user ( context -> tempat untuk menyimpan nilai yg bisa diakses dr tempat lain)

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
}

token, err := authService.ValidateToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo1fQ.4G2N7tjtshRsA3ry4yrCCl5nMyfEftvdVwmFEOc2Vvg")
	if err != nil {
		fmt.Println("ERROR")
	}
	if token.Valid {
		fmt.Println("VALID")
	} else {
		fmt.Println("INVALID")
	}

*/
