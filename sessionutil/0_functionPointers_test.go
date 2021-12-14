package sessionutil

import (
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"testing"

	"github.com/google/uuid"

	"github.com/stretchr/testify/assert"
	webserver "github.com/zhongjie-cai/web-server"
)

var (
	strconvItoaExpected                    int
	strconvItoaCalled                      int
	uuidNewExpected                        int
	uuidNewCalled                          int
	fmtSprintfExpected                     int
	fmtSprintfCalled                       int
	stringsHasSuffixExpected               int
	stringsHasSuffixCalled                 int
	stringsHasPrefixExpected               int
	stringsHasPrefixCalled                 int
	stringsSplitExpected                   int
	stringsSplitCalled                     int
	stringsTrimSpaceExpected               int
	stringsTrimSpaceCalled                 int
	extractCNameFuncExpected               int
	extractCNameFuncCalled                 int
	getCorrelationIDFromHeaderFuncExpected int
	getCorrelationIDFromHeaderFuncCalled   int
	getCNameFromHTTPHeaderFuncExpected     int
	getCNameFromHTTPHeaderFuncCalled       int
	getCNameFromHTTPRequestFuncExpected    int
	getCNameFromHTTPRequestFuncCalled      int
	getCallerIDFromTLSStateFuncExpected    int
	getCallerIDFromTLSStateFuncCalled      int
	prepareCorrelationIDFuncExpected       int
	prepareCorrelationIDFuncCalled         int
	preparePassThroughHeadersFuncExpected  int
	preparePassThroughHeadersFuncCalled    int
	prepareCallerIDFuncExpected            int
	prepareCallerIDFuncCalled              int
	stringsContainsExpected                int
	stringsContainsCalled                  int
	getCorrelationIDFuncExpected           int
	getCorrelationIDFuncCalled             int
	cnRegexFindStringSubmatchExpeted       int
	cnRegexFindStringSubmatchCalled        int
	webserverGetBadRequestExpected         int
	webserverGetBadRequestCalled           int
)

func createMock(t *testing.T) {
	strconvItoaExpected = 0
	strconvItoaCalled = 0
	strconvItoa = func(i int) string {
		strconvItoaCalled++
		return ""
	}
	uuidNewExpected = 0
	uuidNewCalled = 0
	uuidNew = func() uuid.UUID {
		uuidNewCalled++
		return uuid.Nil
	}
	webserverGetBadRequestExpected = 0
	webserverGetBadRequestCalled = 0
	webserverGetBadRequest = func(errorMessage string, innerErrors ...error) webserver.AppError {
		return nil
	}
	fmtSprintfExpected = 0
	fmtSprintfCalled = 0
	fmtSprintf = func(format string, a ...interface{}) string {
		fmtSprintfCalled++
		return ""
	}
	stringsHasSuffixExpected = 0
	stringsHasSuffixCalled = 0
	stringsHasSuffix = func(s, suffix string) bool {
		stringsHasSuffixCalled++
		return false
	}
	stringsHasPrefixExpected = 0
	stringsHasPrefixCalled = 0
	stringsHasPrefix = func(s, prefix string) bool {
		stringsHasPrefixCalled++
		return false
	}
	stringsSplitExpected = 0
	stringsSplitCalled = 0
	stringsSplit = func(s, sep string) []string {
		stringsSplitCalled++
		return nil
	}
	stringsTrimSpaceExpected = 0
	stringsTrimSpaceCalled = 0
	stringsTrimSpace = func(s string) string {
		stringsTrimSpaceCalled++
		return ""
	}

	cnRegexFindStringSubmatchExpeted = 0
	cnRegexFindStringSubmatchCalled = 0
	cnRegexFindStringSubmatch = func(s string) []string {
		cnRegexFindStringSubmatchCalled++
		return nil
	}
}

func verifyAll(t *testing.T) {
	strconvItoa = strconv.Itoa
	assert.Equal(t, strconvItoaExpected, strconvItoaCalled, "Unexpected number of calls to method strconvItoa")
	uuidNew = uuid.New
	assert.Equal(t, uuidNewExpected, uuidNewCalled, "Unexpected number of calls to method uuidNew")
	fmtSprintf = fmt.Sprintf
	assert.Equal(t, fmtSprintfExpected, fmtSprintfCalled, "Unexpected number of calls to method fmtSprintf")
	stringsHasSuffix = strings.HasSuffix
	assert.Equal(t, stringsHasSuffixExpected, stringsHasSuffixCalled, "Unexpected number of calls to method stringsHasSuffix")
	stringsHasPrefix = strings.HasPrefix
	assert.Equal(t, stringsHasPrefixExpected, stringsHasPrefixCalled, "Unexpected number of calls to method stringsHasPrefix")
	stringsSplit = strings.Split
	assert.Equal(t, stringsSplitExpected, stringsSplitCalled, "Unexpected number of calls to method stringsSplit")
	stringsTrimSpace = strings.TrimSpace
	assert.Equal(t, stringsTrimSpaceExpected, stringsTrimSpaceCalled, "Unexpected number of calls to method stringsTrimSpace")
	webserverGetBadRequest = webserver.GetBadRequest
	assert.Equal(t, webserverGetBadRequestExpected, webserverGetBadRequestCalled, "Unexpected number of calls to webserverGetBadRequest")
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
	assert.Fail(session.t, "Unexpected call to GetID")
	return uuid.Nil
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

type dummyResponseWriter struct {
	t *testing.T
}

func (drw *dummyResponseWriter) Header() http.Header {
	assert.Fail(drw.t, "Unexpected number of calls to method Header")
	return nil
}

func (drw *dummyResponseWriter) Write([]byte) (int, error) {
	assert.Fail(drw.t, "Unexpected number of calls to method Write")
	return 0, nil
}

func (drw *dummyResponseWriter) WriteHeader(statusCode int) {
	assert.Fail(drw.t, "Unexpected number of calls to method WriteHeader")
}

type dummyTransaction struct {
	http.ResponseWriter
	t          *testing.T
	attributes map[string]struct {
		value interface{}
		err   error
	}
	noticeError struct {
		input  error
		output error
	}
}

func (dt *dummyTransaction) End() error {
	assert.Fail(dt.t, "Unexpected number of calls to method End")
	return nil
}

func (dt *dummyTransaction) Ignore() error {
	assert.Fail(dt.t, "Unexpected number of calls to method Ignore")
	return nil
}

func (dt *dummyTransaction) SetName(name string) error {
	assert.Fail(dt.t, "Unexpected number of calls to method SetName")
	return nil
}

func (dt *dummyTransaction) NoticeError(err error) error {
	assert.Equal(dt.t, dt.noticeError.input, err)
	return dt.noticeError.output
}

func (dt *dummyTransaction) AddAttribute(key string, value interface{}) error {
	var item, found = dt.attributes[key]
	if !found {
		assert.Fail(dt.t, "Unexpected number of calls to method AddAttribute", key)
		return nil
	}
	assert.Equal(dt.t, item.value, value)
	return item.err
}

func (dt *dummyTransaction) IsSampled() bool {
	assert.Fail(dt.t, "Unexpected number of calls to method IsSampled")
	return false
}
