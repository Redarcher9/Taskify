package bootstrap

import (
	"database/sql"
	"fmt"
	"taskify/database"
)

/*
Description: Bootstrap function to setup database

params: Env variables

returns: sql database instance
*/

func SetupDB(EnvVariables map[string]string) *sql.DB {
	db, err := database.NewPsqlDB(EnvVariables)
	if err != nil {
		fmt.Println(err)
	}
	return db
}
