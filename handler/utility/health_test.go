package utility

import (
	"testing"

	"github.com/google/uuid"
	"github.com/sachinmahanin/passwordstrength/config"
	"github.com/stretchr/testify/assert"
	webserver "github.com/zhongjie-cai/web-server"
)

func TestServiceHealth(t *testing.T) {

	// arrange
	var dummySessionID = uuid.New()

	var dummySessionObject = &dummySession{
		t:  t,
		id: dummySessionID,
	}

	var expectedResult = config.AppVersion

	// mock
	createMock(t)

	// expect
	logMethodLogicExpected = 1

	dummySessionObject.logFunc = func(logLevel webserver.LogLevel, category string, subcategory string, messageFormat string, parameters ...interface{}) {
		logMethodLogicCalled++
		assert.Equal(t, webserver.LogLevelInfo, logLevel)
		assert.Equal(t, "Health", category)
		assert.Equal(t, "Summary", subcategory)
		assert.Equal(t, "AppVersion = %v", messageFormat)
		assert.Equal(t, 1, len(parameters))
		assert.Equal(t, interface{}(expectedResult), parameters[0])
	}
	// SUT + act
	result, err := Health(dummySessionObject)

	// assert
	assert.Equal(t, expectedResult, result)
	assert.NoError(t, err)

	// verify
	verifyAll(t)
}
