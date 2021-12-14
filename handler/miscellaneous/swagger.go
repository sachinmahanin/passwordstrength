package miscellaneous

import (
	"net/http"
)

func SwaggerRedirect() http.Handler {
	return httpRedirectHandler(
		"/docs/",
		http.StatusPermanentRedirect,
	)
}

func SwaggerHandler() http.Handler {
	return httpStripPrefix(
		"/docs/",
		httpFileServer(
			http.Dir("./docs"),
		),
	)
}
