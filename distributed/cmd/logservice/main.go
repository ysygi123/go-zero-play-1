package main

import (
	"context"
	"fmt"
	"go-zero-play-1/distributed/log"
	"go-zero-play-1/distributed/registry"
	"go-zero-play-1/distributed/service"
	stlog "log"
)

func main() {
	log.Run("./distributed.log")
	host, port := "localhost", "4000"
	serviceAddress := fmt.Sprintf("http://%s:%s", host, port)
	r := registry.Registration{
		ServiceName: "LogService",
		ServiceURL:  serviceAddress,
	}
	ctx, err := service.Start(context.Background(), host, port, r, log.RegisterHandler)
	if err != nil {
		stlog.Fatalln(err)
	}
	<-ctx.Done()
	fmt.Println("over")
}
