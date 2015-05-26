package main

import (
	"os"

	"github.com/fear-the-dice/api/controllers"
	"github.com/gin-gonic/gin"
	"github.com/tommy351/gin-cors"
)

func main() {
	port := ":" + os.Getenv("PORT")
	router := gin.New()
	router.Use(cors.Middleware(cors.Options{}))

	controllers.PlayerController.Attach(router)
	controllers.MonsterController.Attach(router)

	router.Run(port)
}
