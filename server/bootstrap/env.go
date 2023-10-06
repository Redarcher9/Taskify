package bootstrap

import (
	"fmt"

	"github.com/joho/godotenv"
)

/*
Description: Loads the environment variables in the server

Params: NA

Returns: Returns the environment variables as a string map
*/
func NewEnv() *map[string]string {
	EnvFile := "../../.env.dev"
	Env, err := godotenv.Read(EnvFile)
	if err != nil {
		fmt.Println(err)
	}
	return &Env
}
