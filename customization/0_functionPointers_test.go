package cusomization

import (
	"net/http"
	"reflect"
	"testing"
	"time"

	"github.com/google/uuid"
	logger "github.com/sachinmahanin/passwordStrength/middleware/logger"
	"github.com/sachinmahanin/passwordStrength/sessionutil"
	web "github.com/sachinmahanin/passwordStrength/web"
	"github.com/stretchr/testify/assert"
	webserver "github.com/zhongjie-cai/web-server"
)

var (
	timeParseDurationExpected    int
	timeParseDurationCalled      int
	customizeLoggingFuncExpected int
	customizeLoggingFuncCalled   int
	prepareSessionFuncExcepcted  int
	prepareSessionFuncCalled     int
	webRegisteredRoutesExpected  int
	webRegisteredRoutesCalled    int
	webRegisteredStaticsExpected int
	webRegisteredStaticsCalled   int
	webserverWrapErrorExpected   int
	webserverWrapErrorCalled     int
)

func createMock(t *testing.T) {

	webRegisteredStaticsExpected = 0
	webRegisteredStaticsCalled = 0
	webRegisteredStatics = func() []webserver.Static {
		webRegisteredStaticsCalled++
		return []webserver.Static{}
	}

	timeParseDurationExpected = 0
	timeParseDurationCalled = 0
	timeParseDuration = func(s string) (time.Duration, error) {
		timeParseDurationCalled++
		return time.Duration(0), nil
	}
	customizeLoggingFuncExpected = 0
	customizeLoggingFuncCalled = 0
	customizeLoggingFunc = func(
		session webserver.Session,
		logType webserver.LogType,
		logLevel webserver.LogLevel,
		category string,
		subcategory string,
		description string,
	) {
		customizeLoggingFuncCalled++

	}
	prepareSessionFuncExcepcted = 0
	prepareSessionFuncCalled = 0
	prepareSessionFunc = func(session webserver.Session) error {
		prepareSessionFuncCalled++
		return nil
	}

	webRegisteredRoutesExpected = 0
	webRegisteredRoutesCalled = 0
	webRegisteredRoutes = func() []webserver.Route {
		webRegisteredRoutesCalled++
		return []webserver.Route{}
	}
	webserverWrapErrorExpected = 0
	webserverWrapErrorCalled = 0
	webserverWrapError = func(sourceError error, innerErrors ...error) webserver.AppError {
		webserverWrapErrorCalled++
		return nil
	}
}

func verifyAll(t *testing.T) {
	webRegisteredStatics = web.RegisteredStatics
	assert.Equal(t, webRegisteredStaticsExpected, webRegisteredStaticsCalled, "Unexpected number of calls to method webRegisteredStatics")
	webRegisteredRoutes = web.RegisteredRoutes
	assert.Equal(t, webRegisteredRoutesExpected, webRegisteredRoutesCalled, "Unexpected number of calls to method webRegisteredRoutes")
	timeParseDuration = time.ParseDuration
	assert.Equal(t, timeParseDurationExpected, timeParseDurationCalled, "Unexpected number of calls to method timeParseDuration")
	customizeLoggingFunc = logger.CustomizeLoggingFunc
	assert.Equal(t, customizeLoggingFuncExpected, customizeLoggingFuncCalled, "Unexpected number of calls to method customizeLoggingFunc")
	prepareSessionFunc = sessionutil.PrepareSession
	assert.Equal(t, prepareSessionFuncExcepcted, prepareSessionFuncCalled, "Unexpected number of calls to method prepareSessionFunc")
	webserverWrapError = webserver.WrapError
	assert.Equal(t, webserverWrapErrorExpected, webserverWrapErrorCalled, "Unexpected number of calls to method webserverWrapError")
}

// mock struct
type dummySession struct {
	t                *testing.T
	id               uuid.UUID
	name             string
	request          *http.Request
	responseWriter   http.ResponseWriter
	attachment       map[string]interface{}
	headerValue      map[string]interface{}
	headerError      map[string]webserver.AppError
	attachSetValue   map[string]interface{}
	attachSetSuccess map[string]bool
	attachGetValue   map[string]interface{}
	attachGetSuccess map[string]bool
}

func (session *dummySession) GetID() uuid.UUID {
	if session.id == uuid.Nil {
		assert.Fail(session.t, "Unexpected call to GetID")
		return uuid.Nil
	}
	return session.id
}

func (session *dummySession) GetName() string {
	if session.name == "" {
		assert.Fail(session.t, "Unexpected call to GetName")
		return ""
	}
	return session.name
}

func (session *dummySession) GetRequest() *http.Request {
	if session.request == nil {
		assert.Fail(session.t, "Unexpected call to GetRequest")
		return nil
	}
	return session.request
}

