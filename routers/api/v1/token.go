package v1

import (
	"com/EasyApp_Server/util/app"
	"com/EasyApp_Server/util/e"
	"github.com/gin-gonic/gin"
	"net/http"
)

func checkToken(c *gin.Context) {
	app := app.Gin{C: c}
	app.Response(http.StatusOK, e.OK, nil)
}
