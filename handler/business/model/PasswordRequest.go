package model

type PasswordRequest struct {
	Password string `json:"password"`
}
type AnalysisResult struct {
	CapitalLetter                      int
	LowercaseLetters                   int
	Numbers                            int
	OtherString                        int
	Length                             int
	LettersOnly                        bool
	NumbersOnly                        bool
	NoOfMiddleSymbolOrNumber           int
	CountOfConsecutiveUppercaseLetters int
	CountOfConsecutiveLowercaseLetters int
	CountOfConsecutiveNumbers          int
	CountOfSequentialLetters           int
	CountOfSequencialNumbers           int
}
type PasswordStrengthResponse struct {
	Message string `json:"message"`
}
