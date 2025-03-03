package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path"

	"github.com/vtievsky/codegen-cli/gen/clienthttp"
)

func main() {
	ctx := context.Background()

	// Клиентское приложение открывает файл спецификации
	data, err := os.ReadFile("../../docs/openapi/swagger.yaml")
	if err != nil {
		log.Fatal("ошибка чтения спецификации по указанному пути")
	}

	cli, err := clienthttp.NewClientWithResponses("http://127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}

	_, err = cli.UploadSpecHttpWithResponse(ctx, clienthttp.UploadSpecHttpRequest{
		Name: "codegen",
		Spec: data,
	})
	if err != nil {
		log.Fatal(err)
	}

	respCli, err := cli.GenerateSpecServerHttpWithResponse(ctx, &clienthttp.GenerateSpecServerHttpParams{
		Name: "codegen",
	})
	if err != nil {
		log.Fatal(err)
	}

	outputDir := fmt.Sprintf("./tmp/%s", "codegen")

	// Клиентское приложение удаляет предыдущую версию файла спецификации
	if err := os.RemoveAll(outputDir); err != nil {
		log.Fatal(err)
	}

	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		log.Fatal(err)
	}

	// Клиентское приложение сохраняет новую версию файла спецификации
	outputFile := path.Join(outputDir, "clienthttp.go")

	err = os.WriteFile(outputFile, respCli.JSON200.Spec, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("exit")
}
