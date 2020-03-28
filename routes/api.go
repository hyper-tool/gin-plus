package routes

import (
	"gin-plus/app/Controllers"
	"github.com/gin-gonic/gin"
)

func Api() *gin.Engine {
	r := gin.Default()
	r.GET("user", Controllers.User)
	return r
}
