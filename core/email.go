package core

import (
	"fmt"
	"net/smtp"
	"time"
)

// generic function to send email
func SendEmail(from, to, subject, message string) error {
	mime := "Content-Type: text/plain; charset=\"UTF-8\"\r\nMIME-version: 1.0\r\n"
	subj := fmt.Sprintf("Subject: %s\r\n", subject)
	frm := fmt.Sprintf("From: %s\r\n", from)
	t := fmt.Sprintf("To: %s\r\n", to)
	date := fmt.Sprintf("Date: %s\r\n\r\n", time.Now().UTC().Format(time.RFC1123Z))
	body := mime + subj + frm + t + date + message

	return smtp.SendMail(
		Config.Email.Url,
		nil,
		from,
		[]string{to},
		[]byte(body),
	)
}
