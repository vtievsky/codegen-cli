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

	app.Commands = []*cli.Command{
		httpserver.CommandGenerateHttpServer(),
		httpclient.CommandGenerateHttpClient(),
		httpserver.CommandUploadHttpServer(),
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
