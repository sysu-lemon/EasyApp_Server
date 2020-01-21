package routers

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/sysu-lemon/EasyApp_Server/routers/api/v1"
)

func Init() *gin.Engine {
	//gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	api := r.Group("/api")
	{
		v1.Register(api)
	}
	return r
}
