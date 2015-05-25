package controllers

import (
	"fmt"
	"net/http"

	"github.com/fear-the-dice/api/models"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

type playerController struct{}

var PlayerController *playerController = new(playerController)

// Attach accepts a Gin object and then creates all routes
func (this *playerController) Attach(router *gin.Engine) {
	players := router.Group("/players")
	{
		players.POST("", this.newPlayer)
		players.GET("", this.getPlayers)
		players.GET("/:id", this.getPlayer)
		players.PUT("/:id", this.updatePlayer)
		players.DELETE("/:id", this.deletePlayer)
	}
}

func (this *playerController) getPlayers(c *gin.Context) {
	players, err := models.PopulatePlayers()
	if err != nil {
		fmt.Printf("%s\n", err)
	}

	if players == nil {
		c.String(http.StatusNotFound, "")
	} else {
		c.JSON(http.StatusOK, players)
	}
}

func (this *playerController) getPlayer(c *gin.Context) {
	oid := bson.ObjectIdHex(c.Params.ByName("id"))
	player, err := models.FindPlayer(oid)
	if err != nil {
		fmt.Printf("%s\n", err)
	}

	if player == nil {
		c.String(http.StatusNotFound, "")
	} else {
		c.JSON(http.StatusOK, player)
	}
}

func (this *playerController) newPlayer(c *gin.Context) {
	var playerJSON *models.Player
	playerJSON = models.NewPlayer()
	c.Bind(&playerJSON)

	player, err := models.InsertPlayer(*playerJSON)
	if err != nil {
		fmt.Printf("%s\n", err)
		c.String(http.StatusNotFound, "")
	} else {
		url := fmt.Sprintf("fear-the-dice-api.herokuapp.com/%s", player.ID)
		c.Writer.Header().Set("Location", url)
		c.JSON(http.StatusCreated, player)
	}
}

func (this *playerController) deletePlayer(c *gin.Context) {
	oid := bson.ObjectIdHex(c.Params.ByName("id"))
	if err := models.DeletePlayer(oid); err != nil {
		fmt.Printf("%s\n", err)
		c.String(http.StatusNotFound, "")
	} else {
		c.String(http.StatusOK, "")
	}
}

func (this *playerController) updatePlayer(c *gin.Context) {
	oid := bson.ObjectIdHex(c.Params.ByName("id"))
	var player models.Player
	c.Bind(&player)

	if err := models.UpdatePlayer(oid, player); err != nil {
		fmt.Printf("%s\n", err)
		c.String(http.StatusNotFound, "")
	} else {
		c.JSON(http.StatusOK, player)
	}
}
