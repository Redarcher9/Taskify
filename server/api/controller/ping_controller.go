package controller

import (
	"net/http"
	"taskify/domain"

	"github.com/gin-gonic/gin"
)

type PingController struct {
	PingUsecase domain.PingUsecase
}

func (pc *PingController) Ping(c *gin.Context) {
	err := pc.PingUsecase.CheckPing(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{
		"Message": "Server and Database are working just fine!",
	})
}
