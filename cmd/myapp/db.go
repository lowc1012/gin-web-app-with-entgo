package main

import (
	"github.com/lowc1012/gin-web-app-with-entgo/internal/persistence"
	"github.com/lowc1012/gin-web-app-with-entgo/internal/log"
	"github.com/urfave/cli/v2"
)

var DBMigrate = &cli.Command{
	Name:   "db:migrate",
	Usage:  "Execute auto-migrate database",
	Action: dbMigrateRun,
}

var DBReset = &cli.Command{
	Name:   "db:reset",
	Usage:  "Reset all database data",
	Action: dbResetRun,
}

func dbResetRun(ctx *cli.Context) error {
	if err := initDB(); err != nil {
		return err
	}

	log.Info("Reset and migrate database")
	err := persistence.ResetTables(persistence.MustClient())
	if err != nil {
		log.Fatalw("Database reset and migrate failed",
			"error", err.Error(),
		)
		return err
	}

	return nil
}

func dbMigrateRun(*cli.Context) error {
	if err := initDB(); err != nil {
		log.Errorw("Database initialization failed", "error", err.Error())
		return err
	}
	if err := dbMigrate(); err != nil {
		log.Errorw("Database migration failed", "error", err.Error())
		return err
	}
	return nil
}

func dbMigrate() error {
	log.Infow("Migrating database")
	err := persistence.AutoMigrate(persistence.MustClient())
	if err != nil {
		log.Fatalw("Database migration failed",
			"error", err.Error(),
		)
		return err
	}

	log.Infow("Database migration completed successfully")
	return nil
}

func initDB() error {
	log.Infow("Initialize the database")
	if err := persistence.Init(); err != nil {
		log.Fatalw("Database initialization failed",
			"error", err.Error(),
		)
		return err
	}

	log.Infow("Checking the connection to database")
	if err := persistence.Ping(); err != nil {
		log.Fatalw("Database unhealthy", "error", err.Error())
		return err
	}
	return nil
}
