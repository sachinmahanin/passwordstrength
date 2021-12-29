package business

import (
	"net/http"
	"testing"

	"github.com/google/uuid"
	"github.com/sachinmahanin/passwordstrength/handler/business/model"
	webserver "github.com/zhongjie-cai/web-server"

	"github.com/stretchr/testify/assert"
)

var (
	analysePasswordFuncExpected   int
	analysePasswordFuncCalled     int
	calculateStrengthFuncExpected int
	calculateStrengthFuncCalled   int
	sessionGetRequestBodyExpected int
	sessionGetRequestBodyCalled   int
)

func createMock(t *testing.T) {
	sessionGetRequestBodyExpected = 0
	sessionGetRequestBodyCalled = 0

	analysePasswordFuncExpected = 0
	analysePasswordFuncCalled = 0
	analysePasswordFunc = func(password string) (analysisData model.AnalysisResult) {
		analysePasswordFuncCalled++
		return model.AnalysisResult{}
	}
	calculateStrengthFuncExpected = 0
	calculateStrengthFuncCalled = 0
	calculateStrengthFunc = func(analysisData model.AnalysisResult) int {
		calculateStrengthFuncCalled++
		return 0
	}

}

func verifyAll(t *testing.T) {
	analysePasswordFunc = analysePassword
	assert.Equal(t, analysePasswordFuncExpected, analysePasswordFuncCalled, "Unexpected number of calls to method analysePasswordFunc")
	calculateStrengthFunc = calculateStrength
	assert.Equal(t, calculateStrengthFuncExpected, calculateStrengthFuncCalled, "Unexpected number of calls to method calculateStrengthFunc")
	assert.Equal(t, sessionGetRequestBodyExpected, sessionGetRequestBodyCalled, "Unexpected number of calls to method sessionGetRequestBody")

}

//mock structs
type dummySession struct {
	t        *testing.T
	id       uuid.UUID
	bodyFunc func(dataTemplate interface{}) error
	logFunc  func(logLevel webserver.LogLevel, category string, subcategory string, messageFormat string, parameters ...interface{})
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
	if session.bodyFunc == nil {
		assert.Fail(session.t, "Unexpected call to GetRequestBody")
		return nil
	}
	return session.bodyFunc(dataTemplate)
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
