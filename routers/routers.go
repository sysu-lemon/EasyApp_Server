package routers

import (
	v1 "com/EasyApp_Server/routers/api/v1"
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	r := gin.Default()
	api := r.Group("/api")
	{
		v1.Register(api)
	}
	return r
}