package main

import (
	"fmt"
	"log"
	"os"

	"github.com/faraazkhan/gauth/pkg/totp"
	"github.com/urfave/cli/v2"
)

func main() {
	configFile := os.Getenv("GAUTH_CONFIG_FILE")
	app := &cli.App{
		Name:  "token",
		Usage: "generate TOTP using configuration file for alfred-workflow-gauth. Example: gauth token my-token",
		Action: func(c *cli.Context) error {
			accountName := c.Args().Get(1)
			if accountName == "" {
				return fmt.Errorf("You must specify a valid account name. Not: %s", accountName)
			}
			token, err := (totp.Token(accountName, configFile))
			fmt.Println(token)
			return err
		},
	}
	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
