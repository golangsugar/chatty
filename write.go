package chatty

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

var severityLevelMap = map[string]byte{
	severityLevelDebug:   0,
	severityLevelInfo:    1,
	severityLevelWarning: 2,
	severityLevelError:   3,
	severityLevelFatal:   4,
}

// Simplest error logger for avoiding recursive loops
func writeInternalError(now, level string, err error) {
	if outputFormatDefault == outputFormatJSON {
		fmt.Print(`{"ts":"` + now + `","level":"` + level + `","msg":"` + err.Error() + `"}`)
		return
	}

	fmt.Print(now + "\t" + level + "\t" + err.Error())
}

func jsonString(now, level, msg string, pairs ...logPair) string {
	tmp := `{"ts":"` + now + `","level":"` + level + `","msg":"` + msg + `"`

	if len(pairs) < 1 {
		return tmp + "}"
	}

	var buffer strings.Builder

	var value string

	for _, kv := range pairs {
		if s, ok := kv.Value.(string); ok {
			value = s
		} else {
			value = fmt.Sprintf("%v", kv.Value)
		}

		if _, err := buffer.WriteString(`,"` + kv.Key + `":"` + value + `"`); err != nil {
			writeInternalError(now, level, err)
		}
	}

	return tmp + buffer.String() + "}"
}

func jsonMarshalledString(now, level, msg string, pairs ...logPair) string {
	dic := map[string]interface{}{
		"ts":    now,
		"level": level,
		"msg":   msg,
	}

	if len(pairs) > 0 {
		details := make(map[string]interface{})

		for _, kv := range pairs {
			details[kv.Key] = kv.Value
		}

		dic["details"] = details
	}

	b, err := json.Marshal(dic)
	if err != nil {
		ErrorErr(err)
		// Fallback json made with string handling
		return jsonString(now, level, msg, pairs...)
	}

	return string(b)
}

func plainString(now, level, msg string, pairs ...logPair) string {
	tmp := now + "\t" + level + "\t" + msg

	if len(pairs) < 1 {
		return tmp + "\n"
	}

	var buffer strings.Builder
	var value string

	for _, kv := range pairs {
		if s, ok := kv.Value.(string); ok {
			value = s
		} else {
			value = fmt.Sprintf("%v", kv.Value)
		}

		if _, err := buffer.WriteString(",\t" + kv.Key + `=` + value); err != nil {
			writeInternalError(now, level, err)
		}
	}

	return buffer.String() + "\n"
}

func write(level string, msg string, pairs ...logPair) {
	if msg == "" && len(pairs) < 1 {
		return
	}

	if severityLevelMap[severityLevelDefault] > severityLevelMap[level] {
		return
	}

	now := time.Now().Format(time.RFC3339)

	record := ""

	if outputFormatDefault == outputFormatJSON {
		if !escapeInputForJSON {
			record = jsonString(now, level, msg, pairs...)
		} else {
			record = jsonMarshalledString(now, level, msg, pairs...)
		}
	} else {
		record = plainString(now, level, msg, pairs...)
	}

	instanceLastRecordMutex.Lock()
	instanceLastRecord = record
	instanceLastRecordMutex.Unlock()

	fmt.Print(instanceLastRecord)
}