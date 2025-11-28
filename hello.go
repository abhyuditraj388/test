package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "hello world",
		})

	})

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "home",
		})
	})

	r.POST("/add", func(c *gin.Context) {
		var user struct {
			Name string `json:"name"`
		}

		err := c.BindJSON(&user)
		if err != nil {
			c.JSON(500, http.StatusBadRequest)
		}

		c.JSON(200, gin.H{
			"message": "hello",
			"user":    user.Name,
		})
	})

	r.Run(":12710")

}
