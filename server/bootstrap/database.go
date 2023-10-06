package bootstrap

import (
	"database/sql"
	"fmt"
	"taskify/database"
)

func SetupDB(EnvVariables map[string]string) *sql.DB {
	db, err := database.NewPsqlDB(EnvVariables)
	if err != nil {
		fmt.Println(err)
	}
	return db
}
