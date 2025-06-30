package app

import (
	"context"
	"os"

	"github.com/silentnova42/library-store-with-elasticsearch/internal/api"
	"github.com/silentnova42/library-store-with-elasticsearch/internal/config"
	"github.com/silentnova42/library-store-with-elasticsearch/internal/db/storage"
	"github.com/silentnova42/library-store-with-elasticsearch/internal/server"
)

func Run() error {
	config := config.NewConfigDb()
	if err := config.GetConfigDbFromYaml("configs", "configs"); err != nil {
		return err
	}

	ctx := context.Background()
	database, err := storage.NewDb(ctx, config, 5)
	if err != nil {
		return err
	}

	handler := api.NewHandler(database)
	server := server.NewServer()

	return server.RunServer(os.Getenv("PORT_APP"), handler.InitRouter())
}
