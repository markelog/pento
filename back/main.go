package main

import (
	"os"

	"github.com/kataras/iris"
	"github.com/markelog/pento/back/database"
	"github.com/markelog/pento/back/env"
	"github.com/markelog/pento/back/logger"
	"github.com/markelog/pento/back/routes"
	"github.com/markelog/pento/back/routes/common"
	"github.com/markelog/pento/back/routes/track"
	"github.com/sirupsen/logrus"
)

func main() {
	env.Up()

	var (
		port    = os.Getenv("PORT")
		address = ":" + port
	)

	var (
		app = routes.Up()
		db  = database.Up()
		log = logger.Up()
	)

	defer db.Close()

	track.Up(app, db, log)
	common.Up(app, db, log)

	log.WithFields(logrus.Fields{
		"port": port,
	}).Info("Started")
	app.Run(iris.Addr(address))
}
