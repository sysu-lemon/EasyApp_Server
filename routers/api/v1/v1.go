package v1

import (
	"com/EasyApp_Server/middleware"
	"github.com/gin-gonic/gin"
)

func Register(router *gin.RouterGroup) {
	v1 := router.Group("/v1")
	{
		v1.POST("/login", signIn)
		v1.POST("/users", signUp)

		v1.GET("/token", checkToken).Use(middleware.JWT())
	}
}
