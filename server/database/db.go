package database

import (
	"database/sql"
)

/*
Description: Struct for PostGreSQL database
*/
type PsqlDatabase struct {
	dbInstance *sql.DB
}

/*
Description: Builds connection string for Postgresql database

params: EnvVariables map from env file

returns: Connection string
*/
func BuildConnString(EnvVariables map[string]string) string {
	connString := "postgres://"
	return connString
}

/*
Description: Establishes a connection to PostGreSQL database

params: EnvVariables map from env file

returns: PSQLDatabase struct with instance of postgresql server,

	error: nil - database connection is successful
			err - if any
*/
func NewPsqlDB(EnvVariables map[string]string) (*PsqlDatabase, error) {
	connString := BuildConnString(EnvVariables)
	db, err := sql.Open("postgres", connString)

	return &PsqlDatabase{dbInstance: db}, err
}
