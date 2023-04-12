package service

import (
	"context"
	"log"
	"mopoen-remake/config"
	"mopoen-remake/pkg/mail"
	"mopoen-remake/repository"
	"mopoen-remake/service/servicemodel"
	"testing"

	"github.com/stretchr/testify/assert"
)

var repo repository.Repository
var mailSnder mail.MailSender

func TestMain(m *testing.M) {
	env, err := config.New("../")
	if err != nil {
		log.Fatalln(err)
	}

	mailSnder = mail.NewMailSender(env.GmailUser, env.GmailPass)

	if err != nil {
		log.Fatal(err)
	}

	repo, err = repository.NewRepository(env.DBDriver, env.DBSource)
	if err != nil {
		log.Fatalln(err)
	}
	m.Run()
}

func TestDaftarMonitoring(t *testing.T) {
	mservice := NewMonitoringService(repo, mailSnder)
	err := mservice.DaftarMonitoring(context.TODO(), servicemodel.DaftarMonitoring{
		TipeSensor_id: 1,
		Location_id:   1,
		Email:         "somemail@gmail.com",
		Author:        "author",
		Nama:          "mr user",
		Keterangan:    "nganu kulo nuwun",
	})
	assert.NoError(t, err)
}
