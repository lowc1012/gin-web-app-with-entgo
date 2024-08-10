package main

import (
	"github.com/lowc1012/gin-web-app-with-entgo/internal/db"
	"github.com/lowc1012/gin-web-app-with-entgo/internal/log"
)

func dbMigrate() error {
	log.Infow("Migrate database")
	err := db.AutoMigrate(db.MustClient())
	if err != nil {
		log.Fatalw("Database migration failed",
			"error", err.Error(),
		)
		return err
	}

	return nil
}

func initDB() error {
	log.Infow("Initialize the database")
	if err := db.Init(); err != nil {
		log.Fatalw("Database initialization failed",
			"error", err.Error(),
		)
		return err
	}

	log.Infow("Checking the connection to database")
	if err := db.Ping(); err != nil {
		log.Fatalw("Database unhealthy", "error", err.Error())
		return err
	}
	return nil
}
