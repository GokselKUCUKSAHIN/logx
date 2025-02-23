# logx

A structured and flexible logging library built on top of `logrus`, providing various log levels and JSON-based logging support.

## Installation

```sh
go get github.com/GokselKUCUKSAHIN/logx
```

## Usage

### Create a Logger

```go
import "github.com/GokselKUCUKSAHIN/logx"

log := logx.NewLogger("debug")
log.Info("Application started")
```

### Logging Levels

```go
log.Debug("This is a debug message")
log.Info("This is an info message")
log.Warn("This is a warning message")
log.Error("This is an error message")
```

### JSON Logging

```go
log.Printj(logger.JSON{"event": "user_login", "user": "john_doe"})
```

### Alert Logging

```go
log.Alert("Critical issue detected!")
log.Alertf("Service %s is down", "database")
```

### Panic Logging

```go
log.AlertPanic("Fatal error occurred")
```
