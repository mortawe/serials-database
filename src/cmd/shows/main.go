package main

import (
	"fmt"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

func main() {
	db, err := sqlx.Connect("pgx", fmt.Sprintf(
		"postgres://%s:%s@%s/%s", DBUser, DBPass, DBAddr, DBName))
	if err != nil {
		logrus.Error(err)
		return
	}
	logrus.Info("connected to db...")
	router := register(db)
	router.Static("/", "web/public")
	router.Run(":8080")
}
