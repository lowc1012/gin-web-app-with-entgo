package main

import (
	"fmt"
	"os"

	"github.com/lowc1012/gin-web-app-with-entgo/internal/log"
	"github.com/urfave/cli/v2"
)

func main() {
	root := &cli.App{
		Name: "myapp",
		Action: func(c *cli.Context) error {
			fmt.Printf("This is my first golang app")
			return nil
		},
		Commands: []*cli.Command{
			Start,
			DBMigrate,
			DBReset,
		},
	}

	if err := root.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
