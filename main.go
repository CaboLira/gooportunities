package main

import (
	"fmt"

	"github.com/CaboLira/gooportunities.git/config"
	"github.com/CaboLira/gooportunities.git/router"
)

func main() {

	// Initialize Configs
	err := config.Init()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Initialize Router
	router.Inicialize()
}
