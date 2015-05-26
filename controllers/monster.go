package controllers

import (
	"fmt"
	"net/http"

	"github.com/fear-the-dice/api/models"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

type monsterController struct{}

var MonsterController *monsterController = new(monsterController)

// Attach accepts a Gin object and then creates all routes
func (this *monsterController) Attach(router *gin.Engine) {
	monsters := router.Group("/monsters")
	{
		monsters.POST("", this.newMonster)
		monsters.GET("", this.getMonsters)
		monsters.GET("/:id", this.getMonster)
		monsters.PUT("/:id", this.updateMonster)
		monsters.DELETE("/:id", this.deleteMonster)
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
	oid := bson.ObjectIdHex(c.Params.ByName("id"))
	monster, err := models.FindMonster(oid)
	if err != nil {
		fmt.Errorf("%s", err)
	}

	if monster == nil {
		c.String(http.StatusNotFound, "")
	} else {
		c.JSON(http.StatusOK, monster)
	}
}

func (this *monsterController) newMonster(c *gin.Context) {
	var monsterJSON *models.Monster
	monsterJSON = models.NewMonster()
	c.Bind(&monsterJSON)

	monster, err := models.InsertMonster(*monsterJSON)
	if err != nil {
		fmt.Printf("%s\n", err)
		c.String(http.StatusNotFound, "")
	} else {
		url := fmt.Sprintf("fear-the-dice-api.herokuapp.com/%s", monster.ID)
		c.Writer.Header().Set("Location", url)
		c.JSON(http.StatusCreated, monster)
	}
}

func (this *monsterController) deleteMonster(c *gin.Context) {
	oid := bson.ObjectIdHex(c.Params.ByName("id"))
	if err := models.DeleteMonster(oid); err != nil {
		fmt.Errorf("%s", err)
		c.String(http.StatusNotFound, "")
	} else {
		c.String(http.StatusOK, "")
	}
}

func (this *monsterController) updateMonster(c *gin.Context) {
	oid := bson.ObjectIdHex(c.Params.ByName("id"))
	var monster models.Monster
	c.Bind(&monster)

	if err := models.UpdateMonster(oid, monster); err != nil {
		fmt.Printf("%s\n", err)
		c.String(http.StatusNotFound, "")
	} else {
		c.JSON(http.StatusOK, monster)
	}
}
