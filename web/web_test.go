package web

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/sachinmahanin/passwordStrength/handler/business"
	"github.com/sachinmahanin/passwordStrength/handler/utility"
	"github.com/stretchr/testify/assert"
	webserver "github.com/zhongjie-cai/web-server"
)

func funcEquals(
	t *testing.T,
	expected interface{},
	actual interface{},
	msgAndArgs ...interface{},
) {
	var expectedPointer = fmt.Sprintf(
		"%v",
		reflect.ValueOf(
			expected,
		),
	)
	var actualPointer = fmt.Sprintf(
		"%v",
		reflect.ValueOf(
			actual,
		),
	)
	assert.Equal(
		t,
		expectedPointer,
		actualPointer,
		msgAndArgs...,
	)
}

func TestRegisteredUtilityRoutes(t *testing.T) {
	// arrange

	// mock
	createMock(t)

	// SUT + act
	var result = registeredUtilityRoutes()

	// assert
	assert.Equal(t, 1, len(result))

	assert.Equal(t, "utility.Health", result[0].Endpoint)
	assert.Equal(t, http.MethodGet, result[0].Method)
	assert.Equal(t, "/health", result[0].Path)
	funcEquals(t, utility.Health, result[0].ActionFunc)
	assert.Empty(t, result[0].Parameters)
	assert.Empty(t, result[0].Queries)

	// verify
	verifyAll(t)
}

func TestRegisteredBusinessRoutes(t *testing.T) {
	// arrange

	// mock
	createMock(t)

	// SUT + act
	var result = registeredBusinessRoutes()

	// assert
	assert.Equal(t, 1, len(result))
	assert.Equal(t, "business.passwordStrength", result[0].Endpoint)
	assert.Equal(t, http.MethodPost, result[0].Method)
	assert.Equal(t, "/PassportStrength", result[0].Path)
	funcEquals(t, business.Strength, result[0].ActionFunc)
	assert.Empty(t, result[0].Parameters)
	assert.Empty(t, result[0].Queries)

	// verify
	verifyAll(t)
}

func TestRegisteredRoutes(t *testing.T) {
	// arrange
	var dummyUtilityRoutes = []webserver.Route{
		webserver.Route{
			Endpoint: "dummyUtilityRoute1",
		},
		webserver.Route{
			Endpoint: "dummyUtilityRoute2",
		},
	}
	var dummyBusinessRoutes = []webserver.Route{
		webserver.Route{
			Endpoint: "dummyBusinessRoute1",
		},
		webserver.Route{
			Endpoint: "dummyBusinessRoute2",
		},
	}

	// mock
	createMock(t)

	// expect
	registeredUtilityRoutesFuncExpected = 1
	registeredUtilityRoutesFunc = func() []webserver.Route {
		registeredUtilityRoutesFuncCalled++
		return dummyUtilityRoutes
	}
	registeredBusinessRoutesFuncExpected = 1
	registeredBusinessRoutesFunc = func() []webserver.Route {
		registeredBusinessRoutesFuncCalled++
		return dummyBusinessRoutes
	}

	// SUT + act
	var result = RegisteredRoutes()

	// assert
	assert.Equal(t, 4, len(result))
	assert.Equal(t, dummyUtilityRoutes[0], result[0])
	assert.Equal(t, dummyUtilityRoutes[1], result[1])
	assert.Equal(t, dummyBusinessRoutes[0], result[2])
	assert.Equal(t, dummyBusinessRoutes[1], result[3])

	// verify
	verifyAll(t)
}
