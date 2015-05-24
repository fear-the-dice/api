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
		monsters.GET("/:id", this.getMonster)
	}
}

func (this *monsterController) getMonsters(c *gin.Context) {
	monsters, err := models.PopulateMonsters()
	if err != nil {
		fmt.Errorf("%s", err)
	}

	c.JSON(http.StatusOK, monsters)
}

func (this *monsterController) getMonster(c *gin.Context) {
	monster, err := models.FindMonster(c.Params.ByName("id"))
	if err != nil {
		fmt.Errorf("%s", err)
	}

	if monster == nil {
		c.String(http.StatusNotFound, "")
	} else {
		c.JSON(http.StatusOK, monster)
	}
}
