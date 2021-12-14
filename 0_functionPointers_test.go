package main

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/sachinmahanin/passwordStrength/config"
	"github.com/stretchr/testify/assert"
	webserver "github.com/zhongjie-cai/web-server"
)

var (
	configSetupApplicationExpected  int
	configSetupApplicationCalled    int
	strconvAtoiExpected             int
	strconvAtoiCalled               int
	fmtErrorfExpected               int
	fmtErrorfCalled                 int
	webserverNewApplicationExpected int
	webserverNewApplicationCalled   int
	applicationStopExpected         int
	applicationStopCalled           int
	applicationStartExpected        int
	applicationStartCalled          int
	panicExpected                   int
	panicCalled                     int
)

func createMock(t *testing.T) {

	configSetupApplicationExpected = 0
	configSetupApplicationCalled = 0
	panicExpected = 0
	panicCalled = 0
	configSetupApplication = func() error {
		configSetupApplicationCalled++
		return nil
	}
	strconvAtoiExpected = 0
	strconvAtoiCalled = 0
	strconvAtoi = func(s string) (int, error) {
		strconvAtoiCalled++
		return 0, nil
	}
	fmtErrorfExpected = 0
	fmtErrorfCalled = 0
	fmtErrorf = func(format string, a ...interface{}) error {
		fmtErrorfCalled++
		return nil
	}
	webserverNewApplicationExpected = 0
	webserverNewApplicationCalled = 0
	webserverNewApplication = func(
		name string,
		port int,
		version string,
		customization webserver.Customization,
	) webserver.Application {
		webserverNewApplicationCalled++
		return nil
	}
	applicationStopExpected = 0
	applicationStopCalled = 0
	applicationStartExpected = 0
	applicationStartCalled = 0
}

func verifyAll(t *testing.T) {
	configSetupApplication = config.SetupApplication
	assert.Equal(t, configSetupApplicationExpected, configSetupApplicationCalled, "Unexpected number of calls to method configSetupApplication")
	strconvAtoi = strconv.Atoi
	assert.Equal(t, strconvAtoiExpected, strconvAtoiCalled, "Unexpected number of calls to method strconvAtoi")
	fmtErrorf = fmt.Errorf
	assert.Equal(t, fmtErrorfExpected, fmtErrorfCalled, "Unexpected number of calls to method fmtErrorf")
	webserverNewApplication = webserver.NewApplication
	assert.Equal(t, webserverNewApplicationExpected, webserverNewApplicationCalled, "Unexpected number of calls to method webserverNewApplication")
	assert.Equal(t, applicationStartExpected, applicationStartCalled, "Unexpected number of calls to method applicationStart")
	assert.Equal(t, applicationStopExpected, applicationStopCalled, "Unexpected number of calls to method applicationStop")
	assert.Equal(t, panicExpected, panicCalled, "Unexpected number of calls to method panic")
}

// mock struct
type dummyApplication struct {
	t           *testing.T
	testSession webserver.Session
	isRunning   bool
	startfunc   func()
	stopFunc    func()
}

func (app *dummyApplication) Start() {
	if app.startfunc == nil {
		assert.Fail(app.t, "Unexpected call to Start")
		return
	}
	app.startfunc()
}

func (app *dummyApplication) Stop() {
	if app.stopFunc == nil {
		assert.Fail(app.t, "Unexpected call to Stop")
		return
	}
	app.stopFunc()
}

func (app *dummyApplication) IsRunning() bool {
	return true
}

func (app *dummyApplication) Session() webserver.Session {
	if app.testSession == nil {
		assert.Fail(app.t, "Unexpected call to Session")
		return app.testSession
	}
	return nil
}
