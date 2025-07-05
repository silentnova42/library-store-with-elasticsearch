package app

import (
	"context"
	"os"

	"github.com/silentnova42/library-store-with-elasticsearch/internal/api"
	"github.com/silentnova42/library-store-with-elasticsearch/internal/config"
	storage "github.com/silentnova42/library-store-with-elasticsearch/internal/db"
	"github.com/silentnova42/library-store-with-elasticsearch/internal/server"
	"github.com/silentnova42/library-store-with-elasticsearch/pkg/auth"
	"github.com/silentnova42/library-store-with-elasticsearch/pkg/hash"
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

	hash := new(hash.Hash)
	auth := new(auth.Auth)
	handler := api.NewHandler(database, hash, auth)
	server := server.NewServer()

	return server.RunServer(os.Getenv("PORT_APP"), handler.InitRouter())
}
