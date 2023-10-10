package controller

import (
	"fmt"
	"net/http"
	"taskify/domain"

	"github.com/gin-gonic/gin"
)

/*
Description: Signup Controller struct
*/

type SignUpController struct {
	SignupUsecase domain.SignupUsecase
}

/*
Description: Signup Controller reciever function

Params: gin context passed from router
*/

func (sc *SignUpController) SignUp(c *gin.Context) {
	var requestPayload domain.UserStruct
	c.ShouldBind(&requestPayload)

	fmt.Println(requestPayload)
	err := sc.SignupUsecase.SignUp(c, requestPayload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{
		"Message": "User has been created!",
	})
}
