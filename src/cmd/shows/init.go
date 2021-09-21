package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"

	"shows/src/api"
	"shows/src/repository"
)

var (
	DBAddr = "localhost:5432"
	DBUser = "user"
	DBPass = "pass"
	DBName = "shows-db"
)

func init() {
	if addr := os.Getenv("DB_ADDR"); addr != "" {
		DBAddr = addr
	}
	if name := os.Getenv("DB_NAME"); name != "" {
		DBName = name
	}
	if user := os.Getenv("DB_USER"); user != "" {
		DBUser = user
	}
	if pass := os.Getenv("DB_PASS"); pass != "" {
		DBPass = pass
	}
}

func register(db *sqlx.DB) *gin.Engine {
	router := gin.Default()
	sR := repository.NewShowRepo(db)
	pR := repository.NewPersonRepo(db)

	pH := api.NewPersonHandler(pR, sR)
	sH := api.NewShowHandler(sR, pR)

	sH.Register(router.Group("show"))
	pH.Register(router.Group("person"))
	return router
}
