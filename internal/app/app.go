package app

import (
	"github.com/silentnova42/library-store-with-elasticsearch/internal/api"
	"github.com/silentnova42/library-store-with-elasticsearch/internal/db/storage"
	"github.com/silentnova42/library-store-with-elasticsearch/internal/server"
)

func Run() error {
	server := server.NewServer()
	database := storage.NewDb()
	handler := api.NewHandler(database)

	return server.RunServer(":8080", handler.InitRouter())
}
