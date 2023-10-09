package routes

import (
	"database/sql"
	"fmt"
	"strconv"
	"taskify/api/controller"
	"taskify/repository"
	"taskify/usecase"

	"github.com/gin-gonic/gin"
)

func NewLoginRoute(env map[string]string, db *sql.DB, r *gin.RouterGroup) {
	timeout, err := strconv.Atoi(env["CONTEXT_TIMEOUT"])
	if err != nil {
		fmt.Println(err)
	}
	lr := repository.NewUserRepository(db, timeout)
	lc := &controller.LoginController{
		LoginUsecase: usecase.NewLoginUsecase(lr, timeout),
	}
	r.POST("/login", lc.Login)
}
