package controller

import (
	"net/http"
	"taskify/domain"
	"taskify/utils"

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
	ss, _ := utils.CreateToken(c, requestPayload.Email)
	c.SetCookie("Auth_Token", ss, 3600, "/", "localhost", false, true)
	_ = utils.ValidateToken(c, ss)
	c.JSON(http.StatusAccepted, gin.H{
		"Message": "User has been Logged in!",
	})
}
