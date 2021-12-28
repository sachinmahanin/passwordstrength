package cusomization

import (
	"fmt"
	"time"

	"github.com/sachinmahanin/passwordStrength/config"
	webserver "github.com/zhongjie-cai/web-server"
)

// Customization inherits from the default customization so you can skip setting up all customization methods
//   alternatively, you could bring in your own struct that instantiate the webserver.Customization interface to have a verbosed control over what to customize
type Customization struct {
	webserver.DefaultCustomization
}

// PreBootstrap is to customize the pre-processing logic before bootstrapping
func (customization *Customization) PreBootstrap() error {
	var _, parseError = timeParseDuration(config.DefaultNetworkTimeout)
	if parseError != nil {
		return webserverWrapError(fmt.Errorf("Failed to parse the default network timeout setting"), parseError)
	}
	return nil
}

// DefaultTimeout is to customize the default timeout for any webcall communications through HTTP/HTTPS by session
func (customization *Customization) DefaultTimeout() time.Duration {
	var defaultNetworkTimeout, _ = timeParseDuration(config.DefaultNetworkTimeout)
	return defaultNetworkTimeout
}

// Log is to customize the logging backend for the whole application
func (customization *Customization) Log(session webserver.Session, logType webserver.LogType, logLevel webserver.LogLevel, category, subcategory, description string) {
	customizeLoggingFunc(session, logType, logLevel, category, subcategory, description)
}

// PreAction is to customize the pre-action used before each route action takes place, e.g. authorization, etc.
func (customization *Customization) PreAction(session webserver.Session) error {
	return prepareSessionFunc(session)
}

func (customization *Customization) Routes() []webserver.Route {
	return webRegisteredRoutes()
}

func (customization *Customization) Statics() []webserver.Static {
	return webRegisteredStatics()
}
