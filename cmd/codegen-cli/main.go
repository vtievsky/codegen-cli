package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
	httpclient "github.com/vtievsky/codegen-cli/cmd/codegen-cli/internal/http-client"
	httpserver "github.com/vtievsky/codegen-cli/cmd/codegen-cli/internal/http-server"
)

func main() {
	app := cli.NewApp()
	app.DefaultCommand = "command-list"

	app.Commands = []*cli.Command{{
		Name:         "codegen",
		Aliases:      nil,
		Usage:        "Утилита для кодогенерации",
		UsageText:    "",
		Description:  "",
		ArgsUsage:    "",
		Category:     "",
		BashComplete: cli.DefaultAppComplete,
		Before:       nil,
		After:        nil,
		Action:       nil,
		OnUsageError: nil,
		Subcommands: []*cli.Command{
			httpclient.Command(),
			httpserver.Command(),
		},
		Flags:                  nil,
		SkipFlagParsing:        false,
		HideHelp:               false,
		HideHelpCommand:        false,
		Hidden:                 false,
		UseShortOptionHandling: false,
		HelpName:               "",
		CustomHelpTemplate:     "",
	}}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
