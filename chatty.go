// Package chatty is a lightweight helper for events logging.
// It consists in a simple wrapper over native fmt methods.
// The OutputFormat can be defined with LOG_OUTPUT_FORMAT environment variable.
// The SeverityLevel can be defined with LOG_SEVERITY_LEVEL environment variable.
package chatty

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

const (
	severityLevelDebug   byte = 0
	severityLevelInfo    byte = 1
	severityLevelWarning byte = 2
	severityLevelError   byte = 3
	severityLevelFatal   byte = 4

	outputFormatJSON  = "json"
	outputFormatPlain = "plain"

	outputTemplateJSON  = `{"ts":"%s","level":"%s","msg":"%s"}`
	outputTemplatePlain = "%s\t%s\t%s\n"
)

var (
	severityLevelDefault       = severityLevelInfo
	outputFormatDefault        = outputFormatPlain
	outputTemplateDefault      = outputTemplatePlain
	escapeMessageStringForJSON = false
)

func init() {
	defineSeverityLevel()

	defineOutputFormat()
}

func defineSeverityLevel() {
	lsl := strings.ToLower(strings.TrimSpace(os.Getenv("LOG_SEVERITY_LEVEL")))

	if lsl == "" {
		return
	}

	samePrefixOrEqual := func(mnemonic, s string) bool {
		return mnemonic[:len(s)] == s
	}

	switch {
	case samePrefixOrEqual("debug", lsl), samePrefixOrEqual("verbose", lsl):
		severityLevelDefault = severityLevelDebug
	case samePrefixOrEqual("info", lsl):
		severityLevelDefault = severityLevelInfo
	case samePrefixOrEqual("warning", lsl):
		severityLevelDefault = severityLevelWarning
	case samePrefixOrEqual("error", lsl):
		severityLevelDefault = severityLevelError
	case samePrefixOrEqual("fatal", lsl), samePrefixOrEqual("critical", lsl):
		severityLevelDefault = severityLevelFatal
	}
}

func defineOutputFormat() {
	lof := strings.ToLower(strings.TrimSpace(os.Getenv("LOG_OUTPUT_FORMAT")))

	if lof == outputFormatJSON {
		outputFormatDefault = outputFormatJSON
		outputTemplateDefault = outputTemplateJSON
	} else {
		outputFormatDefault = outputFormatPlain
		outputTemplateDefault = outputTemplatePlain
	}
}

// SetSeverityLevelDebug defines the default severity level as debug
func SetSeverityLevelDebug() {
	severityLevelDefault = severityLevelDebug
}

// SetSeverityLevelInfo defines the default severity level as info
func SetSeverityLevelInfo() {
	severityLevelDefault = severityLevelInfo
}

// SetSeverityLevelWarning defines the default severity level as "warning"
func SetSeverityLevelWarning() {
	severityLevelDefault = severityLevelWarning
}

// SetSeverityLevelError defines the default severity level as "error"
func SetSeverityLevelError() {
	severityLevelDefault = severityLevelError
}

// SetSeverityLevelFatal defines the default severity level as "fatal"
func SetSeverityLevelFatal() {
	severityLevelDefault = severityLevelFatal
}

// SetOutputFormatJSON defines the format used for printing out messages as JSON
func SetOutputFormatJSON() {
	outputFormatDefault = outputFormatJSON
	outputTemplateDefault = outputTemplateJSON
}

// SetOutputFormatPlainText defines the format used for printing out messages as plain text
func SetOutputFormatPlainText() {
	outputFormatDefault = outputFormatPlain
	outputTemplateDefault = outputTemplatePlain
}

// EscapeMessageStringForJSON turns marshalling of message string on.
// It's not enabled by default because it impacts dramatically the performance
func EscapeMessageStringForJSON(b bool) {
	escapeMessageStringForJSON = b
}

func severityLevelAsString(sl byte) string {
	switch sl {
	case severityLevelDebug:
		return "debug"
	case severityLevelInfo:
		return "info"
	case severityLevelWarning:
		return "warning"
	case severityLevelError:
		return "error"
	case severityLevelFatal:
		return "critical"
	}

	return "info"
}

func write(level byte, msg string) {
	if msg == "" {
		return
	}

	if severityLevelDefault > level {
		return
	}

	if outputFormatDefault == outputFormatJSON && escapeMessageStringForJSON {
		if b, err := json.Marshal(msg); err != nil {
			fmt.Printf(outputTemplatePlain, time.Now().String(), severityLevelAsString(severityLevelError), "error "+err.Error()+" marshalling message: "+msg)
		} else {
			// Trim the beginning and trailing " character
			msg = string(b[1 : len(b)-1])
		}
	}

	fmt.Printf(outputTemplateDefault, time.Now().String(), severityLevelAsString(level), msg)
}

// Debug writes messages with severityLevel=debug
func Debug(msg string) {
	write(severityLevelDebug, msg)
}

