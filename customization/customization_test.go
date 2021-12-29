package cusomization

import (
	"errors"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/sachinmahanin/passwordstrength/config"
	"github.com/stretchr/testify/assert"
	webserver "github.com/zhongjie-cai/web-server"
)

func TestPreBootstrap_NoErrors(t *testing.T) {
	// arrange
	var dummyCustomization Customization
	// mock
	createMock(t)

	// expect
	timeParseDurationExpected = 1
	timeParseDuration = func(s string) (time.Duration, error) {
		timeParseDurationCalled++
		return time.Duration(0), nil
	}

	// SUT + act
	var err = dummyCustomization.PreBootstrap()

	// assert
	assert.NoError(t, err)

	// verify
	verifyAll(t)
}

func TestPreBootstrap_TimeParseError(t *testing.T) {
	// arrange
	var dummyCustomization Customization
	var expectedErrorMessage = "Failed to parse the default network timeout setting"
	var dummyError = fmt.Errorf(expectedErrorMessage)
	var dummyAppError = webserver.WrapError(dummyError)
	// mock
	createMock(t)

	// expect
	webserverWrapErrorExpected = 1
	webserverWrapError = func(sourceError error, innerErrors ...error) webserver.AppError {
		webserverWrapErrorCalled++
		assert.Equal(t, 1, len(innerErrors))
		assert.Equal(t, dummyError, innerErrors[0])
		assert.Equal(t, dummyError, sourceError)
		return dummyAppError
	}
	timeParseDurationExpected = 1
	timeParseDuration = func(s string) (time.Duration, error) {
		timeParseDurationCalled++
		return time.Duration(0), dummyError
	}

	// SUT + act
	var err = dummyCustomization.PreBootstrap()

	// assert
	assert.Error(t, err)
	assert.Equal(t, dummyAppError, err)

	// verify
	verifyAll(t)
}

func TestDefaultTimeout_Success(t *testing.T) {
	// arrange
	var dummyCustomization Customization
	var dummyDuration = time.Duration(100)

	// mock
	createMock(t)

	// expect
	timeParseDurationExpected = 1
	timeParseDuration = func(s string) (time.Duration, error) {
		timeParseDurationCalled++
		assert.Equal(t, config.DefaultNetworkTimeout, s)
		return dummyDuration, nil
	}

	// SUT + act
	var timeout = dummyCustomization.DefaultTimeout()

	// assert
	assert.Equal(t, dummyDuration, timeout)

	// verify
	verifyAll(t)
}

func TestLog_Success(t *testing.T) {
	// arrange
	var dummyCustomization Customization
	var dummySessionObject = &dummySession{}
	var dummyCategory = "category"
	var dummySubCategory = "subCategory"
	var dummyDiscription = "description"
	// mock
	createMock(t)

	// expect
	customizeLoggingFuncExpected = 1
	customizeLoggingFunc = func(
		session webserver.Session,
		logType webserver.LogType,
		logLevel webserver.LogLevel,
		category string,
		subcategory string,
		description string,
	) {
		customizeLoggingFuncCalled++
		assert.Equal(t, dummySessionObject, session)
		assert.Equal(t, config.AllowedLogType, logType)
		assert.Equal(t, config.AllowedLogLevel, logLevel)
		assert.Equal(t, dummyCategory, category)
		assert.Equal(t, dummySubCategory, subcategory)
		assert.Equal(t, dummyDiscription, description)
	}

	// SUT + act
	dummyCustomization.Log(dummySessionObject, config.AllowedLogType, config.AllowedLogLevel, dummyCategory, dummySubCategory, dummyDiscription)

	// assert

	// verify
	verifyAll(t)
}
func TestPreAction_Success(t *testing.T) {
	// arrange
	var dummyCustomization Customization
	var dummySessionObject = &dummySession{}

	// mock
	createMock(t)

	// expect
	prepareSessionFuncExcepcted = 1
	prepareSessionFunc = func(session webserver.Session) error {
		prepareSessionFuncCalled++
		assert.Equal(t, dummySessionObject, session)
		return nil
	}

	// SUT + act
	var err = dummyCustomization.PreAction(dummySessionObject)

	// assert
	assert.NoError(t, err)

	// verify
	verifyAll(t)
}

func TestPreAction_Error(t *testing.T) {
	// arrange
	var dummyCustomization Customization
	var dummySessionObject = &dummySession{}
	var dummyError = errors.New("some error")
	// mock
	createMock(t)

	// expect
	prepareSessionFuncExcepcted = 1
	prepareSessionFunc = func(session webserver.Session) error {
		assert.Equal(t, dummySessionObject, session)
		prepareSessionFuncCalled++
		return dummyError
	}

	// SUT + act
	var err = dummyCustomization.PreAction(dummySessionObject)

	// assert
	assert.Error(t, err)
	assert.Equal(t, dummyError, err)

	// verify
	verifyAll(t)
}

func TestRoutes_Success(t *testing.T) {
	// arrange
	var dummyCustomization Customization
	var dummyWebserverRoute = webserver.Route{
		Endpoint: "Endpoint",
		Method:   http.MethodGet,
		Path:     "Path",
	}
	var dummyWebserverRoutes = []webserver.Route{
		dummyWebserverRoute,
	}

	// mock
	createMock(t)

	// expect
	webRegisteredRoutesExpected = 1
	webRegisteredRoutes = func() []webserver.Route {
		webRegisteredRoutesCalled++
		return dummyWebserverRoutes
	}

	// SUT + act
	var result = dummyCustomization.Routes()

	// assert
	assert.Equal(t, 1, len(result))
	assert.Equal(t, dummyWebserverRoute.Endpoint, result[0].Endpoint)
	assert.Equal(t, dummyWebserverRoute.Method, result[0].Method)
	assert.Equal(t, dummyWebserverRoute.Path, result[0].Path)

	// verify
	verifyAll(t)
}

func TestStatics_Success(t *testing.T) {
	// arrange
	var dummyCustomization Customization
	var dummyWebserverStatic = webserver.Static{
		Name:       "name",
		PathPrefix: "PathPrefix",
	}
	var dummyWebserverStatics = []webserver.Static{
		dummyWebserverStatic,
	}

	// mock
	createMock(t)

	// expect
	webRegisteredStaticsExpected = 1
	webRegisteredStatics = func() []webserver.Static {
		webRegisteredStaticsCalled++
		return dummyWebserverStatics
	}

	// SUT + act
	var result = dummyCustomization.Statics()

	// assert
	assert.Equal(t, 1, len(result))
	assert.Equal(t, dummyWebserverStatic.Name, result[0].Name)
	assert.Equal(t, dummyWebserverStatic.PathPrefix, result[0].PathPrefix)

	// verify
	verifyAll(t)
}
