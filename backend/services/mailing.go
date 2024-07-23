package services

import (
	"dot_conf/configs"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/smtp"
	"sync"
)

type IMailingService interface {
	SendMail(subject, content string, to ...string)
	address() string
}

type MailingService struct {
	host     string
	port     int
	password string
	from     string
}

var instance *MailingService
var once sync.Once

func newMailingService() *MailingService {
	config := configs.NewMailingConfig()
	return &MailingService{
		host:     config.Host,
		port:     config.Port,
		password: config.Password,
		from:     config.From,
	}
}

func GetMailingService() *MailingService {
	if instance == nil {
		once.Do(func() {
			instance = newMailingService()
		})
	}

	return instance
}

func (m *MailingService) SendMail(subject, content string, to ...string) error {
	message := fmt.Sprintf("Subject: %s\r\n\r\n%s", subject, content)
	auth := smtp.PlainAuth("", m.from, m.password, m.host)
	err := smtp.SendMail(m.address(), auth, m.from, to, []byte(message))

	if err != nil {
		log.Error("Error while sending email: ", err)
		return err
	}
	log.Info("Mail sent successfully")
	return nil
}

func (m *MailingService) address() string {
	return fmt.Sprintf("%s:%d", m.host, m.port)
}
