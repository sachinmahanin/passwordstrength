package main

import (
	"github.com/sachinmahanin/passwordstrength/config"
	customization "github.com/sachinmahanin/passwordstrength/customization"
)

// This is a sample of how to setup application for running the server
func main() {

	var configError = configSetupApplication()
	if configError != nil {
		panic(
			configError,
		)
	}

	var port, convErr = strconvAtoi(config.AppPort)
	if convErr != nil {
		panic(
			fmtErrorf(
				"Invalid port number provided: %v",
				config.AppPort,
			),
		)
	}
	var application = webserverNewApplication(
		config.AppName,
		port,
		config.AppVersion,
		&customization.Customization{},
	)
	defer application.Stop()
	application.Start()
}
