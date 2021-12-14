package config

import (
	"os"

	webserver "github.com/zhongjie-cai/web-server"
)

// func pointers for injection / testing: config.go
var (
	osGetenv             = os.Getenv
	webserverNewLogType  = webserver.NewLogType
	webserverNewLogLevel = webserver.NewLogLevel
)
