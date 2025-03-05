package httpserver

import (
	"context"
	"fmt"
	"net/http"
	"os"

	codegenhttpclient "github.com/vtievsky/codegen-cli/gen/httpclient/codegen"
	"github.com/vtievsky/codegen-cli/internal/conf"
	"github.com/vtievsky/golibs/runtime/logger"
	"go.uber.org/zap"
)

func UploadHTTPServerSpec(
	ctx context.Context,
	service string,
	source string,
) error {
	conf := conf.New()
	logger := logger.CreateZapLogger(conf.Debug, conf.Log.EnableStacktrace)

	cli, err := codegenhttpclient.NewClientWithResponses(conf.CodegenSvc.URL)
	if err != nil {
		logger.Error("failed to create codegen client",
			zap.Error(err),
		)

		return fmt.Errorf("failed to create codegen client | %w", err)
	}

	data, err := os.ReadFile(source)
	if err != nil {
		return fmt.Errorf("failed to read source for httpserver  | %w", err)
	}

	respCli, err := cli.UploadSpecHttpWithResponse(ctx, codegenhttpclient.UploadSpecHttpRequest{
		Name: service,
		Spec: data,
	})
	if err != nil {
		logger.Error("failed to upload spec for httpserver",
			zap.String("service", service),
			zap.Error(err),
		)

		return fmt.Errorf("failed to upload spec for httpserver | %w", err)
	}

	if respCli.HTTPResponse.StatusCode != http.StatusOK {
		logger.Error("failed to upload spec for httpserver",
			zap.String("service", service),
			zap.String("status", respCli.HTTPResponse.Status),
			zap.String("description", respCli.JSON500.Status.Description),
		)

		return fmt.Errorf("failed to upload spec for httpserver | status: %s", respCli.HTTPResponse.Status)
	}

	return nil
}
