# Chatty ![Chatty](etc/chatty.png)

### Simple Lightweight Logger for Events

[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](http://golang.org)
&nbsp;
[![Go Report Card](https://goreportcard.com/badge/github.com/golangsugar/chatty)](https://goreportcard.com/report/github.com/golangsugar/chatty)

---
Chatty is a lightweight helper for events logging. <br />
It consists in a simple wrapper over native fmt methods. <br />

#### Downloading

```console
go get -u github.com/golangsugar/chatty
```

#### Severity Levels

Chatty can handle **DEBUG**, **INFO**, **WARNING**, **ERROR** and **FATAL** severity levels

### Environment Variables

- The OutputFormat can be defined with `LOG_OUTPUT_FORMAT` environment variable;
- The SeverityLevel can be defined with `LOG_SEVERITY_LEVEL` environment variable.

#### Examples

🔗 Run the code below online at https://go.dev/play/p/la8C2jFZ9IA

```go
package main

import (
	"fmt"
	"github.com/golangsugar/chatty"
)

func main() {
	// Options
	// Chatty has local and global configuring routines.

	chatty.SetSeverityLevelDebug()

	chatty.SetOutputFormatPlainText()

	// Set the log output format, globally internally changing env vars
	//chatty.SetGlobalOutputFormat("plain")

	// Set the log minimum severity level, globally internally changing env vars
	//chatty.SetGlobalSeverityLevel("debug")

	// Highlights

	errx := fmt.Errorf("demonstration error")

	// You have a possible error that you need to log and nothing more.
	chatty.ConditionalErrorErr("this message and the error will be printed only if the error is not nil", errx)

	// You need to log the error and exit. You got a one-liner convenient helper. 
	go func() error {
		return chatty.ErrorErrReturn(errx)
	}()

	// InfoKV writes messages including a sequence of key-value pairs, with severityLevel=info
	chatty.InfoKV("including bank branch", chatty.KVMap{"code": 1588, "branch": "münch", "address": nil})

	// LastRecord returns the last recorded message.
	// It's designed for testing, but can also be used for sending the same message for two or more output engines
	fmt.Println(chatty.LastRecord())

	// General Examples:
	// For every severity level chatty enables at least 3 function signatures:
	// X(message), Xf(message, params...) and XKV(message, kvPairs)
	const world = "world"

	chatty.Debug("this message appears if the level was defined as debug")

	chatty.Debugf("debug message is 'hello %s'", world)

	// DebugKV writes messages including a sequence of key-value pairs, with severityLevel=debug
	chatty.InfoKV("server scheduled health-checker", chatty.KVMap{"foo": "bar", "answer": 42, "null": nil})

	chatty.Info("database connected")

	chatty.Warn("blocking user for too many login attempts")

	chatty.Error("could not connect external service xyz")

	chatty.ErrorErr(errx)

	// Fatal writes messages with severityLevel=fatal, and stop program with os.Exit(1)
	// chatty.Fatal("ooops! critical failure") {

	chatty.FatalErr(errx)
}
```

Outputs

```console
2009-11-10T23:00:00Z	warning	empty severity level. assuming default info
2009-11-10T23:00:00Z	error	this message and the error will be printed only if the error is not nil demonstration error
2009-11-10T23:00:00Z	info	including bank branch,	address=<nil>,	code=1588,	branch=münch
2009-11-10T23:00:00Z	info	including bank branch,	address=<nil>,	code=1588,	branch=münch

2009-11-10T23:00:00Z	debug	this message appears if the level was defined as debug
2009-11-10T23:00:00Z	debug	debug message is hello world
2009-11-10T23:00:00Z	info	server scheduled health-checker,	null=<nil>,	foo=bar,	answer=42
2009-11-10T23:00:00Z	info	database connected
2009-11-10T23:00:00Z	warning	blocking user for too many login attempts
2009-11-10T23:00:00Z	error	could not connect external service xyz
2009-11-10T23:00:00Z	error	demonstration error
2009-11-10T23:00:00Z	fatal	demonstration error

Program exited.
```
