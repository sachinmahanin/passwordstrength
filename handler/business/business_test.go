package business

import (
	"errors"
	"strconv"
	"testing"

	"github.com/sachinmahanin/passwordStrength/handler/business/model"
	"github.com/stretchr/testify/assert"
)

func TestAnalysePassword_EmptyPassword(t *testing.T) {

	// arrange
	var expectedResult = model.AnalysisResult{
		Length: 0,
	}

	// mock
	createMock(t)

	// expect

	// SUT + act
	var result = analysePassword("")
	// assert
	assert.Equal(t, expectedResult, result)

	// verify
	verifyAll(t)
}

func TestAnalysePassword_SingleNumberOnly(t *testing.T) {

	// arrange
	var expectedResult = model.AnalysisResult{
		Length:      1,
		NumbersOnly: true,
		Numbers:     1,
	}

	// mock
	createMock(t)

	// expect

	// SUT + act
	var result = analysePassword("1")
	// assert
	assert.Equal(t, expectedResult, result)

	// verify
	verifyAll(t)
}

func TestAnalysePassword_SequencialNumberOnly(t *testing.T) {

	// arrange
	var expectedResult = model.AnalysisResult{
		Length:                    3,
		NumbersOnly:               true,
		Numbers:                   3,
		NoOfMiddleSymbolOrNumber:  1,
		CountOfSequencialNumbers:  1,
		CountOfConsecutiveNumbers: 2,
	}

	// mock
	createMock(t)

	// expect

	// SUT + act
	var result = analysePassword("123")
	// assert
	assert.Equal(t, expectedResult, result)

	// verify
	verifyAll(t)
}

func TestAnalysePassword_SingleCapitalLetter(t *testing.T) {

	// arrange
	var expectedResult = model.AnalysisResult{
		Length:        1,
		CapitalLetter: 1,
		LettersOnly:   true,
	}

	// mock
	createMock(t)

	// expect

	// SUT + act
	var result = analysePassword("S")
	// assert
	assert.Equal(t, expectedResult, result)

	// verify
	verifyAll(t)
}

func TestAnalysePassword_AllType(t *testing.T) {

	// arrange
	var expectedResult = model.AnalysisResult{
		Length:                   4,
		CapitalLetter:            1,
		LowercaseLetters:         1,
		OtherString:              1,
		NoOfMiddleSymbolOrNumber: 1,
		Numbers:                  1,
		LettersOnly:              false,
	}

	// mock
	createMock(t)

	// expect

	// SUT + act
	var result = analysePassword("Ss1@")
	// assert
	assert.Equal(t, expectedResult, result)

	// verify
	verifyAll(t)
}

func TestAnalysePassword_AllTypeMiddleNumberAndSymbol(t *testing.T) {

	// arrange
	var expectedResult = model.AnalysisResult{
		Length:                   4,
		CapitalLetter:            1,
		LowercaseLetters:         1,
		OtherString:              1,
		NoOfMiddleSymbolOrNumber: 2,
		Numbers:                  1,
		LettersOnly:              false,
	}

	// mock
	createMock(t)

	// expect

	// SUT + act
	var result = analysePassword("S1@s")
	// assert
	assert.Equal(t, expectedResult, result)

	// verify
	verifyAll(t)
}

func TestAnalysePassword_ConsecutiveAlphabetsOnly(t *testing.T) {

	// arrange
	var expectedResult = model.AnalysisResult{
		Length:                             5,
		CountOfConsecutiveLowercaseLetters: 4,
		CountOfSequentialLetters:           1,
		LettersOnly:                        true,
		LowercaseLetters:                   5,
	}

	// mock
	createMock(t)

	// expect

	// SUT + act
	var result = analysePassword("abcde")
	// assert
	assert.Equal(t, expectedResult, result)

	// verify
	verifyAll(t)
}

func Test_CalculateStrength_AllType(t *testing.T) {

	// arrange
	var dummyAnalysisResult = model.AnalysisResult{
		Length:           4,
		LowercaseLetters: 1,
		CapitalLetter:    1,
		OtherString:      1,
		Numbers:          1,
	}

	var expectedResult = 46

	// mock
	createMock(t)

	// expect

	// SUT + act
	var result = calculateStrength(dummyAnalysisResult)
	// assert
	assert.Equal(t, expectedResult, result)

	// verify
	verifyAll(t)
}

