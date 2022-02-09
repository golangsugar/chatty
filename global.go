package chatty

import "os"

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

// SetGlobalSeverityLevel  changes os environment variables in order to globally define the default severity level
func SetGlobalSeverityLevel(sl string) {
	severityLevel, err := guessSeverityLevel(sl)
	if err != nil {
		ErrorErr(err)
	}

	switch severityLevel {
	case severityLevelDebug:
		SetSeverityLevelDebug()
	case severityLevelInfo:
		SetSeverityLevelInfo()
	case severityLevelWarning:
		SetSeverityLevelWarning()
	case severityLevelError:
		SetSeverityLevelError()
	case severityLevelFatal:
		SetSeverityLevelFatal()
	}
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

// SetGlobalOutputFormat changes os environment variables in order to globally define the format used for printing out messages
func SetGlobalOutputFormat(format string) {
	if format == outputFormatJSON {
		SetGlobalOutputFormatJSON()
	} else if format == outputFormatPlain {
		SetGlobalOutputFormatPlainText()
	} else {
		Errorf("unknown format given: %s", format)
	}
}
