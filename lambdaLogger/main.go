package lambdaLogger

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

type Logger struct {
	infoLogger    *log.Logger
	warningLogger *log.Logger
	errorLogger   *log.Logger
	debugLogger   *log.Logger
}

// ToDo There must be a better way for injection and log levels
const (
	Info int = iota
	Warning
	Error
	Debug
)

var logger *Logger

// New creates a new Logger with the specified name and debug level.
// The output parameter determines where the logs will be written to.
//
// The name is included in every log message, allowing you to distinguish
// between logs from different parts of your application.
//
// If debug is set to true, the logger will log messages at the debug level.
// Otherwise, debug-level messages will be discarded.
//
// The output parameter should implement the io.Writer interface. This allows
// you to specify where the logs should be written to. This could be os.Stdout
// for console output, a file (e.g., os.File), or any other type that implements
// io.Writer. If you want to discard all logs, you can pass in io.Discard.
//
// Example usage:
//
//	// Log to the console:
//	l := logger.New("myapp", true, os.Stdout)
//	l.Print(logger.Info, "Hello, Info!")
//
//	// Log to a file:
//	file, _ := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
//	defer file.Close()
//	l = logger.New("myapp", true, file)
//	l.Print(logger.Info, "Hello, Info!")
//
//	// Discard all logs (useful in testing):
//	l = logger.New("myapp", false, io.Discard)
func New(name string, debug bool, output io.Writer) *Logger {
	logger = &Logger{
		infoLogger:    log.New(output, "INFO "+name, 0),
		warningLogger: log.New(output, "WARNING "+name, 0),
		errorLogger:   log.New(output, "ERROR "+name, 0),
		debugLogger:   log.New(output, "DEBUG "+name, 0),
	}

	if !debug {
		logger.debugLogger = log.New(io.Discard, "DEBUG "+name, 0)
	}

	return logger
}

func GetLogger() *Logger {
	if logger == nil {
		return New("unknown", true, os.Stdout)
	}

	return logger
}

func (l *Logger) Print(lvl int, v ...any) {
	switch lvl {
	case Info:
		l.infoLogger.Print(v...)
	case Warning:
		l.warningLogger.Print(v...)
	case Error:
		l.errorLogger.Print(v...)
	case Debug:
		l.debugLogger.Print(v...)
	}
}

func (l *Logger) Println(lvl int, v ...any) {
	switch lvl {
	case Info:
		l.infoLogger.Println(v...)
	case Warning:
		l.warningLogger.Println(v...)
	case Error:
		l.errorLogger.Println(v...)
	case Debug:
		l.debugLogger.Println(v...)
	}
}

func (l *Logger) Printf(lvl int, format string, v ...any) {
	l.Print(lvl, format, v)
}

func (l *Logger) PrintJson(lvl int, j []byte) {
	l.Println(lvl, string(j))
}

func (l *Logger) PrintStruct(lvl int, i interface{}) {
	logData, err := json.Marshal(i)
	if err != nil {
		l.Printf(lvl, "logger printStruct: marshal failed with:", err)
	}
	l.PrintJson(lvl, logData)
}
