package logger

import (
	"fmt"
	"log"
)

var (
	loggerPrefix  = ""
	hashMaxLength = 10
)

func InitLogger(appName, appHash, appBuild string) {
	if len(appHash) > hashMaxLength {
		appHash = appHash[0:9]
	}
	loggerPrefix = fmt.Sprintf("[%s] (%s,%s) ", appName, appHash, appBuild)
	log.Print(loggerPrefix, "App stated")
}

func Info(message string, args ...interface{}) {
	log.Print(loggerPrefix, message, fmt.Sprint(args...))
}

func Warning(message string, args ...interface{}) {
	log.Print(loggerPrefix, message, fmt.Sprint(args...))
}

func Error(message string, err error) {
	log.Print("[ERROR]", loggerPrefix, message, err.Error())
}

func ErrorDetailed(message string, err error, logFields map[string]interface{}) {
	log.Printf("[ERROR] %s %s err: %s, %v", loggerPrefix, message, err.Error(), logFields)
}

func Fatal(message string, err error) {
	log.Fatalf("[FATALITY] %s %s err: %s", loggerPrefix, message, err.Error())
}

func FatalDetailed(message string, err error, logFields map[string]interface{}) {
	log.Fatalf("[FATALITY] %s %s err: %s, %v", loggerPrefix, message, err.Error(), logFields)
}