// Debugf writes messages with severityLevel=debug, taking arguments in fmt.Printf format
func Debugf(format string, args ...interface{}) {
	write(severityLevelDebug, fmt.Sprintf(format, args...))
}

// Info writes messages with severityLevel=info
func Info(msg string) {
	write(severityLevelInfo, msg)
}

// Infof writes messages with severityLevel=info, taking arguments in fmt.Printf format
func Infof(format string, args ...interface{}) {
	write(severityLevelInfo, fmt.Sprintf(format, args...))
}

// Warn writes messages with severityLevel=warning
func Warn(msg string) {
	write(severityLevelWarning, msg)
}

// Warnf writes messages with severityLevel=warning, taking arguments in fmt.Printf format
func Warnf(format string, args ...interface{}) {
	write(severityLevelWarning, fmt.Sprintf(format, args...))
}

// Error writes messages with severityLevel=error
func Error(msg string) {
	write(severityLevelError, msg)
}

// ErrorErr writes messages with severityLevel=error, taking an argument of type error
func ErrorErr(err error) {
	write(severityLevelError, err.Error())
}

// ErrorErrReturn writes messages with severityLevel=error, taking arguments in fmt.Printf format
// It's provided as convenient one-liner return, for functions that returns an error
// Examples:
// if err != nil {
//     return logger.ErrorErrReturn(err)
// }
func ErrorErrReturn(err error) error {
	write(severityLevelError, err.Error())

	return err
}

// Errorf writes messages with severityLevel=error, taking arguments in fmt.Printf format
func Errorf(format string, args ...interface{}) {
	write(severityLevelError, fmt.Sprintf(format, args...))
}

// Fatal writes messages with severityLevel=fatal, and stop program with os.Exit(1)
func Fatal(msg string) {
	write(severityLevelFatal, msg)
	os.Exit(1)
}

// Fatalf writes messages with severityLevel=fatal, taking arguments in fmt.Printf format
// Fatalf stops the program with os.Exit(1)
func Fatalf(format string, args ...interface{}) {
	write(severityLevelFatal, fmt.Sprintf(format, args...))
	os.Exit(1)
}

// Global Setters
// Sets the values not only for current imported instance, but for the whole application and OS through environment variables

// SetGlobalSeverityLevelDebug  changes os environment variables in order to globally define the default severity level as debug
func SetGlobalSeverityLevelDebug() {
	if err := os.Setenv("LOG_SEVERITY_LEVEL", "debug"); err != nil {
		Errorf("chatty.SetGlobalSeverityLevelDebug error setting environment variable %s", err.Error())
	}

	SetSeverityLevelDebug()
}

// SetGlobalSeverityLevelInfo  changes os environment variables in order to globally define the default severity level as info
func SetGlobalSeverityLevelInfo() {
	if err := os.Setenv("LOG_SEVERITY_LEVEL", "info"); err != nil {
		Errorf("chatty.SetGlobalSeverityLevelInfo error setting environment variable %s", err.Error())
	}

	SetSeverityLevelInfo()
}

// SetGlobalSeverityLevelWarning  changes os environment variables in order to globally define the default severity level as "warning"
func SetGlobalSeverityLevelWarning() {
	if err := os.Setenv("LOG_SEVERITY_LEVEL", "warning"); err != nil {
		Errorf("chatty.SetGlobalSeverityLevelWarning error setting environment variable %s", err.Error())
	}

	SetSeverityLevelWarning()
}

// SetGlobalSeverityLevelError  changes os environment variables in order to globally define the default severity level as "error"
func SetGlobalSeverityLevelError() {
	if err := os.Setenv("LOG_SEVERITY_LEVEL", "error"); err != nil {
		Errorf("chatty.SetGlobalSeverityLevelError error setting environment variable %s", err.Error())
	}

	SetSeverityLevelError()
}

// SetGlobalSeverityLevelFatal  changes os environment variables in order to globally define the default severity level as "fatal"
func SetGlobalSeverityLevelFatal() {
	if err := os.Setenv("LOG_SEVERITY_LEVEL", "fatal"); err != nil {
		Errorf("chatty.SetGlobalSeverityLevelFatal error setting environment variable %s", err.Error())
	}

	SetSeverityLevelFatal()
}

// SetGlobalOutputFormatJSON changes os environment variables in order to globally define the format used for printing out messages as JSON
func SetGlobalOutputFormatJSON() {
	if err := os.Setenv("LOG_OUTPUT_FORMAT", outputFormatJSON); err != nil {
		Errorf("chatty.SetGlobalOutputFormatJSON error setting environment variable %s", err.Error())
	}

	SetOutputFormatJSON()
}

// SetGlobalOutputFormatPlainText changes os environment variables in order to globally define the format used for printing out messages as plain text
func SetGlobalOutputFormatPlainText() {
	if err := os.Setenv("LOG_OUTPUT_FORMAT", outputFormatPlain); err != nil {
		Errorf("chatty.SetGlobalOutputFormatPlainText error setting environment variable %s", err.Error())
	}

	SetOutputFormatPlainText()
}
