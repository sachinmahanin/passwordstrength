package sessionutil

import (
	"strings"
	"testing"
)

func TestPrepareSession_HappyPath(t *testing.T) {
	// arrange
	var dummyAuthorization = "bearer some authorization"
	var dummySessionObject = &dummySession{
		t: t,
		headerValue: map[string]interface{}{
			authorizationHeaderName: dummyAuthorization,
		},
		attachSetValue: map[string]interface{}{
			authorizationHeaderName: dummyAuthorization,
		},
		attachSetSuccess: map[string]bool{
			authorizationHeaderName: true,
		},
	}
	// mock
	createMock(t)

	// expect
	stringsHasPrefixExpected = 1
	stringsHasPrefix = func(s, prefix string) bool {
		stringsHasPrefixCalled++
		return strings.HasPrefix(s, prefix)
	}
	// SUT + act
	var _ = PrepareSession(
		dummySessionObject,
	)

	// assert
	//	assert.Equal(t, dummyCorrelationIDError, err)

	// verify
	verifyAll(t)
}
