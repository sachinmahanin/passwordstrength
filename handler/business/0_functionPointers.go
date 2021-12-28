package business

// func pointers for injection / testing: sessionutil.go
var (
	analysePasswordFunc   = analysePassword
	calculateStrengthFunc = calculateStrength
)
