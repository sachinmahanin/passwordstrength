package logger

import (
	"github.com/google/uuid"
	"github.com/sachinmahanin/passwordstrength/config"
	webserver "github.com/zhongjie-cai/web-server"
)

// Elastic default values
const (
	indexTypeLog           string = "log"
	loggingDataVersion     byte   = 2
	defaultTimeStampFormat string = "2006-01-02T15:04:05.999Z07:00"
)

type logEntry struct {
	Time        string    `json:"@timestamp"`
	DataVersion byte      `json:"data_version"`
	IndexType   string    `json:"type"`
	RoleType    string    `json:"roletype"`
	Host        string    `json:"host"`
	Application string    `json:"application"`
	AppVersion  string    `json:"application_version"`
	Endpoint    string    `json:"endpoint"`
	Descrition  string    `json:"description"`
	LogLevel    string    `json:"level"`
	Category    string    `json:"category"`
	Subcategory string    `json:"subcategory"`
	SessionID   uuid.UUID `json:"session_id"`
	Name        string    `json:"name"`
}

// CustomizeLoggingFunc customizes the logging for application
func CustomizeLoggingFunc(
	session webserver.Session,
	logType webserver.LogType,
	logLevel webserver.LogLevel,
	category string,
	subcategory string,
	description string,
) {
	var sessionID = session.GetID()
	var endpoint = session.GetName()
	var logTime = timeutilGetTimeNowUTC()
	var logTimeValue = logTime.Format(defaultTimeStampFormat)
	if !config.AllowedLogType.HasFlag(logType) || config.AllowedLogLevel < logLevel {
		return
	}
	var logEntry = logEntry{
		Application: config.AppName,
		AppVersion:  config.AppVersion,
		DataVersion: loggingDataVersion,
		Descrition:  description,
		Endpoint:    endpoint,
		LogLevel:    logLevel.String(),
		Category:    category,
		Subcategory: subcategory,
		SessionID:   sessionID,
		Host:        config.HostName,
		IndexType:   indexTypeLog,
		Time:        logTimeValue,
		Name:        logType.String(),
	}
	var logEntryBytes, _ = jsonMarshal(
		logEntry,
	)
	fmtPrintln(
		string(
			logEntryBytes,
		),
	)
}
