package main

import (
	"context"
	"log"

	"github.com/sigstore/cosign/v2/pkg/providers"
	_ "github.com/sigstore/cosign/v2/pkg/providers/github"
)

func main() {
	ctx := context.Background()
	if !providers.Enabled(ctx) {
		log.Fatalf("incorrect environment")
	}
	token, err := providers.Provide(ctx, "guac")
	if err != nil {
		log.Fatalf(err.Error())
	}
	if token == "" {
		log.Fatalf("empty token")
	}
	log.Printf("Token aquired")
}
