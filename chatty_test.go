package chatty

import (
	"fmt"
	"testing"
)

func TestAll(t *testing.T) {
	SetGlobalOutputFormat("plain")

	SetGlobalSeverityLevel("debug")

	Debug("this message appears if the level was defined as debug")

	const world = "world"

	Debugf("debug message is hello %s", world)

	Info("database connected")

	const userID = 10

	Infof("user %d changed its password", userID)

	// InfoKV writes messages including a sequence of key-value pairs, with severityLevel=info
	InfoKV("including bank branch", KVMap{"code": 1588, "branch": "münch", "address": nil})

	Warnf("blocking user %d for too many login attempts", userID)

	Error("could not connect external service xyz")

	err := fmt.Errorf("this is an example error")

	ErrorErr(err)

	// ErrorErrReturn writes messages with severityLevel=error, taking arguments in fmt.Printf format
	// It's provided as convenient one-liner return, for functions that returns an error
	// Example:
	// if err != nil {
	//     return logger.ErrorErrReturn(err)
	// }

	Errorf("error querying user %d investments: %v", userID, err)

	// ErrorKV writes messages including a sequence of key-value pairs, with severityLevel=error
	ErrorKV("error including bank branch", KVMap{"code": 1588, "branch": "münch", "error": err})

	// Fatal writes messages with severityLevel=fatal, and stop program with os.Exit(1)
	// chatty.Fatal("ooops! critical failure") {

	// Fatalf writes messages with severityLevel=fatal, taking arguments in fmt.Printf format
	// Fatalf stops the program with os.Exit(1)
	// chatty.Fatalf("this app is going to faint")

	// LastRecord returns the last recorded message.
	// It's designed for testing, but can also be used for sending the same message for two or more output engines
	fmt.Println(LastRecord())

	//FatalErr(err)
}
