package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	webserver "github.com/zhongjie-cai/web-server"
)

var (
	osGetenvExpected             int
	osGetenvCalled               int
	webserverNewLogTypeExpected  int
	webserverNewLogTypeCalled    int
	webserverNewLogLevelExpected int
	webserverNewLogLevelCalled   int
)

func createMock(t *testing.T) {
	osGetenvExpected = 0
	osGetenvCalled = 0
	osGetenv = func(key string) string {
		osGetenvCalled++
		return ""
	}
	webserverNewLogTypeExpected = 0
	webserverNewLogTypeCalled = 0
	webserverNewLogType = func(value string) webserver.LogType {
		webserverNewLogTypeCalled++
		return 0
	}
	webserverNewLogLevelExpected = 0
	webserverNewLogLevelCalled = 0
	webserverNewLogLevel = func(value string) webserver.LogLevel {
		webserverNewLogLevelCalled++
		return 0
	}
}

func verifyAll(t *testing.T) {
	osGetenv = os.Getenv
	assert.Equal(t, osGetenvExpected, osGetenvCalled, "Unexpected number of calls to method osGetenv")
	webserverNewLogType = webserver.NewLogType
	assert.Equal(t, webserverNewLogTypeExpected, webserverNewLogTypeCalled, "Unexpected number of calls to method webserverNewLogType")
	webserverNewLogLevel = webserver.NewLogLevel
	assert.Equal(t, webserverNewLogLevelExpected, webserverNewLogLevelCalled, "Unexpected number of calls to method webserverNewLogLevel")
}
