package main

import (
	"github.com/CaboLira/gooportunities.git/config"
	"github.com/CaboLira/gooportunities.git/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	// Initialize Configs
	err := config.Init()
	if err != nil {
		logger.Errorf("config inicialization error: %v", err)
		return
	}

	// Initialize Router
	router.Inicialize()
}
