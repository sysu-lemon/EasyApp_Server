package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/sysu-lemon/EasyApp_Server/middleware"
)

func Register(router *gin.RouterGroup) {
	v1 := router.Group("/v1")
	{
		v1.POST("/login", signIn)
		v1.POST("/users", signUp)

		v1.GET("/token", checkToken).Use(middleware.JWT())
	}
}