func Test_CalculateStrength_NumbersOnlyLength8(t *testing.T) {

	// arrange
	var dummyAnalysisResult = model.AnalysisResult{
		Length:                   8,
		Numbers:                  8,
		CountOfSequencialNumbers: 8,
		NumbersOnly:              true,
	}

	var expectedResult = 68

	// mock
	createMock(t)

	// expect

	// SUT + act
	var result = calculateStrength(dummyAnalysisResult)
	// assert
	assert.Equal(t, expectedResult, result)

	// verify
	verifyAll(t)
}

func Test_CalculateStrength_LowercaseLetterOnlyLength8(t *testing.T) {

	// arrange
	var dummyAnalysisResult = model.AnalysisResult{
		Length:                   8,
		LettersOnly:              true,
		CountOfSequentialLetters: 8,
		LowercaseLetters:         8,
	}

	var expectedResult = 20

	// mock
	createMock(t)

	// expect

	// SUT + act
	var result = calculateStrength(dummyAnalysisResult)
	// assert
	assert.Equal(t, expectedResult, result)

	// verify
	verifyAll(t)
}

func Test_Strength_InvalidBody(t *testing.T) {

	// arrange
	var dummySession = &dummySession{t: t}
	var dummyPassword = "dummy"
	var dummyPasswordRequest = model.PasswordRequest{Password: dummyPassword}
	var dummyError = errors.New("someError")

	// mock
	createMock(t)

	// expect
	sessionGetRequestBodyExpected = 1
	dummySession.bodyFunc = func(dataTemplate interface{}) error {
		sessionGetRequestBodyCalled++
		(*(dataTemplate).(*model.PasswordRequest)) = dummyPasswordRequest
		return dummyError
	}

	// SUT + act
	var result, err = Strength(dummySession)
	// assert
	assert.Nil(t, result)
	assert.Equal(t, dummyError, err)

	// verify
	verifyAll(t)
}

func Test_Strength_VeryWeek(t *testing.T) {

	// arrange
	var dummySession = &dummySession{t: t}
	var dummyAnalysisResult = model.AnalysisResult{}
	var dummyPassword = "dummy"
	var dummyPasswordRequest = model.PasswordRequest{Password: dummyPassword}
	var dummyStrength = 59
	var dummyComplexityType = "Very Weak"
	var dummyMsg = "Your password strength is " + strconv.Itoa(dummyStrength) + ". complexity Type =" + dummyComplexityType
	var dummyResult = model.PasswordStrengthResponse{
		Message: dummyMsg,
	}
	// mock
	createMock(t)

	// expect
	sessionGetRequestBodyExpected = 1
	dummySession.bodyFunc = func(dataTemplate interface{}) error {
		sessionGetRequestBodyCalled++
		(*(dataTemplate).(*model.PasswordRequest)) = dummyPasswordRequest
		return nil
	}
	analysePasswordFuncExpected = 1
	analysePasswordFunc = func(password string) (analysisData model.AnalysisResult) {
		analysePasswordFuncCalled++
		assert.Equal(t, dummyPassword, password)
		return dummyAnalysisResult
	}
	calculateStrengthFuncExpected = 1
	calculateStrengthFunc = func(analysisData model.AnalysisResult) int {
		calculateStrengthFuncCalled++
		return dummyStrength
	}

	// SUT + act
	var result, err = Strength(dummySession)
	// assert
	assert.Equal(t, dummyResult, (*(result).(*model.PasswordStrengthResponse)))
	assert.Nil(t, err)

	// verify
	verifyAll(t)
}

