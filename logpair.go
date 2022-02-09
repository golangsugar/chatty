package chatty

// logPair is a key-value pair for printing values in a structured way
type logPair struct {
	Key   string
	Value interface{}
}

// KVMap is a key-value pairs dictionary for printing values in a structured way
type KVMap map[string]interface{}

func kvMapAsLogPairs(kvm KVMap) []logPair {
	var pairs []logPair

	for k, v := range kvm {
		p := logPair{
			Key:   k,
			Value: v,
		}

		pairs = append(pairs, p)
	}

	return pairs
}

// DebugKV writes messages including a sequence of key-value pairs, with severityLevel=debug
func DebugKV(msg string, keyValuePairsMap KVMap) {
	write(severityLevelDebug, msg, kvMapAsLogPairs(keyValuePairsMap)...)
}

// InfoKV writes messages including a sequence of key-value pairs, with severityLevel=info
func InfoKV(msg string, keyValuePairsMap KVMap) {
	write(severityLevelInfo, msg, kvMapAsLogPairs(keyValuePairsMap)...)
}

// WarnKV writes messages including a sequence of key-value pairs, with severityLevel=warning
func WarnKV(msg string, keyValuePairsMap KVMap) {
	write(severityLevelWarning, msg, kvMapAsLogPairs(keyValuePairsMap)...)
}

// ErrorKV writes messages including a sequence of key-value pairs, with severityLevel=error
func ErrorKV(msg string, keyValuePairsMap KVMap) {
	write(severityLevelError, msg, kvMapAsLogPairs(keyValuePairsMap)...)
}

// FatalKV writes messages including a sequence of key-value pairs, with severityLevel=fatal
func FatalKV(msg string, keyValuePairsMap KVMap) {
	write(severityLevelFatal, msg, kvMapAsLogPairs(keyValuePairsMap)...)
}
