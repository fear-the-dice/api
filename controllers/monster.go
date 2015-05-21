package controllers

import (
	"fmt"
	"net/http"

	"github.com/fear-the-dice/api/models"
	"github.com/gin-gonic/gin"
)

type monsterController struct{}

var MonsterController *monsterController = new(monsterController)

// Attach accepts a Gin object and then creates all routes
func (this *monsterController) Attach(router *gin.Engine) {
	monsters := router.Group("/monsters")
	{
		monsters.GET("", this.getMonsters)
	}
}

func (this *monsterController) getMonsters(c *gin.Context) {
	monsters, err := models.PopulateMonsters()
	if err != nil {
		fmt.Printf("%s", err.Error())
	}

	c.JSON(http.StatusOK, monsters)
}
