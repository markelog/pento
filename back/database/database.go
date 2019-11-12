package database

import (
	"os"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/markelog/pento/back/database/models"
	"github.com/markelog/pento/back/env"
	"github.com/markelog/pento/back/logger"
	"github.com/qor/validations"
)

const enableLogs = false

// Up database
func Up() *gorm.DB {
	var (
		log = logger.Up()
		dsn = os.Getenv("DATABASE_URL")

		err error
		db  *gorm.DB
	)

	if len(dsn) != 0 {
		db, err = models.ConnectDSN(dsn)
	} else {
		db, err = models.Connect(
			&models.ConnectArgs{
				User:     os.Getenv("DATABASE_USER"),
				Password: os.Getenv("DATABASE_PASSWORD"),
				Name:     os.Getenv("DATABASE_NAME"),
				Host:     os.Getenv("DATABASE_HOST"),
				Port:     os.Getenv("DATABASE_PORT"),
				SSL:      os.Getenv("DATABASE_SSL"),
			},
		)
	}
	if err != nil {
		log.Panic(err)
	}

	// Logs?
	db.LogMode(enableLogs)

	// Flush out dead connections
	db.DB().SetConnMaxLifetime(time.Second * 500)

	// Plugins
	validations.RegisterCallbacks(db)

	// Migrations
	err = db.AutoMigrate(
		&models.User{},
		&models.Track{},
	).Error

	if err != nil {
		log.Panic(err)
	}

	// Foreign keys
	db.Model(
		&models.Track{},
	).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")

	fixtures := os.Getenv("DATABASE_FIXTURES_PATH")
	if len(fixtures) > 0 {
		_, err = env.Fixtures(fixtures, db)
		if err != nil {
			log.Panic(err)
		}
	}

	return db
}
