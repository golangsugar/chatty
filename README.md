# Chatty ![Chatty](etc/chatty.png)
### Simple Lightweight Logger for Events
[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](http://golang.org)
&nbsp;
[![Go Report Card](https://goreportcard.com/badge/github.com/golangsugar/chatty)](https://goreportcard.com/report/github.com/golangsugar/chatty)
---
Chatty is a lightweight helper for events logging. <br />
It consists in a simple wrapper over native fmt methods. </br />
The OutputFormat can be defined with LOG_OUTPUT_FORMAT environment variable. </br />
The SeverityLevel can be defined with LOG_SEVERITY_LEVEL environment variable. </br />

#### Downloading
```bash
go get -u github.com/golangsugar/chatty
```

#### Using
Run the code below online at https://goplay.tools/snippet/tWKaK04eiPY

```go
package main

import (
	"fmt"
	"github.com/golangsugar/chatty"
)

func main() {
	chatty.SetGlobalOutputFormat("plain")

	chatty.SetGlobalSeverityLevel("debug")

	chatty.Debug("this message appears if the level was defined as debug")

	const world = "world"
	
	chatty.Debugf("debug message is hello %s", world)

	chatty.Info("database connected")

	const userID = 10
	
	chatty.Infof("user %d changed its password", userID)

	chatty.Warn("query took more than 10 seconds")

	chatty.Warnf("blocking user %d for too many login attempts", userID)

	chatty.Error("could not connect external service xyz")

	err := fmt.Errorf("this is an example error")
	
	chatty.ErrorErr(err)

	// ErrorErrReturn writes messages with severityLevel=error, taking arguments in fmt.Printf format
	// It's provided as convenient one-liner return, for functions that returns an error
	// Example:
	// if err != nil {
	//     return logger.ErrorErrReturn(err)
	// }
	
	chatty.Errorf("error querying user %d investments: %v", userID, err)
	
	// Fatal writes messages with severityLevel=fatal, and stop program with os.Exit(1)
	// chatty.Fatal("ooops! critical failure") {

	// Fatalf writes messages with severityLevel=fatal, taking arguments in fmt.Printf format
	// Fatalf stops the program with os.Exit(1)
	// chatty.Fatalf("this app is going to faint")

	// LastRecord returns the last recorded message.
	// It's designed for testing, but can also be used for sending the same message for two or more output engines
	fmt.Println(chatty.LastRecord())
    
	chatty.FatalErr(err)
}
```
```bash
2009-11-10 23:00:00 +0000 UTC m=+0.000000001	error	empty severity level
2009-11-10 23:00:00 +0000 UTC m=+0.000000001	debug	this message appears if the level was defined as debug
2009-11-10 23:00:00 +0000 UTC m=+0.000000001	debug	debug message is hello world
2009-11-10 23:00:00 +0000 UTC m=+0.000000001	info	database connected
2009-11-10 23:00:00 +0000 UTC m=+0.000000001	info	user 10 changed its password
2009-11-10 23:00:00 +0000 UTC m=+0.000000001	warning	query took more than 10 seconds
2009-11-10 23:00:00 +0000 UTC m=+0.000000001	warning	blocking user 10 for too many login attempts
2009-11-10 23:00:00 +0000 UTC m=+0.000000001	error	could not connect external service xyz
2009-11-10 23:00:00 +0000 UTC m=+0.000000001	error	this is an example error
2009-11-10 23:00:00 +0000 UTC m=+0.000000001	error	error querying user 10 investments: this is an example error
2009-11-10 23:00:00 +0000 UTC m=+0.000000001	error	error querying user 10 investments: this is an example error

2009-11-10 23:00:00 +0000 UTC m=+0.000000001	fatal	this is an example error
[T+0000ms]
Program exited.
```
