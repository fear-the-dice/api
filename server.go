package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/fear-the-dice/api/controllers"
	"github.com/fear-the-dice/api/models"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"github.com/obihann/gin-cors"
)

func checkAuth(pool *redis.Pool, verifyKey []byte) gin.HandlerFunc {
	return func(c *gin.Context) {
		conn := pool.Get()
		defer conn.Close()

		auth := c.Request.Header.Get("Authorization")
		auth = strings.Trim(auth[6:], " ")

		token, err := jwt.Parse(auth, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			return verifyKey, nil
		})

		if err != nil {
			c.Fail(http.StatusUnauthorized, errors.New("Error parsing token"))
			return
		}

		sub := token.Claims["sub"].(string)

		if sub != "8205d3b2-3e73-432b-b7eb-b73f73818d83" {
			c.Fail(http.StatusUnauthorized, errors.New("Invalid subscriber"))
			return
		}

		jti := token.Claims["jti"].(string)
		key := fmt.Sprintf("ftd/%s/%s", sub, jti)

		conn.Send("EXISTS", key)
		conn.Flush()
		_, err = conn.Receive()

		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			c.Fail(http.StatusUnauthorized, errors.New("JTI expired or invalid"))
			return
		}

		c.Next()
		return
	}
}

func main() {
	type ()

	var (
		redisServer = os.Getenv("REDISCLOUD_URL")
		port        = os.Getenv("PORT")
		mongoServer = os.Getenv("MONGOLAB_URI")
		mongoDb     = os.Getenv("DB")
		authKey     = os.Getenv("FTDAUTHKEY")
	)

	if len(redisServer) <= 1 {
		redisServer = ":6379"
	}

	if len(port) <= 1 {
		port = "3000"
	}

	if len(mongoDb) <= 1 {
		mongoDb = "heroku_app37083199"
	}

	if len(mongoServer) <= 1 {
		mongoServer = "mongodb://localhost"
	}

	if len(authKey) <= 1 {
		authKey = "supersecret"
	}

	pool := &redis.Pool{
		MaxIdle: 100,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", redisServer)

			if err != nil {
				return nil, err
			}

			return c, nil
		},
	}

	dbOptions := models.DbOptions{
		Host:     mongoServer,
		Database: mongoDb,
	}

	models.SetConfig(&dbOptions)

	router := gin.New()

	router.Use(cors.Middleware(cors.Options{
		AllowHeaders: []string{"Origin", "Accept", "Content-Type", "Authorization", "X-HTTP-Method-Override"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "HEAD"},
	}))

	router.Use(checkAuth(pool, []byte(authKey)))

	controllers.PlayerController.Attach(router)
	controllers.MonsterController.Attach(router)

	router.Run(fmt.Sprintf(":%s", port))
}
