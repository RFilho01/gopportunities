package main

import (
	"fmt"

	"github.com/RFilho01/gopportunities/config"
	"github.com/RFilho01/gopportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	// Initialize Configs
	err := config.Init()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Initialize Router
	router.Initialize()
}
