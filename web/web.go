package web

import (
	"fmt"
	"net/http"

	"github.com/sachinmahanin/passwordstrength/handler/business"
	"github.com/sachinmahanin/passwordstrength/handler/utility"
	webserver "github.com/zhongjie-cai/web-server"
)

//RegisteredStatics returns the registered static content handlers for web service hosting
func RegisteredStatics() []webserver.Static {
	fmt.Println("RegisteredStatics")
	return []webserver.Static{
		//add static routes
	}
}

func registeredBusinessRoutes() []webserver.Route {
	fmt.Println("registeredBusinessRoutes")
	return []webserver.Route{
		webserver.Route{
			Endpoint:   "business.passwordstrength",
			Method:     http.MethodPost,
			Path:       "/PassportStrength",
			ActionFunc: business.Strength,
		},
	}
}

func registeredUtilityRoutes() []webserver.Route {
	fmt.Println("registeredUtilityRoutes")
	return []webserver.Route{
		webserver.Route{
			Endpoint:   "utility.Health",
			Method:     http.MethodGet,
			Path:       "/health",
			ActionFunc: utility.Health,
		},
	}
}

// RegisteredRoutes returns the registered route handlers for web service hosting
func RegisteredRoutes() []webserver.Route {
	fmt.Println("RegisteredRoutes")
	var allRoutes = []webserver.Route{}
	allRoutes = append(
		allRoutes,
		registeredUtilityRoutesFunc()...,
	)
	allRoutes = append(
		allRoutes,
		registeredBusinessRoutesFunc()...,
	)
	return allRoutes
}
