package main

import (
	"email-send/engine"

	"github.com/pykelysia/pyketools"
)

func main() {
	e := engine.NewDefaultEmailEngine()
	subject := "Test Email"
	body := "<h1>test content.</h1>"
	err := e.SendMail(subject, body)
	if err != nil {
		pyketools.Fatalf("send email failed: %v", err)
	}
}
