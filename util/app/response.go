package app

import "github.com/gin-gonic/gin"

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (g *Gin) Response(httpCode int, err error, data interface{}) {
	g.C.IndentedJSON(httpCode, Response{
		Code:    httpCode,
		Message: err.Error(),
		Data:    data,
	})
}