func (session *dummySession) GetResponseWriter() http.ResponseWriter {
	if session.responseWriter == nil {
		assert.Fail(session.t, "Unexpected call to GetResponseWriter")
		return nil
	}
	return session.responseWriter
}

func (session *dummySession) GetRequestBody(dataTemplate interface{}) error {
	assert.Fail(session.t, "Unexpected call to GetRequestBody")
	return nil
}

func (session *dummySession) GetRequestParameter(name string, dataTemplate interface{}) error {
	assert.Fail(session.t, "Unexpected call to GetRequestParameter")
	return nil
}

func (session *dummySession) GetRequestQuery(name string, index int, dataTemplate interface{}) error {
	assert.Fail(session.t, "Unexpected call to GetRequestQuery")
	return nil
}

func (session *dummySession) GetRequestHeader(name string, index int, dataTemplate interface{}) error {
	if session.headerValue == nil && session.headerError == nil {
		assert.Fail(session.t, "Unexpected call to GetRequestHeader")
		return nil
	}
	var err, errFound = session.headerError[name]
	if errFound && err != nil {
		return err
	}
	var value, valueFound = session.headerValue[name]
	if valueFound {
		if reflect.TypeOf(dataTemplate) == reflect.TypeOf((*string)(nil)) {
			(*(dataTemplate).(*string)) = value.(string)
		} else if reflect.TypeOf(dataTemplate) == reflect.TypeOf((*bool)(nil)) {
			(*(dataTemplate).(*bool)) = value.(bool)
		}
		return nil
	}
	assert.Fail(session.t, "Unexpected call to GetRequestHeader")
	return nil
}

func (session *dummySession) Attach(name string, value interface{}) bool {
	if session.attachSetValue == nil && session.attachSetSuccess == nil {
		assert.Fail(session.t, "Unexpected call to Attach")
		return false
	}
	var val, valFound = session.attachSetValue[name]
	assert.True(session.t, valFound)
	assert.Equal(session.t, val, value)
	var ret, retFound = session.attachSetSuccess[name]
	assert.True(session.t, retFound)
	return ret
}

func (session *dummySession) Detach(name string) bool {
	assert.Fail(session.t, "Unexpected call to Detach")
	return false
}

func (session *dummySession) GetRawAttachment(name string) (interface{}, bool) {
	assert.Fail(session.t, "Unexpected call to GetRawAttachment")
	return nil, false
}

func (session *dummySession) GetAttachment(name string, dataTemplate interface{}) bool {
	if session.attachGetValue == nil && session.attachGetSuccess == nil {
		assert.Fail(session.t, "Unexpected call to GetAttachment")
		return false
	}
	var value, valueFound = session.attachGetValue[name]
	if valueFound {
		if reflect.TypeOf(dataTemplate) == reflect.TypeOf((*string)(nil)) {
			(*(dataTemplate).(*string)) = value.(string)
		} else if reflect.TypeOf(dataTemplate) == reflect.TypeOf((*int)(nil)) {
			(*(dataTemplate).(*int)) = value.(int)
		} else if reflect.TypeOf(dataTemplate) == reflect.TypeOf((*uuid.UUID)(nil)) {
			(*(dataTemplate).(*uuid.UUID)) = value.(uuid.UUID)
		} else if reflect.TypeOf(dataTemplate) == reflect.TypeOf((*map[string]string)(nil)) {
			(*(dataTemplate).(*map[string]string)) = value.(map[string]string)
		} else {
			(*(dataTemplate).(*interface{})) = value
		}
	}
	return session.attachGetSuccess[name]
}

func (session *dummySession) LogMethodEnter() {
	assert.Fail(session.t, "Unexpected call to LogMethodEnter")
}

func (session *dummySession) LogMethodParameter(parameters ...interface{}) {
	assert.Fail(session.t, "Unexpected call to LogMethodParameter")
}

func (session *dummySession) LogMethodLogic(logLevel webserver.LogLevel, category string, subcategory string, messageFormat string, parameters ...interface{}) {
	assert.Fail(session.t, "Unexpected call to LogMethodLogic")
}

func (session *dummySession) LogMethodReturn(returns ...interface{}) {
	assert.Fail(session.t, "Unexpected call to LogMethodReturn")
}

func (session *dummySession) LogMethodExit() {
	assert.Fail(session.t, "Unexpected call to LogMethodExit")
}

func (session *dummySession) CreateWebcallRequest(method string, url string, payload string, sendClientCert bool) webserver.WebRequest {
	assert.Fail(session.t, "Unexpected call to CreateNetworkRequest")
	return nil
}

// mock struct
type dummyRoundTripper struct{}

func (drt *dummyRoundTripper) RoundTrip(*http.Request) (*http.Response, error) {
	return nil, nil
}
