package miscellaneous

import (
	"net/http"
	"testing"

	"github.com/sachinmahanin/passwordStrength/config"
	"github.com/stretchr/testify/assert"
)

func TestSwaggerRedirect(t *testing.T) {

	// arrange
	var dummyHandler = &dummyHandlerStruct{}

	// mock
	createMock(t)

	// expect
	httpRedirectHandlerExpected = 1
	httpRedirectHandler = func(url string, code int) http.Handler {
		httpRedirectHandlerCalled++
		assert.Equal(t, "/docs/", url)
		assert.Equal(t, http.StatusPermanentRedirect, code)
		return dummyHandler
	}

	// SUT + act
	var result = SwaggerRedirect()

	// assert
	assert.Equal(t, dummyHandler, result)

	// verify
	verifyAll(t)
}

func TestSwaggerSwaggerHandler(t *testing.T) {

	// arrange
	var dummyAppPath = "."
	var dummyFileHandler = &dummyHandlerStruct{}
	var dummyForwardedHandler = &dummyHandlerStruct{}
	// stub
	config.AppPath = dummyAppPath

	// mock
	createMock(t)

	// expect
	httpStripPrefixExpected = 1
	httpStripPrefix = func(prefix string, h http.Handler) http.Handler {
		httpStripPrefixCalled++
		assert.Equal(t, "/docs/", prefix)
		assert.Equal(t, dummyFileHandler, h)
		return dummyForwardedHandler
	}
	httpFileServerExpected = 1
	httpFileServer = func(root http.FileSystem) http.Handler {
		httpFileServerCalled++
		assert.Equal(t, http.Dir(dummyAppPath+"/docs"), root)
		return dummyFileHandler
	}

	// SUT + act
	var result = SwaggerHandler()

	// assert
	assert.Equal(t, dummyForwardedHandler, result)

	// verify
	verifyAll(t)
}
