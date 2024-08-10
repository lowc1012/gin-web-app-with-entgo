package db

import (
	"context"
	"database/sql"
	"fmt"

	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/lib/pq"
	"github.com/lowc1012/gin-web-app-with-entgo/internal/config"
	"github.com/lowc1012/gin-web-app-with-entgo/internal/ent"
	"github.com/lowc1012/gin-web-app-with-entgo/internal/ent/migrate"
	"github.com/lowc1012/gin-web-app-with-entgo/internal/log"
	_ "github.com/mattn/go-sqlite3"
	"github.com/xo/dburl"
)

var (
	dbClient *ent.Client
	dbDriver string
	sqlDB    *sql.DB
)

func Init() (err error) {
	var opts []ent.Option

	// set db.client options according to environment
	if config.IsTestEnv() || config.IsDevEnv() {
		opts = append(opts,
			ent.Debug(),
			ent.Log(func(args ...any) {
				log.Infow(fmt.Sprint(args))
			}))
	}

	// initialize the client
	dbClient, err = entConnect(config.Global.DatabaseURL, opts...)
	if err != nil {
		return err
	}

	return
}

func Client() (*ent.Client, error) {
	var err error
	if dbClient == nil {
		err = Init()
	}

	return dbClient, err
}

// MustClient causes panic if the db client isn't created
func MustClient() *ent.Client {
	var err error
	dbClient, err = Client()
	if dbClient == nil || err != nil {
		panic(err)
	}
	return dbClient
}

// AutoMigrate migrates database schema automatically
func AutoMigrate(c *ent.Client) error {
	ctx := context.Background()
	return c.Schema.Create(ctx, migrate.WithDropIndex(true), migrate.WithDropColumn(true))
}

func Ping() error {
	if sqlDB == nil {
		return fmt.Errorf("database instance not found")
	}
	return sqlDB.Ping()
}

// entConnect creates a connection to DB via ent
func entConnect(dbUrl string, opts ...ent.Option) (*ent.Client, error) {
	url, err := dburl.Parse(dbUrl)
	if err != nil {
		return nil, err
	}

	if sqlDB, err = sql.Open(url.Driver, url.DSN); err != nil {
		return nil, err
	}

	dbDriver = url.Driver
	driver := entsql.OpenDB(dbDriver, sqlDB)
	opts = append(opts, ent.Driver(driver))

	log.Infow("Open Database", "env", config.Global.Env, "driver", url.Driver, "host", url.Host, "db_name", url.Path, "dsn", url.DSN)
	return ent.NewClient(opts...), nil
}
