package business

import (
	"github.com/sachinmahanin/passwordStrength/handler/business/model"

	webserver "github.com/zhongjie-cai/web-server"
)

func calculateStrength(analysisData model.AnalysisResult) int {
	return 0
}
func analysePassword(password string) (analysisData model.AnalysisResult) {
	for i := analysisData.OtherString; i < len(password); i++ {
		switch {
		case 64 < password[i] && password[i] < 91:
			analysisData.CapitalLetter += 1
		case 96 < password[i] && password[i] < 123:
			analysisData.LowercaseLetters += 1
		case 47 < password[i] && password[i] < 58:
			analysisData.Number += 1
		default:
			analysisData.OtherString += 1
		}
	}

	analysisData.Length = len(password)

	if analysisData.Length == analysisData.CapitalLetter+analysisData.LowercaseLetters {
		analysisData.LettersOnly = true
	}

	if analysisData.Length == analysisData.Number {
		analysisData.NumbersOnly = true
	}

	return analysisData
}
func Strength(session webserver.Session) (interface{}, error) {

	session.LogMethodLogic(
		webserver.LogLevelInfo,
		"business",
		"Strength", "WELCOME",
	)
	var passwordStrengthRequest model.PasswordRequest
	var bodyError = session.GetRequestBody(
		&passwordStrengthRequest,
	)
	if bodyError != nil {
		return nil, bodyError
	}

	result := analysePassword(passwordStrengthRequest.Password)
	var strength = result.Number * 4
	return strength, nil
}
