package Controllers

import "github.com/gin-gonic/gin"

func User(c *gin.Context) {
	user := c.DefaultQuery("user", "ben")
	c.JSON(200, gin.H{"User": user})
}
