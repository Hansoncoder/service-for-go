package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"sync"
	"time"
)

var (
	once     sync.Once
	instance *Logger
)

// Logger struct for handling log operations
type Logger struct {
	logFile   *os.File
	logWriter *log.Logger
	mutex     sync.Mutex
}

// GetLogger returns the singleton instance of Logger
func GetLogger() *Logger {
	once.Do(func() {
		instance = &Logger{}
		instance.initLogFile()
	})
	return instance
}

// initLogFile initializes the log file
func (l *Logger) initLogFile() {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	// Create logs directory if it doesn't exist
	logDir := "logs"
	if err := os.MkdirAll(logDir, 0755); err != nil {
		fmt.Println("Failed to create log directory:", err)
		return
	}

	// Generate log file path based on the current date
	logFileName := time.Now().Format("2006-01-02") + ".log"
	logFilePath := filepath.Join(logDir, logFileName)

	// Open the log file in append mode
	file, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Failed to open log file:", err)
		return
	}

	// Set log output
	l.logFile = file
	l.logWriter = log.New(file, "", 0) // Disable default flags to customize the format
}

// logToFile writes log messages to the file and console
func (l *Logger) logToFile(level, msg string) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	if l.logWriter == nil {
		fmt.Println("Logger is not initialized")
		return
	}

	// Get the caller's file name and line number
	_, file, line, ok := runtime.Caller(2) // 2 means the caller of the caller (e.g., logger.Info)
	if !ok {
		file = "unknown"
		line = 0
	}

	// Format the log message with the current time, file, line, and level
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	logMsg := fmt.Sprintf("[%s] [%s:%d] [%s] %s", currentTime, filepath.Base(file), line, level, msg)

	// Write to the log file
	l.logWriter.Println(logMsg)

	// Also print to the console
	fmt.Println(logMsg)
}

// Info logs an informational message
func (l *Logger) Info(msg string) {
	l.logToFile("INFO", msg)
}

// Infof logs a formatted informational message
func (l *Logger) Infof(format string, args ...interface{}) {
	l.logToFile("INFO", fmt.Sprintf(format, args...))
}

// Warn logs a warning message
func (l *Logger) Warn(msg string) {
	l.logToFile("WARN", msg)
}

// Warnf logs a formatted warning message
func (l *Logger) Warnf(format string, args ...interface{}) {
	l.logToFile("WARN", fmt.Sprintf(format, args...))
}

// Error logs an error message
func (l *Logger) Error(msg string) {
	l.logToFile("ERROR", msg)
}

// Errorf logs a formatted error message
func (l *Logger) Errorf(format string, args ...interface{}) {
	l.logToFile("ERROR", fmt.Sprintf(format, args...))
}

// Close closes the log file
func (l *Logger) Close() {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	if l.logFile != nil {
		l.logFile.Close()
	}
}
