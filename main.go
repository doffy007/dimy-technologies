package main

import (
	"context"

	"github.com/doffy007/dimy-technologies/config"
	"github.com/doffy007/dimy-technologies/server"
)

func main() {
	config.Init()

	ctx := context.Background()
	server.NewApp(ctx).Start()
}
