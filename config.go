package chatty

import (
	"os"
	"strings"
)

func defineSeverityLevel() {
	var err error

	if severityLevelDefault, err = guessSeverityLevel(os.Getenv("LOG_SEVERITY_LEVEL")); err != nil {
		ErrorErr(err)
	}
}

func defineOutputFormat() {
	lof := strings.ToLower(strings.TrimSpace(os.Getenv("LOG_OUTPUT_FORMAT")))

	if lof == outputFormatJSON {
		outputFormatDefault = outputFormatJSON
	} else {
		outputFormatDefault = outputFormatPlain
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
}

// SetOutputFormatPlainText defines the format used for printing out messages as plain text
func SetOutputFormatPlainText() {
	outputFormatDefault = outputFormatPlain
}
