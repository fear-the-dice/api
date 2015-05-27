package main

import (
	"os"

	"github.com/fear-the-dice/api/controllers"
	"github.com/gin-gonic/gin"
	"github.com/obihann/gin-cors"
)

func main() {
	port := ":" + os.Getenv("PORT")
	router := gin.New()
	router.Use(cors.Middleware(cors.Options{
		AllowHeaders: []string{"Origin", "Accept", "Content-Type", "Authorization", "X-HTTP-Method-Override"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "HEAD"},
	}))

	controllers.PlayerController.Attach(router)
	controllers.MonsterController.Attach(router)

	router.Run(port)
}
