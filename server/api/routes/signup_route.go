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

func NewSignupRoute(env map[string]string, db *sql.DB, r *gin.RouterGroup) {
	timeout, err := strconv.Atoi(env["CONTEXT_TIMEOUT"])
	if err != nil {
		fmt.Println(err)
	}
	sr := repository.NewUserRepository(db, timeout)
	sc := &controller.SignUpController{
		SignupUsecase: usecase.NewSignupUsecase(sr, timeout),
	}
	r.POST("/signup", sc.SignUp)
}
