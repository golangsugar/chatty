// Package chatty is a lightweight helper for events logging.
// It consists in a simple wrapper over native fmt methods.
// The OutputFormat can be defined with LOG_OUTPUT_FORMAT environment variable.
// The SeverityLevel can be defined with LOG_SEVERITY_LEVEL environment variable.
package chatty

import (
	"fmt"
	"os"
	"strings"
	"sync"
)

const (
	severityLevelDebug   = "debug"
	severityLevelInfo    = "info"
	severityLevelWarning = "warning"
	severityLevelError   = "error"
	severityLevelFatal   = "fatal"

	outputFormatJSON  = "json"
	outputFormatPlain = "plain"
)

var (
	severityLevelDefault    = severityLevelInfo
	outputFormatDefault     = outputFormatPlain
	escapeInputForJSON      = false
	instanceLastRecord      string
	instanceLastRecordMutex sync.Mutex
)

func init() {
	defineSeverityLevel()

	defineOutputFormat()
}

func guessSeverityLevel(sl string) (string, error) {
	if sl == "" {
		return severityLevelDefault, fmt.Errorf("empty severity level. assuming default %s", severityLevelDefault)
	}

	samePrefixOrEqual := func(mnemonic, s string) bool {
		if len(mnemonic) >= len(s) {
			return mnemonic[:len(s)] == s
		}

		return strings.HasPrefix(s, mnemonic)
	}

	x := strings.ToLower(strings.TrimSpace(sl))

	switch {
	case samePrefixOrEqual(severityLevelDebug, x), samePrefixOrEqual("verbose", x):
		return severityLevelDebug, nil
	case samePrefixOrEqual(severityLevelInfo, x), samePrefixOrEqual("normal", x):
		return severityLevelInfo, nil
	case samePrefixOrEqual(severityLevelWarning, x):
		return severityLevelWarning, nil
	case samePrefixOrEqual(severityLevelError, x):
		return severityLevelError, nil
	case samePrefixOrEqual(severityLevelFatal, x), samePrefixOrEqual("critical", x):
		return severityLevelFatal, nil
	}

	return severityLevelDefault, fmt.Errorf("unknown severity level %s", sl)
}

// EscapeInputForJSON turns marshalling of message string on.
// It's not enabled by default because it impacts dramatically the performance
func EscapeInputForJSON(b bool) {
	escapeInputForJSON = b
}

// Debug writes messages with severityLevel=debug
func Debug(msg string) {
	write(severityLevelDebug, msg, nil)
}

// Debugf writes messages with severityLevel=debug, taking arguments in fmt.Printf format
func Debugf(format string, args ...interface{}) {
	write(severityLevelDebug, fmt.Sprintf(format, args...), nil)
}

// Info writes messages with severityLevel=info
func Info(msg string) {
	write(severityLevelInfo, msg, nil)
}

// Infof writes messages with severityLevel=info, taking arguments in fmt.Printf format
func Infof(format string, args ...interface{}) {
	write(severityLevelInfo, fmt.Sprintf(format, args...), nil)
}

// Warn writes messages with severityLevel=warning
func Warn(msg string) {
	write(severityLevelWarning, msg, nil)
}

// Warnf writes messages with severityLevel=warning, taking arguments in fmt.Printf format
func Warnf(format string, args ...interface{}) {
	write(severityLevelWarning, fmt.Sprintf(format, args...), nil)
}

// Error writes messages with severityLevel=error
func Error(msg string) {
	write(severityLevelError, msg, nil)
}

// ErrorErr writes messages with severityLevel=error, taking an argument of type error
func ErrorErr(err error) {
	write(severityLevelError, err.Error(), nil)
}

// ConditionalErrorErr received a message and an error. It writes the content ONLY IF the error is not null.
// It sets severityLevel=error
func ConditionalErrorErr(msg string, err error) {
	if err != nil {
		write(severityLevelError, msg+" "+err.Error(), nil)
	}
}

// ErrorErrReturn writes messages with severityLevel=error, taking arguments in fmt.Printf format
// It's provided as convenient one-liner return, for functions that returns an error
// Examples:
// if err != nil {
//     return logger.ErrorErrReturn(err)
// }
func ErrorErrReturn(err error) error {
	write(severityLevelError, err.Error(), nil)

	return err
}

// Errorf writes messages with severityLevel=error, taking arguments in fmt.Printf format
func Errorf(format string, args ...interface{}) {
	write(severityLevelError, fmt.Sprintf(format, args...), nil)
}

// Fatal writes messages with severityLevel=fatal, and stop program with os.Exit(1)
func Fatal(msg string) {
	write(severityLevelFatal, msg, nil)
	os.Exit(1)
}

// FatalErr writes messages with severityLevel=fatal, and stop program with os.Exit(1)
func FatalErr(err error) {
	write(severityLevelFatal, err.Error(), nil)
	os.Exit(1)
}

// Fatalf writes messages with severityLevel=fatal, taking arguments in fmt.Printf format
// Fatalf stops the program with os.Exit(1)
func Fatalf(format string, args ...interface{}) {
	write(severityLevelFatal, fmt.Sprintf(format, args...), nil)
	os.Exit(1)
}

// LastRecord returns the last recorded message.
// It's designed for testing, but can also be used for sending the same message for two or more output engines
func LastRecord() string {
	return instanceLastRecord
}
