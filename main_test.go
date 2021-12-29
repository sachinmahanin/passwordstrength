package main

import (
	"errors"
	"testing"

	"github.com/sachinmahanin/passwordstrength/config"
	customization "github.com/sachinmahanin/passwordstrength/customization"
	"github.com/stretchr/testify/assert"
	webserver "github.com/zhongjie-cai/web-server"
)

func TestMain_HappyPath(t *testing.T) {
	// arrange
	var dummyApp = &dummyApplication{}
	var dummyPort = 80
	var dummyCustomization = &customization.Customization{}
	// mock
	createMock(t)

	// expect
	configSetupApplicationExpected = 1
	configSetupApplication = func() error {
		configSetupApplicationCalled++
		return nil
	}
	strconvAtoiExpected = 1
	strconvAtoi = func(s string) (int, error) {
		strconvAtoiCalled++
		assert.Equal(t, config.AppPort, s)
		return dummyPort, nil
	}
	webserverNewApplicationExpected = 1
	webserverNewApplication = func(
		name string,
		port int,
		version string,
		customization webserver.Customization,
	) webserver.Application {
		webserverNewApplicationCalled++
		assert.Equal(t, config.AppName, name)
		assert.Equal(t, dummyPort, port)
		assert.Equal(t, config.AppVersion, version)
		assert.Equal(t, dummyCustomization, customization)
		return dummyApp
	}
	applicationStartExpected = 1
	dummyApp.startfunc = func() {
		applicationStartCalled++
		return
	}
	applicationStopExpected = 1
	dummyApp.stopFunc = func() {
		applicationStopCalled++
		return
	}

	// SUT + act
	main()

	// assert

	// verify
	verifyAll(t)
}

func TestMain_configSetupApplicationFail(t *testing.T) {
	// arrange
	var dummyError = errors.New("app setup error")
	// mock
	createMock(t)

	// expect
	configSetupApplicationExpected = 1
	configSetupApplication = func() error {
		configSetupApplicationCalled++
		return dummyError
	}
	panicExpected = 1
	defer func() {
		panicCalled++

		// verify
		verifyAll(t)
		recover()
	}()

	// SUT + act
	main()

	// assert

	// verify
}

func TestMain_StrconvAtoiError(t *testing.T) {
	// arrange
	var dummyError = errors.New("app setup error")
	var dummyFmtError = errors.New("some error")
	// mock
	createMock(t)

	// expect
	configSetupApplicationExpected = 1
	configSetupApplication = func() error {
		configSetupApplicationCalled++
		return nil
	}
	strconvAtoiExpected = 1
	strconvAtoi = func(s string) (int, error) {
		strconvAtoiCalled++
		assert.Equal(t, config.AppPort, s)
		return 0, dummyError
	}
	fmtErrorfExpected = 1
	fmtErrorf = func(format string, a ...interface{}) error {
		assert.Equal(t, "Invalid port number provided: %v", format)
		assert.Equal(t, 1, len(a))
		assert.Equal(t, config.AppPort, a[0])
		fmtErrorfCalled++
		return dummyFmtError
	}
	panicExpected = 1
	defer func() {
		panicCalled++

		// verify
		verifyAll(t)
		recover()
	}()

	// SUT + act
	main()

	// assert

	// verify
}
