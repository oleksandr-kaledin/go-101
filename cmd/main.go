package main 

import (
	"github.com/oleksandr-kaledin/go-101"
	"github.com/oleksandr-kaledin/go-101/tree/main/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(goss.Server)
	if err := srv.Run(8000, handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
