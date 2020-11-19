package main

import "github.com/gin-gonic/gin"

func main(){
	g := gin.Default()
	g.GET("/", func(c *gin.Context) {
		c.JSON(200,gin.H{
			"version": "v1",
			"data": "这是第一版，Hello Joker",
		})
	})
	g.Run(":8080")
}
