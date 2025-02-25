package logx

import (
	"io"
	"os"
	"strings"

	"github.com/GokselKUCUKSAHIN/jsonx"
	"github.com/sirupsen/logrus"
)

const (
	DEBUG Lvl = iota + 1
	INFO
	WARN
	ERROR
	OFF
)

type (
	ILogger interface {
		Output() io.Writer
		SetOutput(w io.Writer)
		Prefix() string
		SetPrefix(p string)
		Level() Lvl
		SetLevel(v Lvl)
		SetHeader(h string)
		Print(i ...interface{})
		Printf(format string, args ...interface{})
		Printj(j JSON)
		Debug(i ...interface{})
		Debugf(format string, args ...interface{})
		Debugj(j JSON)
		Info(i ...interface{})
		Infof(format string, args ...interface{})
		Infoj(j JSON)
		Warn(i ...interface{})
		Warnf(format string, args ...interface{})
		Warnj(j JSON)
		Error(i ...interface{})
		Errorf(format string, args ...interface{})
		Errorj(j JSON)
		Fatal(i ...interface{})
		Fatalj(j JSON)
		Fatalf(format string, args ...interface{})
		Panic(i ...interface{})
		Panicj(j JSON)
		Panicf(format string, args ...interface{})
	}
	Lvl  uint8
	JSON map[string]any
)

var logLevels = map[string]Lvl{
	"INFO":  INFO,
	"DEBUG": DEBUG,
	"WARN":  WARN,
	"ERROR": ERROR,
}

func retrieveLogLevel(levelName string) Lvl {
	if level, exist := logLevels[strings.ToUpper(levelName)]; exist {
		return level
	}
	return logLevels["INFO"]
}

type Logger struct {
	*logrus.Logger
}

var logger = &Logger{logrus.New()}

func NewLogger(level string) *Logger {
	logger = &Logger{logrus.New()}
	logger.SetOutput(os.Stdout)
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetLevel(retrieveLogLevel(level))
	return logger
}

func Print(i ...any) {
	logger.Print(i...)
}

func Printf(format string, args ...any) {
	logger.Printf(format, args...)
}

func Debug(i ...any) {
	logger.Debug(i...)
}

func Debugf(format string, args ...any) {
	logger.Debugf(format, args...)
}

func Info(i ...any) {
	logger.Info(i...)
}

func Infof(format string, args ...any) {
	logger.Infof(format, args...)
}

func Warn(i ...any) {
	logger.Warn(i...)
}

func Warnf(format string, args ...any) {
	logger.Warnf(format, args...)
}

func Error(i ...any) {
	logger.Error(i...)
}

func Errorf(format string, args ...any) {
	logger.Errorf(format, args...)
}

func Fatal(i ...any) {
	logger.Fatal(i...)
}

func Fatalf(format string, args ...any) {
	logger.Fatalf(format, args...)
}

func Panic(i ...any) {
	logger.Panic(i...)
}

func Panicf(format string, args ...any) {
	logger.Panicf(format, args...)
}

func Alert(i ...any) {
	logger.Alert(i...)
}

func Alertf(format string, args ...any) {
	logger.Alertf(format, args...)
}

func AlertPanic(i ...any) {
	logger.AlertPanic(i...)
}

func AlertPanicf(format string, args ...any) {
	logger.AlertPanicf(format, args...)
}

func toLogrusLevel(level Lvl) logrus.Level {
	switch level {
	case DEBUG:
		return logrus.DebugLevel
	case INFO:
		return logrus.InfoLevel
	case WARN:
		return logrus.WarnLevel
	case ERROR:
		return logrus.ErrorLevel
	}

	return logrus.InfoLevel
}

func toEchoLevel(level logrus.Level) Lvl {
	switch level {
	case logrus.DebugLevel:
		return DEBUG
	case logrus.InfoLevel:
		return INFO
	case logrus.WarnLevel:
		return WARN
	case logrus.ErrorLevel:
		return ERROR
	}

	return OFF
}

func (l *Logger) IsDebugLevel() bool {
	return l.GetLevel() == logrus.DebugLevel
}

