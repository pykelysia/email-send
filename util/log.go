package util

import (
	"email-send/config"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/pykelysia/pyketools"
)

type (
	Logger struct {
		c       config.Config
		logPath string
		logFile *os.File
		mu      sync.Mutex
	}
)

func NewLogger(c config.Config) *Logger {
	l := &Logger{c: c, logPath: c.LogConfig.LogPath}
	logFile, fileErr := os.OpenFile(l.logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if fileErr != nil {
		pyketools.Errorf("failed to open log file: %v", fileErr)
		logFile = nil
	}
	l.logFile = logFile
	return l
}

func (l *Logger) LogToFile(level, message string) {
	if l.logFile != nil {
		l.mu.Lock()
		defer l.mu.Unlock()
		timeStamp := time.Now().Format("2006-02-06 13:01:02")
		logLine := fmt.Sprintf("[%s] %s : %s\r\n", level, timeStamp, message)
		l.logFile.WriteString(logLine)
		l.logFile.Sync()
	}
}
