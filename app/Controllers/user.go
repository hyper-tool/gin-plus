package Controllers

import (
	"gin-plus/app/Models"
	"github.com/gin-gonic/gin"
)

func User(c *gin.Context) {

	user := c.DefaultQuery("user", "ben")
	age := c.Query("age")
	Models.Create(user, age)
	c.JSON(200, gin.H{"User": user})
}
