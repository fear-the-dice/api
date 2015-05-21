package main

import (
	"flag"
	"github.com/gin-gonic/gin"

	"github.com/fear-the-dice/api/tree/real-api/controllers"
)

func main() {
	var (
		port   = flag.String("PORT", ":3000", "The server port")
		router = gin.Default()
	)

	flag.Parse()

	controllers.PlayerController.Attach(router)
	controllers.MonsterController.Attach(router)

	router.Run(*port)
}
