package engine

import (
	"fmt"
	"net/smtp"
	"strings"
)

func (e *EmailEngine) SendMail(subject, body string) error {
	hp := strings.Split(e.config.UserConfig.EmailHost, ":")
	auth := smtp.PlainAuth(
		"",
		e.config.UserConfig.UserEmail,
		e.config.UserConfig.EmailPsw,
		hp[0],
	)
	contentType := "Content-Type: text/html; charset=UTF-8"
	user := e.config.UserConfig
	toAddresses := e.config.EmailTo.Addresses

	if len(toAddresses) == 0 {
		return fmt.Errorf("no email to")
	}
	to := toAddresses[0]
	for i := 1; i < len(toAddresses); i++ {
		to += ";" + toAddresses[i]
	}
	msg := []byte("To: " + to + "\r\nForm: " + user.UserEmail + ">\r\nSubject: " + subject + "\r\n" + contentType + "\r\n\r\n" + body)
	sentTo := strings.Split(to, ";")
	err := smtp.SendMail(user.EmailHost, auth, user.UserEmail, sentTo, msg)
	return err
}
