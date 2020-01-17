package middleware

import (
	"com/gin-learning/util/e"
	"github.com/gin-gonic/gin"
	"github.com/sysu-lemon/EasyApp_Server/util/app"
	"github.com/sysu-lemon/EasyApp_Server/util/jwt"
	"net/http"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		app := app.Gin{C: c}

		token := c.GetHeader("token")
		claims, err := jwt.ParseToken(token)

		switch {
		case token == "":
			app.Response(http.StatusUnauthorized, e.ErrTokenEmpty, nil)
		case err != nil:
			app.Response(http.StatusUnauthorized, e.ErrTokenInvalid, nil)
		case time.Now().Unix() > claims.ExpiresAt:
			app.Response(http.StatusUnauthorized, e.ErrTokenTimeout, nil)
		default:
			c.Next()
		}
		c.Abort()
	}
}
