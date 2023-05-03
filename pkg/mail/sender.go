package mail

import (
	"fmt"
	"net/smtp"

	"github.com/google/uuid"
	"github.com/jordan-wright/email"
)

const SMPT_AUTH = "smtp.gmail.com"
const SMTP_SERVER = SMPT_AUTH + ":587"

type MailSender struct {
	smtpUser     string
	smtpPassword string
}

func NewMailSender(smtpUser string, smtpPassword string) MailSender {
	return MailSender{
		smtpUser:     smtpUser,
		smtpPassword: smtpPassword,
	}
}

// SendRegisteredMonitoring send email to registered mopoen user
// @to: email destination
// @id: mopoen id
// @author
// @secret: mopoen secret
func (s MailSender) SendRegisteredMonitoring(to string, id uuid.UUID, author string, secret string) error {
	// send email to registered mopoen
	body := fmt.Sprintf(
		"<h2>Berhasil daftar mopoen atas nama %s </h2>"+
			"Gunakan id berikut sebagai id monitoring mopoen: <b>%s</b><br>"+
			"Gunakan secret berikut sebagai secret monitoring mopoen: <b>%s</b><br>"+
			"keterangan lebih lanjut kunjungi <a href='https://github.com/arisygdc/mopoen-remake/blob/master/README.md'>API Documentation Mopoen</a>",
		author,
		secret,
		id.String(),
	)
	return s.Send(to, "Registered mopoen", body)
}

func (s MailSender) Send(to string, subject string, body string) error {
	e := email.Email{
		From:    fmt.Sprintf("Mopoen <%s>", s.smtpUser),
		To:      []string{to},
		Subject: subject,
	}
	e.HTML = []byte(body)
	return e.Send(SMTP_SERVER, s.plainAuth())
}

func (s MailSender) plainAuth() smtp.Auth {
	return smtp.PlainAuth("", s.smtpUser, s.smtpPassword, SMPT_AUTH)
}
