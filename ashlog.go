package ashlog

import (
	"errors"
	"fmt"
	"io"
	"log"
	"time"
)

// Category type
type Category string

// LogLevel type
type LogLevel int

const (
	// LOGAPP is a log category related to application
	LOGAPP Category = "app"
	// LOGDBM is a log category related to database
	LOGDBM Category = "dbm"
	// LOGNET is a log category related to net
	LOGNET Category = "net"
	// LOGCONF is a log category related to configuration
	LOGCONF Category = "cnf"

	// TRACE logs
	TRACE LogLevel = 0
	// DEBUG logs
	DEBUG LogLevel = 1
	// INFO logs
	INFO LogLevel = 2
	// WARNING logs
	WARNING LogLevel = 3
	// ERROR logs
	ERROR LogLevel = 4
	// CRITICAL logs
	CRITICAL LogLevel = 5
)

var showThisLogLevelAndBelow = 5

// createMsgPrintLog logs the given message with the given level only if the
// given level is not bigger than the configured valid log level
func createMsgPrintLog(cat Category, lvl LogLevel, msg string) {
	if int(lvl) > showThisLogLevelAndBelow {
		return
	}

	var lvlString string

	switch lvl {
	case TRACE:
		lvlString = "T"
	case DEBUG:
		lvlString = "D"
	case INFO:
		lvlString = "I"
	case WARNING:
		lvlString = "W"
	case ERROR:
		lvlString = "E"
	case CRITICAL:
		lvlString = "C"
	}
	log.Println(fmt.Sprintf(" %s/%s: %s", lvlString, cat, msg))
}

// InitLogger initialized log file
func InitLogger(writers []io.Writer, lvl int) error {
	if lvl > 5 || lvl < 0 {
		return errors.New("Log level is not between 0 and 5. uses default log level: 5")
	}

	mw := io.MultiWriter(writers...)
	log.SetOutput(mw)
	log.SetPrefix(time.Now().Format("@[20060102-150405]")) // log prefix (timestamp)
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))   // Don't use default timestamp
	showThisLogLevelAndBelow = lvl
	return nil
}

// Info can be used to log info
func Info(cat Category, msg string) {
	createMsgPrintLog(cat, INFO, msg)
}

// Trace can be used for trace
func Trace(cat Category, msg string) {
	createMsgPrintLog(cat, TRACE, msg)
}

// Debug can be used for debug
func Debug(cat Category, msg string) {
	createMsgPrintLog(cat, DEBUG, msg)
}

// Warning can be used for log warnings
func Warning(cat Category, msg string) {
	createMsgPrintLog(cat, WARNING, msg)
}

// Error can be used for log error mesages
func Error(cat Category, msg string) {
	createMsgPrintLog(cat, ERROR, msg)
}

// Critical can be used for log critical mesages
func Critical(cat Category, msg string) {
	createMsgPrintLog(cat, CRITICAL, msg)
}
