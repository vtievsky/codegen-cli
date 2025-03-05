package httpserver

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func CommandUploadHttpServer() *cli.Command {
	return &cli.Command{
		Name:        "upload-http-server",
		Usage:       "Выгрузка спецификации HTTP-сервера",
		Description: "Выгрузка спецификации HTTP-сервера для сервиса (--service/-s) на основании спецификации из Spec Storage \n",
		Flags: []cli.Flag{
			&cli.StringFlag{ //nolint:exhaustruct
				Name:     "service",
				Aliases:  []string{"s"},
				Required: true,
				Usage:    "Название сервиса",
				Value:    "",
			},
			&cli.StringFlag{ //nolint:exhaustruct
				Name:     "source",
				Aliases:  []string{"f"},
				Required: true,
				Usage:    "Файл спецификации",
				Value:    "",
			},
		},
		Action: func(cCtx *cli.Context) error {
			if err := UploadHTTPServerSpec(
				cCtx.Context,
				cCtx.String("service"),
				cCtx.String("source"),
			); err != nil {
				return fmt.Errorf("failed to exec command | %w", err)
			}

			return nil
		},
	}
}
