package miscellaneous

import (
	"net/http"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	webserver "github.com/zhongjie-cai/web-server"
)

var (
	httpRedirectHandlerExpected int
	httpRedirectHandlerCalled   int
	httpStripPrefixExpected     int
	httpStripPrefixCalled       int
	httpFileServerExpected      int
	httpFileServerCalled        int
)

func createMock(t *testing.T) {
	httpRedirectHandlerExpected = 0
	httpRedirectHandlerCalled = 0
	httpRedirectHandler = func(url string, code int) http.Handler {
		httpRedirectHandlerCalled++
		return nil
	}

	httpStripPrefixExpected = 0
	httpStripPrefixCalled = 0
	httpStripPrefix = func(prefix string, h http.Handler) http.Handler {
		httpStripPrefixCalled++
		return nil
	}
	httpFileServerExpected = 0
	httpFileServerCalled = 0
	httpFileServer = func(root http.FileSystem) http.Handler {
		httpFileServerCalled++
		return nil
	}
}

func verifyAll(t *testing.T) {
	httpRedirectHandler = http.RedirectHandler
	assert.Equal(t, httpRedirectHandlerExpected, httpRedirectHandlerCalled, "Unexpected number of calls to method httpRedirectHandler")
	httpStripPrefix = http.StripPrefix
	assert.Equal(t, httpStripPrefixExpected, httpStripPrefixCalled, "Unexpected number of calls to method httpStripPrefix")
	httpFileServer = http.FileServer
	assert.Equal(t, httpFileServerExpected, httpFileServerCalled, "Unexpected number of calls to method httpFileServer")
}

//mock structs
type dummySession struct {
	t              *testing.T
	id             uuid.UUID
	name           string
	request        *http.Request
	responseWriter http.ResponseWriter
	paramFunc      func(name string, dataTemplate interface{}) error
	logFunc        func(logLevel webserver.LogLevel, category string, subcategory string, messageFormat string, parameters ...interface{})
}

func (session *dummySession) GetID() uuid.UUID {
	assert.Fail(session.t, "Unexpected call to GetID")
	return uuid.Nil
}

func (session *dummySession) GetName() string {
	assert.Fail(session.t, "Unexpected call to GetName")
	return ""
}

func (session *dummySession) GetRequest() *http.Request {
	assert.Fail(session.t, "Unexpected call to GetRequest")
	return nil
}

func (session *dummySession) GetResponseWriter() http.ResponseWriter {
	assert.Fail(session.t, "Unexpected call to GetResponseWriter")
	return nil
}

func (session *dummySession) GetRequestBody(dataTemplate interface{}) error {
	assert.Fail(session.t, "Unexpected call to GetRequestBody")
	return nil
}

func (session *dummySession) GetRequestParameter(name string, dataTemplate interface{}) error {
	if session.paramFunc == nil {
		assert.Fail(session.t, "Unexpected call to GetRequestParameter")
		return nil
	}
	return session.paramFunc(name, dataTemplate)
}

func (session *dummySession) GetRequestQuery(name string, index int, dataTemplate interface{}) error {
	assert.Fail(session.t, "Unexpected call to GetRequestQuery")
	return nil
}

func (session *dummySession) GetRequestHeader(name string, index int, dataTemplate interface{}) error {
	assert.Fail(session.t, "Unexpected call to GetRequestHeader")
	return nil
}

func (session *dummySession) Attach(name string, value interface{}) bool {
	assert.Fail(session.t, "Unexpected call to Attach")
	return false
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
	assert.Fail(session.t, "Unexpected call to GetAttachment")
	return false
}

func (session *dummySession) LogMethodEnter() {
	assert.Fail(session.t, "Unexpected call to LogMethodEnter")
}

func (session *dummySession) LogMethodParameter(parameters ...interface{}) {
	assert.Fail(session.t, "Unexpected call to LogMethodParameter")
}

func (session *dummySession) LogMethodLogic(logLevel webserver.LogLevel, category string, subcategory string, messageFormat string, parameters ...interface{}) {
	if session.logFunc == nil {
		assert.Fail(session.t, "Unexpected call to LogMethodLogic")
		return
	}
	session.logFunc(logLevel, category, subcategory, messageFormat, parameters...)
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

type dummyHandlerStruct struct {
}

func (dhs *dummyHandlerStruct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}
