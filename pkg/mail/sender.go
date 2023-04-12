package mail

import (
	"fmt"
	"net/smtp"

	"github.com/google/uuid"
	"github.com/jordan-wright/email"
)

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

func (s MailSender) SendRegisteredMonitoring(to string, id uuid.UUID, author string) error {
	// send email to registered mopoen
	body := fmt.Sprintf(
		"<h2>Berhasil daftar mopoen atas nama %s </h2>"+
			"Gunakan id berikut sebagai id monitoring mopoen: <b>%s</b><br>"+
			"keterangan lebih lanjut kunjungi <a href='https://github.com/arisygdc/mopoen-remake/blob/master/README.md'>API Documentation Mopoen</a>",
		author,
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
	return e.Send("smtp.gmail.com:587", s.plainAuth())
}

func (s MailSender) plainAuth() smtp.Auth {
	return smtp.PlainAuth("", s.smtpUser, s.smtpPassword, "smtp.gmail.com")
}
