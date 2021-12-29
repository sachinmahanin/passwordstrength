package web

import (
	"net/http"
	"testing"

	"github.com/sachinmahanin/passwordstrength/handler/miscellaneous"
	"github.com/stretchr/testify/assert"
	webserver "github.com/zhongjie-cai/web-server"
)

var (
	miscellaneousSwaggerHandlerExpected  int
	miscellaneousSwaggerHandlerCalled    int
	miscellaneousSwaggerRedirectExpected int
	miscellaneousSwaggerRedirectCalled   int
	registeredMiscRoutesFuncExpected     int
	registeredMiscRoutesFuncCalled       int
	registeredUtilityRoutesFuncExpected  int
	registeredUtilityRoutesFuncCalled    int
	registeredBusinessRoutesFuncExpected int
	registeredBusinessRoutesFuncCalled   int
)

func createMock(t *testing.T) {
	miscellaneousSwaggerHandlerExpected = 0
	miscellaneousSwaggerHandlerCalled = 0
	miscellaneousSwaggerHandler = func() http.Handler {
		miscellaneousSwaggerHandlerCalled++
		return nil
	}
	miscellaneousSwaggerRedirectExpected = 0
	miscellaneousSwaggerRedirectCalled = 0
	miscellaneousSwaggerRedirect = func() http.Handler {
		miscellaneousSwaggerRedirectCalled++
		return nil
	}
	registeredUtilityRoutesFuncExpected = 0
	registeredUtilityRoutesFuncCalled = 0
	registeredUtilityRoutesFunc = func() []webserver.Route {
		registeredUtilityRoutesFuncCalled++
		return nil
	}
	registeredBusinessRoutesFuncExpected = 0
	registeredBusinessRoutesFuncCalled = 0
	registeredBusinessRoutesFunc = func() []webserver.Route {
		registeredBusinessRoutesFuncCalled++
		return nil
	}
}

func verifyAll(t *testing.T) {
	miscellaneousSwaggerHandler = miscellaneous.SwaggerHandler
	assert.Equal(t, miscellaneousSwaggerHandlerExpected, miscellaneousSwaggerHandlerCalled, "Unexpected number of calls to method miscellaneousSwaggerHandler")
	miscellaneousSwaggerRedirect = miscellaneous.SwaggerRedirect
	assert.Equal(t, miscellaneousSwaggerHandlerExpected, miscellaneousSwaggerHandlerCalled, "Unexpected number of calls to method miscellaneousSwaggerRedirect")
	registeredUtilityRoutesFunc = registeredUtilityRoutes
	assert.Equal(t, registeredUtilityRoutesFuncExpected, registeredUtilityRoutesFuncCalled, "Unexpected number of calls to method registeredUtilityRoutesFunc")
	registeredBusinessRoutesFunc = registeredBusinessRoutes
	assert.Equal(t, registeredBusinessRoutesFuncExpected, registeredBusinessRoutesFuncCalled, "Unexpected number of calls to method registeredBusinessRoutesFunc")
}

// mock struct
type dummyHandlerStruct struct {
}

func (dhs *dummyHandlerStruct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}