func (l *Logger) Output() io.Writer {
	return l.Out
}

func (l *Logger) SetOutput(w io.Writer) {
	l.Out = w
}

func (l *Logger) Level() Lvl {
	return toEchoLevel(l.Logger.Level)
}

func (l *Logger) Lvl() string {
	return l.Logger.Level.String()
}

func (l *Logger) SetLevel(v Lvl) {
	l.Logger.Level = toLogrusLevel(v)
}

func (l *Logger) SetHeader(h string) {
	// do nothing
}

func (l *Logger) Formatter() logrus.Formatter {
	return l.Logger.Formatter
}

func (l *Logger) SetFormatter(formatter logrus.Formatter) {
	l.Logger.Formatter = formatter
}

func (l *Logger) Prefix() string {
	return ""
}

func (l *Logger) SetPrefix(p string) {
	// do nothing
}

func (l *Logger) Print(i ...any) {
	l.Logger.Print(i...)
}

func (l *Logger) Printf(format string, args ...any) {
	l.Logger.Printf(format, args...)
}

func (l *Logger) Println(i ...any) {
	l.Logger.Println(i...)
}

func (l *Logger) Printj(j JSON) {
	b, err := jsonx.Marshal(j)
	if err != nil {
		panic(err)
	}
	l.Logger.Println(string(b))
}

func (l *Logger) Debug(i ...any) {
	l.Logger.Debug(i...)
}

func (l *Logger) Debugf(format string, args ...any) {
	l.Logger.Debugf(format, args...)
}

func (l *Logger) Debugj(j JSON) {
	b, err := jsonx.Marshal(j)
	if err != nil {
		panic(err)
	}
	l.Logger.Debugln(string(b))
}

func (l *Logger) Info(i ...any) {
	l.Logger.Info(i...)
}

func (l *Logger) Infof(format string, args ...any) {
	l.Logger.Infof(format, args...)
}

func (l *Logger) Infoj(j JSON) {
	b, err := jsonx.Marshal(j)
	if err != nil {
		panic(err)
	}
	l.Logger.Infoln(string(b))
}

func (l *Logger) Warn(i ...any) {
	l.Logger.Warn(i...)
}

func (l *Logger) Warnf(format string, args ...any) {
	l.Logger.Warnf(format, args...)
}

func (l *Logger) Warnj(j JSON) {
	b, err := jsonx.Marshal(j)
	if err != nil {
		panic(err)
	}
	l.Logger.Warnln(string(b))
}

func (l *Logger) Error(i ...any) {
	l.Logger.Error(i...)
}

func (l *Logger) Errorf(format string, args ...any) {
	l.Logger.Errorf(format, args...)
}

func (l *Logger) Errorj(j JSON) {
	b, err := jsonx.Marshal(j)
	if err != nil {
		panic(err)
	}
	l.Logger.Errorln(string(b))
}

func (l *Logger) Fatal(i ...any) {
	l.Logger.Fatal(i...)
}

func (l *Logger) Fatalf(format string, args ...any) {
	l.Logger.Fatalf(format, args...)
}

func (l *Logger) Fatalj(j JSON) {
	b, err := jsonx.Marshal(j)
	if err != nil {
		panic(err)
	}
	l.Logger.Fatalln(string(b))
}

func (l *Logger) Panic(i ...any) {
	l.Logger.Panic(i...)
}

func (l *Logger) Panicf(format string, args ...any) {
	l.Logger.Panicf(format, args...)
}

func (l *Logger) Panicj(j JSON) {
	b, err := jsonx.Marshal(j)
	if err != nil {
		panic(err)
	}
	l.Logger.Panicln(string(b))
}

func (l *Logger) Alert(i ...any) {
	l.WithField("alert", "true").Error(i...)
}

func (l *Logger) Alertf(format string, args ...any) {
	l.WithField("alert", "true").Errorf(format, args...)
}

func (l *Logger) AlertPanic(i ...any) {
	l.WithField("alert", "true").Panic(i...)
}

func (l *Logger) AlertPanicf(format string, args ...any) {
	l.WithField("alert", "true").Panicf(format, args...)
}
