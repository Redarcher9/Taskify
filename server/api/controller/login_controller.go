package controller

import (
	"net/http"
	"taskify/domain"
	"taskify/utils"

	"github.com/gin-gonic/gin"
)

/*
Description: Login Controller struct
*/

type LoginController struct {
	LoginUsecase domain.LoginUsecase
	HMAC_KEY     string
}

/*
Description: Login Controller reciever function

Params: gin context passed from router
*/

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
	ss, _ := utils.CreateToken(c, requestPayload.Email, lc.HMAC_KEY)
	c.SetCookie("Auth_Token", ss, 3600, "/", "localhost", false, true)
	c.JSON(http.StatusAccepted, gin.H{
		"Message": "User has been Logged in!",
	})
}
