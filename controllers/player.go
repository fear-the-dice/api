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
		players.GET("/:id", this.getPlayer)
	}
}

func (this *playerController) getPlayers(c *gin.Context) {
	players, err := models.PopulatePlayers()
	if err != nil {
		fmt.Errorf("%s", err)
	}

	if players == nil {
		c.String(http.StatusNotFound, "")
	} else {
		c.JSON(http.StatusOK, players)
	}
}

func (this *playerController) getPlayer(c *gin.Context) {
	player, err := models.FindPlayer(c.Params.ByName("id"))
	if err != nil {
		fmt.Errorf("%s", err)
	}

	if player == nil {
		c.String(http.StatusNotFound, "")
	} else {
		c.JSON(http.StatusOK, player)
	}
}
