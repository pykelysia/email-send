package util

import (
	"email-send/config"
	"testing"
)

func TestLogEngine(t *testing.T) {
	c := config.LoadConfig("../log_test.yaml")
	l := NewLogger(c)
	l.LogToFile("Info", "Test success.")
}
