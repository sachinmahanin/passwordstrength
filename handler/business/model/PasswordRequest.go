package model

type PasswordRequest struct {
	Password string `json:"password"`
}
type AnalysisResult struct {
	CapitalLetter    int
	LowercaseLetters int
	Number           int
	OtherString      int
	Length           int
	LettersOnly      bool
	NumbersOnly      bool
}
