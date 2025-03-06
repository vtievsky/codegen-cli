package httpserver

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"path"

	codegenhttpclient "github.com/vtievsky/codegen-cli/gen/httpclient/codegen"
	"github.com/vtievsky/codegen-cli/internal/conf"
	"github.com/vtievsky/codegen-cli/pkg/shortcut"
	"github.com/vtievsky/golibs/runtime/logger"
	"go.uber.org/zap"
)

func GenerateHTTPServer(
	ctx context.Context,
	service string,
	packagePath string,
) error {
	var (
		err     error
		workDir string
	)

	if packagePath != "" {
		workDir = packagePath
	} else {
		workDir, err = os.Getwd()
		if err != nil {
			return fmt.Errorf("failed to get workdir")
		}
	}

	outputDir := path.Join(workDir, shortcut.OutputDirHttpServer)
	outputDir = path.Join(outputDir, service)
	outputFile := fmt.Sprintf("%shttpserver.go", service)
	outputFileName := path.Join(outputDir, outputFile)

	conf := conf.New()
	logger := logger.CreateZapLogger(conf.Debug, conf.Log.EnableStacktrace)

	cli, err := codegenhttpclient.NewClientWithResponses(conf.CodegenSvc.URL)
	if err != nil {
		logger.Error("failed to create codegen client",
			zap.Error(err),
		)

		return fmt.Errorf("failed to create codegen client | %w", err)
	}

	respCli, err := cli.GenerateSpecServerHttpWithResponse(ctx, service)
	if err != nil {
		logger.Error("failed to generate code for httpserver",
			zap.String("service", service),
			zap.Error(err),
		)

		return fmt.Errorf("failed to generate code for httpserver | %w", err)
	}

	if respCli.HTTPResponse.StatusCode != http.StatusOK {
		logger.Error("failed to generate code for httpserver",
			zap.String("service", service),
			zap.String("status", respCli.HTTPResponse.Status),
			zap.String("description", respCli.JSON500.Status.Description),
		)

		return fmt.Errorf("failed to generate code for httpserver | status: %s", respCli.HTTPResponse.Status)
	}

	if err := os.RemoveAll(outputDir); err != nil {
		logger.Error("failed to remove outputDir",
			zap.String("outputDir", outputDir),
			zap.Error(err),
		)

		return fmt.Errorf("failed to remove outputDir | %w", err)
	}

	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		logger.Error("failed to make outputDir",
			zap.String("outputDir", outputDir),
			zap.Error(err),
		)

		return fmt.Errorf("failed to make outputDir | %w", err)
	}

	err = os.WriteFile(outputFileName, respCli.JSON200.Spec, os.ModePerm)
	if err != nil {
		logger.Error("failed to write outputFileName",
			zap.String("outputFileName", outputFileName),
			zap.Error(err),
		)

		return fmt.Errorf("failed to write outputFileName | %w", err)
	}

	return nil
}
