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

func NewPingRoute(env map[string]string, db *sql.DB, r *gin.RouterGroup) {
	timeout, err := strconv.Atoi(env["CONTEXT_TIMEOUT"])
	if err != nil {
		fmt.Println(err)
	}
	pr := repository.NewPingRepository(db, timeout)
	pc := &controller.PingController{
		PingUsecase: usecase.NewPingUsecase(pr, timeout),
	}
	r.GET("/ping", pc.Ping)
}
