package httpclient

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func Command() *cli.Command {
	return &cli.Command{
		Name:  "gen-http-client",
		Usage: "Генерация HTTP клиента ",
		Description: "Генерация HTTP клиента для сервиса (--service/-s) на основании спецификации из Spec Storage \n" +
			"Результат кодогенерации сохраняется в директории - \"gen/httpclient/{service_name}\" \n" +
			"Название пакета - \"{service_name}httpclient\"  \n",
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
			if err := GenHTTPClient(
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
