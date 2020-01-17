package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/sysu-lemon/EasyApp_Server/util/app"
	"github.com/sysu-lemon/EasyApp_Server/util/e"
	"net/http"
)

func checkToken(c *gin.Context) {
	app := app.Gin{C: c}
	app.Response(http.StatusOK, e.OK, nil)
}
