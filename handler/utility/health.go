package utility

import (
	"github.com/sachinmahanin/passwordStrength/config"
	webserver "github.com/zhongjie-cai/web-server"
)

func Health(session webserver.Session) (interface{}, error) {
	var appVersion = config.AppVersion

	session.LogMethodLogic(
		webserver.LogLevelInfo,
		"Health",
		"Summary",
		"AppVersion = %v",
		appVersion,
	)
	return appVersion, nil
}
