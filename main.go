package main

import (
	"github.com/the-medium-tech/mdl-manager/hashvegas-proxy-go/internal/controller"
	"github.com/the-medium-tech/platform-externals/log"
)

func main() {
	log.Info("Starting hashvegas proxy go")
	c := controller.NewHvController()
	err := c.RegisterRoute()
	if err != nil {
		panic("Registering routes failed")
	}
	c.Start("0.0.0.0:8081")

	return
}
