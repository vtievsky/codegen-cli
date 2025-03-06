package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
	httpclient "github.com/vtievsky/codegen-cli/cmd/codegen-cli/internal/http-client"
	httpserver "github.com/vtievsky/codegen-cli/cmd/codegen-cli/internal/http-server"
)

const (
	ErrExitCode = 1
)

func main() {
	var exitCode int

	defer func() {
		cli.OsExiter(exitCode)
	}()

	app := cli.NewApp()
	app.Name = "codegen-cli"
	app.Usage = "Codegen CLI"
	app.Description = "Утилита кодогенерации"
	app.Suggest = true

	app.CommandNotFound = func(ctx *cli.Context, command string) {
		fmt.Printf("The command was not found: %s\n", command)

		exitCode = ErrExitCode
	}

	app.Commands = []*cli.Command{
		httpserver.CommandGenerateHttpServer(),
		httpclient.CommandGenerateHttpClient(),
		httpserver.CommandUploadHttpServer(),
	}

	if err := app.Run(os.Args); err != nil {
		exitCode = ErrExitCode

		return
	}
}
