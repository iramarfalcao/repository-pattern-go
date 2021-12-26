package main

import (
	"github.com/iramarfalcao/repository-pattern-go/config"
	"github.com/iramarfalcao/repository-pattern-go/database"
	"github.com/iramarfalcao/repository-pattern-go/server"
)

func main() {
	config.Init()

	database.StartDatabase()
	s := server.NewServer()

	s.Run()
}
