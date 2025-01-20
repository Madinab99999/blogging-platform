package main

import (
	"context"
	"fmt"

	"github.com/Madinab99999/blogging-platform/config"
)

func main() {
	ctx, _ := context.WithCancel(context.Background())

	// config
	conf := config.New(ctx)
	fmt.Println("New Server")
	fmt.Println(conf.API.Rest.Port)
}
