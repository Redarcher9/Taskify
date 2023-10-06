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
func BuildConnString(EnvVariables map[string]string) string {
	connString := "postgres://" + EnvVariables["DATABASE_USERNAME"] + ":" + EnvVariables["DATABASE_PASSWORD"] + "@" + EnvVariables["DATABASE_ADDRESS"] + ":" + EnvVariables["DATABASE_PORT"] + "/" + EnvVariables["DATABASE_NAME"] + "?sslmode=" + EnvVariables["DATABASE_SSLMODE"]
	return connString
}

/*
Description: Establishes a connection to PostGreSQL database

params: EnvVariables map from env file

returns: PSQLDatabase struct with instance of postgresql server,

	error: nil - database connection is successful
			err - if any
*/
func NewPsqlDB(EnvVariables map[string]string) (*sql.DB, error) {
	connString := BuildConnString(EnvVariables)
	db, err := sql.Open("postgres", connString)

	return db, err
}
