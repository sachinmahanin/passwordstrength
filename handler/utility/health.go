package utility

import (
	"fmt"

	"github.com/sachinmahanin/passwordStrength/config"
	webserver "github.com/zhongjie-cai/web-server"
)

func Health(session webserver.Session) (interface{}, error) {
	fmt.Println("inside health")
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
