package main

import (
	"log"

	"github.com/silentnova42/library-store-with-elasticsearch/internal/app"
)

func main() {
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
