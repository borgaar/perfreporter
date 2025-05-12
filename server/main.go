package main

import (
	"github.com/borgaar/perfreporter/api"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	router := gin.Default()

	authorized := router.Group("/")

	authPasswd := os.Getenv("PASSWD")

	if authPasswd == "" {
		log.Println("PASSWD env variable not found. Setting PASSWD=development")
		authPasswd = "development"
	}

	authorized.Use(func(ctx *gin.Context) {
		passwd := ctx.GetHeader("passwd")

		if passwd != authPasswd {
			ctx.JSON(401, gin.H{
				"message": "Bad credentials"},
			)
			ctx.Abort()
		} else {
			ctx.Next()
		}
	})
	{
		authorized.GET("/api/stats", func(c *gin.Context) {
			api.GetStats(c)
		})
	}

	runErr := router.Run("0.0.0.0:8080")

	if runErr != nil {
		log.Fatal(runErr)
	}
}
