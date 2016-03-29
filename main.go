package main

import (
	"os"
	"strconv"

	"github.com/lucasvmiguel/GoApi/api"
	"github.com/lucasvmiguel/GoApi/config"
)

func main() {
	env := os.Getenv("env")
	config := config.New("envs")

	if env != "" {
		config.Read(env)
	} else {
		config.Read("development")
	}

	api.Start(":" + strconv.Itoa(config.Api.Port))
}
