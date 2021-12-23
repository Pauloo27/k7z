package main

import (
	"fmt"

	"github.com/Pauloo27/k7z/internal/config"
)

func main() {
	fmt.Println("hello")
	err := config.LoadConfig()
	if err != nil {
		panic(err)
	}
}
