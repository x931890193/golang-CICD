package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"golang-CICD/lib"
)

func TestHandler(c *gin.Context) {
	resp := Response{}
	lib.Logger.WithField("sss", "s").Error("dsadasda")
	msg := CodeMsg{Code:ParamsError, ExtraMsg:"不能为空！"}
	resp.SetMsg(msg)
	c.JSON(http.StatusOK, resp)
}

