package httpserver

import (
	"fmt"

	"github.com/urfave/cli/v2"
	"github.com/vtievsky/codegen-cli/pkg/shortcut"
)

func CommandGenerateHttpServer() *cli.Command {
	return &cli.Command{
		Name:  "gen-http-server",
		Usage: "Генерация HTTP-сервера",
		Description: "Генерация HTTP-сервера для сервиса (--service/-s) на основании спецификации из Spec Storage \n" +
			fmt.Sprintf("Результат кодогенерации сохраняется в директории - \"%s/{service_name}\" \n", shortcut.OutputDirHttpServer) +
			"Название пакета - \"{service_name}httpserver\"  \n",
		Flags: []cli.Flag{
			&cli.StringFlag{ //nolint:exhaustruct
				Name:     "service",
				Aliases:  []string{"s"},
				Required: true,
				Usage:    "Название сервиса",
				Value:    "",
			},
			&cli.StringFlag{ //nolint:exhaustruct
				Name:    "packagepath",
				Aliases: []string{"p"},
				Usage:   "Путь размещения пакета",
				Value:   "",
			},
		},
		Action: func(cCtx *cli.Context) error {
			if err := GenerateHTTPServer(
				cCtx.Context,
				cCtx.String("service"),
				cCtx.String("packagepath"),
			); err != nil {
				return fmt.Errorf("failed to exec command | %w", err)
			}

			return nil
		},
	}
}
