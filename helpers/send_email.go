package helpers

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
	"os"

	"github.com/bayek335/task-5-pbi-btpns-bayu-pamungkas/app"
)

func SendEmailVerifyCode(email string, code string) error {
	// Sender data.
	from := os.Getenv("EMAIL_SENDER")
	password := os.Getenv("EMAIL_SENDER_PASSWORD")
	// Receiver email address.
	to := []string{
		email,
	}
	// smtp server configuration.
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	t, _ := template.ParseFiles("./public/html/send_email_template.html")

	var body bytes.Buffer

	subject := "Email verification"

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("To: %s\r\nSubject: %s \n%s\n\n", to[0], subject, mimeHeaders)))

	t.Execute(&body, app.SendEmail{
		Email: email,
		Code:  code,
	})

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, body.Bytes())
	if err != nil {
		return err
	}
	return nil
}
