package controllers

import (
	"fmt"
	"net/http"

	"github.com/fear-the-dice/api/models"
	"github.com/gin-gonic/gin"
)

type playerController struct{}

var PlayerController *playerController = new(playerController)

// Attach accepts a Gin object and then creates all routes
func (this *playerController) Attach(router *gin.Engine) {
	players := router.Group("/players")
	{
		players.GET("", this.getPlayers)
	}
}

func (this *playerController) getPlayers(c *gin.Context) {
	players, err := models.PopulatePlayers()
	if err != nil {
		fmt.Printf("%s", err.Error())
	}

	c.JSON(http.StatusOK, players)
}
