package logger

import (
	"encoding/json"
	"fmt"

	"github.com/sachinmahanin/passwordstrength/timeutil"
)

// func pointers for injection / testing: logger.go
var (
	jsonMarshal           = json.Marshal
	fmtPrintln            = fmt.Println
	timeutilGetTimeNowUTC = timeutil.GetTimeNowUTC
)
