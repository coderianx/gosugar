package gosugar

import (
	"fmt"
	"os"
	"time"
)

// LogLevel defines the severity level of log messages
type LogLevel int

// Log level constants from lowest to highest severity
const (
	DebugLevel LogLevel = iota // Debug level: detailed diagnostic information
	InfoLevel                  // Info level: general informational messages
	WarnLevel                  // Warn level: warning messages for potentially harmful situations
	ErrorLevel                 // Error level: error messages for serious problems
	FatalLevel                 // Fatal level: fatal error messages that cause program exit
)

// Logger is a simple logging utility that manages log levels and outputs formatted log messages
type Logger struct {
	level LogLevel // Current logging level threshold
}

// NewLogger creates and returns a new Logger instance with the specified log level
func NewLogger(level LogLevel) *Logger {
	return &Logger{
		level: level,
	}
}

// SetLevel updates the logger's log level threshold
func (l *Logger) SetLevel(level LogLevel) {
	l.level = level
}

// log is an internal method that handles the actual logging logic.
// It checks if the message level meets the threshold, formats it with timestamp,
// prints to stdout, and exits if the level is Fatal
func (l *Logger) log(level LogLevel, label string, msg string) {
	if level < l.level {
		return // Skip logging if level is below threshold
	}

	timestamp := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("[%s] [%s] %s\n", timestamp, label, msg)

	if level == FatalLevel {
		os.Exit(1) // Exit program on fatal error
	}
}

// Debug logs a message at DEBUG level
func (l *Logger) Debug(msg string) {
	l.log(DebugLevel, "DEBUG", msg)
}

// Info logs a message at INFO level
func (l *Logger) Info(msg string) {
	l.log(InfoLevel, "INFO", msg)
}

// Warn logs a message at WARN level
func (l *Logger) Warn(msg string) {
	l.log(WarnLevel, "WARN", msg)
}

// Error logs a message at ERROR level
func (l *Logger) Error(msg string) {
	l.log(ErrorLevel, "ERROR", msg)
}

// Fatal logs a message at FATAL level and exits the program
func (l *Logger) Fatal(msg string) {
	l.log(FatalLevel, "FATAL", msg)
}
