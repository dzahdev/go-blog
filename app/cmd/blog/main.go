package main

import (
	"context"
	"dzrise.ru/internal/app"
	"log"
)

func main() {
	ctx := context.Background()

	blog, err := app.New()
	if err != nil {
		log.Fatalf("failed to init app: %s", err.Error())
	}

	err = blog.Run(ctx)
	if err != nil {
		log.Fatalf("failed to run app: %s", err.Error())
	}
}
