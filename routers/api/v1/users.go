package v1

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sysu-lemon/EasyApp_Server/models"
	"github.com/sysu-lemon/EasyApp_Server/util/app"
	"github.com/sysu-lemon/EasyApp_Server/util/e"
	"github.com/sysu-lemon/EasyApp_Server/util/jwt"
	"net/http"
)

func signIn(c *gin.Context) {
	app := app.Gin{C: c}
	var user models.User
	if err := json.NewDecoder(c.Request.Body).Decode(&user); err != nil {
		app.Response(http.StatusBadRequest, err, nil)
		return
	}

	var u models.User
	models.DB.Where("username = ?", user.Username).First(&u)
	if u.Username != user.Username || u.Password != user.Password {
		app.Response(http.StatusBadRequest, e.ErrNameOrPassWrong, nil)
		return
	}
	token, err := jwt.GenerateToken(u)
	if err != nil {
		app.Response(http.StatusInternalServerError, err, nil)
		return
	}
	app.Response(http.StatusOK, e.OK, map[string]string{"token": token})
}

func signUp(c *gin.Context) {
	app := app.Gin{C: c}
	var user models.User
	fmt.Println(c.Request.PostForm)
	if err := json.NewDecoder(c.Request.Body).Decode(&user); err != nil {
		app.Response(http.StatusBadRequest, err, nil)
		return
	}

	if err := models.DB.Create(&user).Error; err != nil {
		app.Response(http.StatusBadRequest, e.ErrUserExisted, nil)
		return
	}
	app.Response(http.StatusOK, e.OK, user)
}
