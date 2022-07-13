package main 

import (
	"https://github.com/oleksandr-kaledin/go-101"
	"https://github.com/oleksandr-kaledin/go-101/tree/main/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(go.Server)

	if err := srv.Run(port: 8000, handlers.initRoutes); err != nil {
		log.Fatalf(format: "error occured while running http server: %s", err.Error())
	}
}
