package sessionutil

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/google/uuid"
	webserver "github.com/zhongjie-cai/web-server"
)

var cnRegex = regexp.MustCompile("[Cc][Nn]=([^,\\/]+)")

// func pointers for injection / testing: sessionutil.go
var (
	strconvItoa               = strconv.Itoa
	uuidNew                   = uuid.New
	fmtSprintf                = fmt.Sprintf
	stringsHasSuffix          = strings.HasSuffix
	stringsHasPrefix          = strings.HasPrefix
	stringsSplit              = strings.Split
	stringsTrimSpace          = strings.TrimSpace
	cnRegexFindStringSubmatch = cnRegex.FindStringSubmatch
	webserverGetBadRequest    = webserver.GetBadRequest
	fmtErrorf                 = fmt.Errorf
)
