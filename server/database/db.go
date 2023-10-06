package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

/*
Description: Builds connection string for Postgresql database

params: EnvVariables map from env file

returns: Connection string
*/
func BuildConnString() string {
	username := "postgres"
	password := "postgres"
	address := "localhost"
	port := "5432"
	database_name := "Taskify"
	sslmode := "false"
	connString := "postgres://" + username + ":" + password + "@" + address + ":" + port + "/" + database_name + "?sslmode=" + sslmode
	return connString
}

/*
Description: Establishes a connection to PostGreSQL database

params: EnvVariables map from env file

returns: PSQLDatabase struct with instance of postgresql server,

	error: nil - database connection is successful
			err - if any
*/
func NewPsqlDB() (*sql.DB, error) {
	connString := BuildConnString()
	db, err := sql.Open("postgres", connString)

	return db, err
}