func Test_Strength_VeryStrong(t *testing.T) {

	// arrange
	var dummySession = &dummySession{t: t}
	var dummyAnalysisResult = model.AnalysisResult{}
	var dummyPassword = "dummy"
	var dummyPasswordRequest = model.PasswordRequest{Password: dummyPassword}
	var dummyStrength = 80
	var dummyComplexityType = "Very Strong"
	var dummyMsg = "Your password strength is " + strconv.Itoa(dummyStrength) + ". complexity Type =" + dummyComplexityType
	var dummyResult = model.PasswordStrengthResponse{
		Message: dummyMsg,
	}
	// mock
	createMock(t)

	// expect
	sessionGetRequestBodyExpected = 1
	dummySession.bodyFunc = func(dataTemplate interface{}) error {
		sessionGetRequestBodyCalled++
		(*(dataTemplate).(*model.PasswordRequest)) = dummyPasswordRequest
		return nil
	}
	analysePasswordFuncExpected = 1
	analysePasswordFunc = func(password string) (analysisData model.AnalysisResult) {
		analysePasswordFuncCalled++
		assert.Equal(t, dummyPassword, password)
		return dummyAnalysisResult
	}
	calculateStrengthFuncExpected = 1
	calculateStrengthFunc = func(analysisData model.AnalysisResult) int {
		calculateStrengthFuncCalled++
		return dummyStrength
	}

	// SUT + act
	var result, err = Strength(dummySession)
	// assert
	assert.Equal(t, dummyResult, (*(result).(*model.PasswordStrengthResponse)))
	assert.Nil(t, err)

	// verify
	verifyAll(t)
}

func Test_Strength_Strong(t *testing.T) {

	// arrange
	var dummySession = &dummySession{t: t}
	var dummyAnalysisResult = model.AnalysisResult{}
	var dummyPassword = "dummy"
	var dummyPasswordRequest = model.PasswordRequest{Password: dummyPassword}
	var dummyStrength = 70
	var dummyComplexityType = "Strong"
	var dummyMsg = "Your password strength is " + strconv.Itoa(dummyStrength) + ". complexity Type =" + dummyComplexityType
	var dummyResult = model.PasswordStrengthResponse{
		Message: dummyMsg,
	}
	// mock
	createMock(t)

	// expect
	sessionGetRequestBodyExpected = 1
	dummySession.bodyFunc = func(dataTemplate interface{}) error {
		sessionGetRequestBodyCalled++
		(*(dataTemplate).(*model.PasswordRequest)) = dummyPasswordRequest
		return nil
	}
	analysePasswordFuncExpected = 1
	analysePasswordFunc = func(password string) (analysisData model.AnalysisResult) {
		analysePasswordFuncCalled++
		assert.Equal(t, dummyPassword, password)
		return dummyAnalysisResult
	}
	calculateStrengthFuncExpected = 1
	calculateStrengthFunc = func(analysisData model.AnalysisResult) int {
		calculateStrengthFuncCalled++
		return dummyStrength
	}

	// SUT + act
	var result, err = Strength(dummySession)
	// assert
	assert.Equal(t, dummyResult, (*(result).(*model.PasswordStrengthResponse)))
	assert.Nil(t, err)

	// verify
	verifyAll(t)
}

func Test_Strength_Week(t *testing.T) {

	// arrange
	var dummySession = &dummySession{t: t}
	var dummyAnalysisResult = model.AnalysisResult{}
	var dummyPassword = "dummy"
	var dummyPasswordRequest = model.PasswordRequest{Password: dummyPassword}
	var dummyStrength = 60
	var dummyComplexityType = "Week"
	var dummyMsg = "Your password strength is " + strconv.Itoa(dummyStrength) + ". complexity Type =" + dummyComplexityType
	var dummyResult = model.PasswordStrengthResponse{
		Message: dummyMsg,
	}
	// mock
	createMock(t)

	// expect
	sessionGetRequestBodyExpected = 1
	dummySession.bodyFunc = func(dataTemplate interface{}) error {
		sessionGetRequestBodyCalled++
		(*(dataTemplate).(*model.PasswordRequest)) = dummyPasswordRequest
		return nil
	}
	analysePasswordFuncExpected = 1
	analysePasswordFunc = func(password string) (analysisData model.AnalysisResult) {
		analysePasswordFuncCalled++
		assert.Equal(t, dummyPassword, password)
		return dummyAnalysisResult
	}
	calculateStrengthFuncExpected = 1
	calculateStrengthFunc = func(analysisData model.AnalysisResult) int {
		calculateStrengthFuncCalled++
		return dummyStrength
	}

	// SUT + act
	var result, err = Strength(dummySession)
	// assert
	assert.Equal(t, dummyResult, (*(result).(*model.PasswordStrengthResponse)))
	assert.Nil(t, err)

	// verify
	verifyAll(t)
}
