package main

import (
	"echo-mangosteen/cmd/wire"
	"echo-mangosteen/pkg/http"
	"log"
)

func main() {
	app, cleanup, err := wire.NewApp()

	if err != nil {
		log.Panic(err)
	}

	http.Run(app, ":8000")

	defer cleanup()
}
