package config

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
	webserver "github.com/zhongjie-cai/web-server"
)

func TestSetupApplication_NoErrors(t *testing.T) {
	// arrange
	var dummyHostName = "hostname"
	var dummyAppVersion = "1.0.0"
	var dummyAppPort = "8080"
	var dummyDefaultAllowedLogType = webserver.LogType(rand.Int())
	var dummyDefaultAllowedLogLevel = webserver.LogLevel(rand.Int())
	var dummyDefaultAllowedLogTypeContent = "dummyDefaultAllowedLogType"
	var dummyDefaultAllowedLogLevelContent = "dummyDefaultAllowedLogLevel"
	var osGetEnvParameters = []string{
		"HOSTNAME",
		"APP_VERSION",
		"APP_PORT",
		"ALLOWED_LOG_LEVEL",
		"ALLOWED_LOG_TYPE",
	}
	var osGetEnvReturns = []string{
		dummyHostName,
		dummyAppVersion,
		dummyAppPort,
		dummyDefaultAllowedLogLevelContent,
		dummyDefaultAllowedLogTypeContent,
	}

	// mock
	createMock(t)

	// expect
	osGetenvExpected = 5
	osGetenv = func(key string) string {
		osGetenvCalled++
		if osGetenvCalled > osGetenvExpected {
			return ""
		}

		assert.Equal(t, osGetEnvParameters[osGetenvCalled-1], key)
		return osGetEnvReturns[osGetenvCalled-1]
	}
	webserverNewLogTypeExpected = 1
	webserverNewLogType = func(value string) webserver.LogType {
		webserverNewLogTypeCalled++
		assert.Equal(t, dummyDefaultAllowedLogTypeContent, value)
		return dummyDefaultAllowedLogType
	}
	webserverNewLogLevelExpected = 1
	webserverNewLogLevel = func(value string) webserver.LogLevel {
		webserverNewLogLevelCalled++
		assert.Equal(t, dummyDefaultAllowedLogLevelContent, value)
		return dummyDefaultAllowedLogLevel
	}

	// SUT + act
	var err = SetupApplication()

	// assert
	assert.NoError(t, err)
	assert.Equal(t, dummyHostName, osGetEnvReturns[0])
	assert.Equal(t, dummyAppVersion, osGetEnvReturns[1])
	assert.Equal(t, dummyAppPort, osGetEnvReturns[2])
	assert.Equal(t, dummyDefaultAllowedLogLevelContent, osGetEnvReturns[3])
	assert.Equal(t, dummyDefaultAllowedLogTypeContent, osGetEnvReturns[4])

	// verify
	verifyAll(t)
}
