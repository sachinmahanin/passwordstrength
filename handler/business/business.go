package business

import (
	"strconv"

	"github.com/sachinmahanin/passwordStrength/handler/business/model"

	"regexp"

	webserver "github.com/zhongjie-cai/web-server"
)

func calculateStrength(analysisData model.AnalysisResult) int {

	var strength = 0
	//Additions
	//Number of characters
	strength += analysisData.Length * 4
	//Uppercase letters
	strength += (analysisData.Length - analysisData.CapitalLetter) * 2
	//Lowercase Letters
	strength += (analysisData.Length - analysisData.LowercaseLetters) * 2
	//Numbers
	strength += analysisData.Numbers * 4
	//Symbols
	strength += analysisData.OtherString * 6
	//Middle numbers or symbols
	strength += analysisData.NoOfMiddleSymbolOrNumber * 2
	//Requirements
	var requirementCount = 0

	if analysisData.Numbers >= 1 {
		requirementCount += 1
	}
	if analysisData.LowercaseLetters >= 1 {
		requirementCount += 1
	}
	if analysisData.CapitalLetter >= 1 {
		requirementCount += 1
	}
	if analysisData.OtherString >= 1 {
		requirementCount += 1
	}
	if analysisData.Length >= 8 {
		requirementCount += 1
	}
	strength += requirementCount * 2

	//Deductions
	//Letters only
	if analysisData.LettersOnly == true {
		strength -= analysisData.Length
	}
	//Numbers only
	if analysisData.NumbersOnly == true {
		strength -= analysisData.Length
	}
	//Consecutive uppercase letters
	strength -= analysisData.CountOfConsecutiveUppercaseLetters * 2

	//Consecutive lowercase letters
	strength -= analysisData.CountOfConsecutiveLowercaseLetters * 2

	//Consecutive numbers
	strength -= analysisData.CountOfConsecutiveNumbers * 2

	//Sequential letters (3+)
	strength -= analysisData.CountOfSequentialLetters * 3

	//Sequential numbers (3+)
	strength -= analysisData.CountOfSequencialNumbers * 3

	return strength

}
func analysePassword(password string) (analysisData model.AnalysisResult) {

	analysisData.Length = len(password)
	if len(password) == 0 {
		return analysisData
	}

	for i := analysisData.OtherString; i < analysisData.Length; i++ {
		switch {
		case 64 < password[i] && password[i] < 91:
			analysisData.CapitalLetter += 1
		case 96 < password[i] && password[i] < 123:
			analysisData.LowercaseLetters += 1
		case 47 < password[i] && password[i] < 58:
			analysisData.Numbers += 1
			if i != 0 && i != analysisData.Length-1 {
				analysisData.NoOfMiddleSymbolOrNumber += 1
			}
		default:
			analysisData.OtherString += 1
			if i != 0 && i != analysisData.Length-1 {
				analysisData.NoOfMiddleSymbolOrNumber += 1
			}
		}
	}

	analysisData.Length = len(password)
	/////////////////////////////////////////////
	if analysisData.Length == analysisData.CapitalLetter+analysisData.LowercaseLetters {
		analysisData.LettersOnly = true
	}

	if analysisData.Length == analysisData.Numbers {
		analysisData.NumbersOnly = true
	}
	var countOfConsecutiveUppercaseLetters = 0
	r := regexp.MustCompile(`[A-Z]+`)
	matches := r.FindAllString(password, -1)
	if len(matches) > 0 {
		for i := 0; i < len(matches); i++ {
			countOfConsecutiveUppercaseLetters = countOfConsecutiveUppercaseLetters + len(matches[i]) - 1
		}
	}
	analysisData.CountOfConsecutiveUppercaseLetters = countOfConsecutiveUppercaseLetters
	/////////////////////////////////////////////
	var countOfConsecutiveLowercaseLetters = 0
	r = regexp.MustCompile(`[a-z]+`)
	matches = r.FindAllString(password, -1)
	if len(matches) > 0 {
		for i := 0; i < len(matches); i++ {
			countOfConsecutiveLowercaseLetters = countOfConsecutiveLowercaseLetters + len(matches[i]) - 1
		}
	}
	analysisData.CountOfConsecutiveLowercaseLetters = countOfConsecutiveLowercaseLetters
	/////////////////////////////////////////////
	var countOfConsecutiveNumbers = 0
	r = regexp.MustCompile(`[0-9]+`)
	matches = r.FindAllString(password, -1)
	if len(matches) > 0 {
		for i := 0; i < len(matches); i++ {
			countOfConsecutiveNumbers = countOfConsecutiveNumbers + len(matches[i]) - 1
		}
	}
	analysisData.CountOfConsecutiveNumbers = countOfConsecutiveNumbers

	/////////////////////////////////////////////
	var countOfSequentialLetters = 0
	r = regexp.MustCompile(`(?i)(abc|bcd|cde|def|efg|fgh|ghi|hij|ijk|jkl|klm|lmn|mno|nop|opq|pqr|qrs|rst|stu|tuv|uvw|vwx|wxy|xyz)+`)
	matches = r.FindAllString(password, -1)
	if len(matches) > 0 {
		var maxLength = 0
		for i := 0; i < len(matches); i++ {
			if maxLength < len(matches[i]) {
				maxLength = len(matches[i])
			}
		}
		if maxLength-2 > 0 {
			countOfSequentialLetters = maxLength - 2
		}
	}
	analysisData.CountOfSequentialLetters = countOfSequentialLetters
	/////////////////////////////////////////////
	/////////////////////////////////////////////
	var countOfSequencialNumbers = 0
	r = regexp.MustCompile(`(012|123|234|345|456|567|678|789)+`)
	matches = r.FindAllString(password, -1)
	if len(matches) > 0 {
		var maxLength = 0
		for i := 0; i < len(matches); i++ {
			if maxLength < len(matches[i]) {
				maxLength = len(matches[i])
			}
			if maxLength-2 > 0 {
				countOfSequencialNumbers = maxLength - 2
			}
		}
	}
	analysisData.CountOfSequencialNumbers = countOfSequencialNumbers
	/////////////////////////////////////////////
	return analysisData
}

func Strength(session webserver.Session) (interface{}, error) {

	var passwordStrengthRequest model.PasswordRequest
	var bodyError = session.GetRequestBody(
		&passwordStrengthRequest,
	)
	if bodyError != nil {
		return nil, bodyError
	}

	result := analysePasswordFunc(passwordStrengthRequest.Password)
	var strength = calculateStrengthFunc(result)
	var complexityType = "Very Weak"
	switch {
	case strength >= 80:
		complexityType = "Very Strong"
	case strength >= 70:
		complexityType = "Strong"
	case strength >= 60:
		complexityType = "Week"
	default:
		complexityType = "Very Weak"
	}
	var msg = "Your password strength is " + strconv.Itoa(strength) + ". complexity Type =" + complexityType
	var resp = &model.PasswordStrengthResponse{
		Message: msg,
	}
	return resp, nil

}
