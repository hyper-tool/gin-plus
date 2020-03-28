package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/test", Handler)
	r.Run()
}

func Handler(c *gin.Context) {
	name := c.DefaultQuery("name", "tizi365")
	id, _ := c.GetQuery("id")
	c.JSON(200, gin.H{"name": name, "id": id})
}
