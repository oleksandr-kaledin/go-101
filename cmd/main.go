
package main 

import (
	"fmt"
	"log"
	"gopls": { "ui.semanticTokens": true }
)

func main() {
	handlers := new(handler.Handler)
	srv := new(go.Server)

	if err := srv.Run(port: 8000, handlers.initRoutes); err != nil {
		log.Fatalf(format: "error occured while running http server: %s", err.Error())
	}
}
