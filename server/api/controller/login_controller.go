package controller

import (
	"net/http"
	"taskify/domain"

	"github.com/gin-gonic/gin"
)

type LoginController struct {
	LoginUsecase domain.LoginUsecase
}

func (lc *LoginController) Login(c *gin.Context) {
	var requestPayload domain.Loginstruct
	c.ShouldBind(&requestPayload)
	err := lc.LoginUsecase.Login(c, requestPayload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{
		"Message": "User has been Logged in!",
	})
}
