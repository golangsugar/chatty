package chatty

// KVMap is a key-value pairs dictionary for printing values in a structured way
type KVMap map[string]interface{}

// DebugKV writes messages including a sequence of key-value pairs, with severityLevel=debug
func DebugKV(msg string, keyValuePairsMap KVMap) {
	write(severityLevelDebug, msg, keyValuePairsMap)
}

// InfoKV writes messages including a sequence of key-value pairs, with severityLevel=info
func InfoKV(msg string, keyValuePairsMap KVMap) {
	write(severityLevelInfo, msg, keyValuePairsMap)
}

// WarnKV writes messages including a sequence of key-value pairs, with severityLevel=warning
func WarnKV(msg string, keyValuePairsMap KVMap) {
	write(severityLevelWarning, msg, keyValuePairsMap)
}

// ErrorKV writes messages including a sequence of key-value pairs, with severityLevel=error
func ErrorKV(msg string, keyValuePairsMap KVMap) {
	write(severityLevelError, msg, keyValuePairsMap)
}

// FatalKV writes messages including a sequence of key-value pairs, with severityLevel=fatal
func FatalKV(msg string, keyValuePairsMap KVMap) {
	write(severityLevelFatal, msg, keyValuePairsMap)
}
