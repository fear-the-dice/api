package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/fear-the-dice/api/controllers"
	"github.com/gin-gonic/gin"
	"github.com/obihann/gin-cors"
)

func checkAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		verifyKey := []byte("secretkey")
		auth := c.Request.Header.Get("Authorization")
		auth = strings.Trim(auth[6:], " ")

		token, err := jwt.Parse(auth, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			return verifyKey, nil
		})

		if err != nil {
			fmt.Printf("Error parsing token: %s\n", err)
			return
		}

		iss := token.Claims["iss"].(string)
		fmt.Printf("iss: %s\n", iss)
	}
}

func main() {
	port := ":" + os.Getenv("PORT")
	router := gin.New()
	router.Use(cors.Middleware(cors.Options{
		AllowHeaders: []string{"Origin", "Accept", "Content-Type", "Authorization", "X-HTTP-Method-Override"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "HEAD"},
	}))

	router.Use(checkAuth())

	controllers.PlayerController.Attach(router)
	controllers.MonsterController.Attach(router)

	router.Run(port)
}
