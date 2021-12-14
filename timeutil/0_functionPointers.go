package timeutil

import (
	"fmt"
	"time"
)

// func pointers for injection / testing: timeutil.go
var (
	timeNow   = time.Now
	timeSleep = time.Sleep
	fmtErrorf = fmt.Errorf
)
