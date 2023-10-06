package bootstrap

import (
	"database/sql"
	"fmt"
	"taskify/database"
)

func SetupDB() *sql.DB {
	db, err := database.NewPsqlDB()
	if err != nil {
		fmt.Println(err)
	}
	return db
}
