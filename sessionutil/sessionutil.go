package sessionutil

import (
	webserver "github.com/zhongjie-cai/web-server"
)

// These are names for session attachments, request headers, etc.
const (
	authorizationHeaderName = "Authorization"
)

// PrepareSession serves as a pre-action function for the session to be prepared for sync service endpoints
func PrepareSession(session webserver.Session) error {
	var authorization string
	session.GetRequestHeader(
		authorizationHeaderName,
		0,
		&authorization,
	)
	if stringsHasPrefix(authorization, "bearer ") {
		session.Attach(
			authorizationHeaderName,
			authorization,
		)
	}
	return nil
}
