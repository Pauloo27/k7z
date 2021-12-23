package main

import (
	"github.com/Pauloo27/k7z/internal/config"
	"github.com/Pauloo27/k7z/internal/server"
)

func main() {
	err := config.LoadConfig()
	if err != nil {
		panic(err)
	}
	err = server.StartHTTPServer()
	if err != nil {
		panic(err)
	}
}
