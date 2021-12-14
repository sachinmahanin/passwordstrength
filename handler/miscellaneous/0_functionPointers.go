package miscellaneous

import (
	"net/http"
)

var (
	httpRedirectHandler = http.RedirectHandler
	httpStripPrefix     = http.StripPrefix
	httpFileServer      = http.FileServer
)
