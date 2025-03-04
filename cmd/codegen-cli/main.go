package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path"

	"github.com/vtievsky/codegen-cli/gen/clienthttp"
	"github.com/vtievsky/codegen-cli/internal/config"
	"github.com/vtievsky/golibs/runtime/logger"
	"go.uber.org/zap"
)

func main() {
	cfg := config.New()
	ctx := context.Background()
	logger := logger.CreateZapLogger(cfg.Debug, cfg.Log.EnableStacktrace)

	if len(os.Args) < 2 {
		log.Fatal("nameSpec, outputDir must be specified")
	}

	// // Клиентское приложение открывает файл спецификации
	// data, err := os.ReadFile("../../docs/openapi/swagger.yaml")
	// if err != nil {
	// 	log.Fatal("ошибка чтения спецификации по указанному пути")
	// }

	cli, err := clienthttp.NewClientWithResponses(cfg.CodegenSvc.URL)
	if err != nil {
		logger.Error("failed to create codegen client",
			zap.Error(err),
		)

		return
	}

	// _, err = cli.UploadSpecHttpWithResponse(ctx, clienthttp.UploadSpecHttpRequest{
	// 	Name: "codegen",
	// 	Spec: data,
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }

	nameSpec := os.Args[1]
	outputDir := fmt.Sprintf("%s/tmp/%s", os.Args[2], nameSpec)
	outputFile := path.Join(outputDir, "clienthttp.go")

	respCli, err := cli.GenerateSpecServerHttpWithResponse(ctx, &clienthttp.GenerateSpecServerHttpParams{
		Name: nameSpec,
	})
	if err != nil {
		logger.Error("failed to generate code for client",
			zap.Error(err),
		)

		return
	}

	// Клиентское приложение удаляет предыдущую версию файла спецификации
	if err := os.RemoveAll(outputDir); err != nil {
		logger.Error("failed to remove outputDir",
			zap.String("outputDir", outputDir),
			zap.Error(err),
		)

		return
	}

	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		logger.Error("failed to make outputDir",
			zap.String("outputDir", outputDir),
			zap.Error(err),
		)

		return
	}

	err = os.WriteFile(outputFile, respCli.JSON200.Spec, os.ModePerm)
	if err != nil {
		logger.Error("failed to write outputFile",
			zap.String("outputFile", outputFile),
			zap.Error(err),
		)

		return
	}

	logger.Info("Successfully",
		zap.String("nameSpec", nameSpec),
		zap.String("outputFile", outputFile),
	)
}
