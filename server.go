package main

import (
	"github.com/gin-gonic/gin"
	"os"

	"github.com/fear-the-dice/api/controllers"
)

func main() {
	var (
		port   = ":" + os.Getenv("PORT")
		router = gin.Default()
	)

	controllers.PlayerController.Attach(router)
	controllers.MonsterController.Attach(router)

	router.Run(port)
}
