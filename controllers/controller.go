package controllers

import (
	"ScreenHttpServer/tcp"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	BaseController
}

func (this *Controller) Getdata(c *gin.Context) {
	screenDat := tcp.Tcp_main()
	c.JSON(http.StatusOK, screenDat)
	
}
