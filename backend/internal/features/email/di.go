package email

import (
	"dbsystemdata-backend/internal/config"
	"dbsystemdata-backend/internal/util/logger"
)

var (
	env = config.GetEnv()
	log = logger.GetLogger()
)

var emailSMTPSender = &EmailSMTPSender{
	log,
	env.SMTPHost,
	env.SMTPPort,
	env.SMTPUser,
	env.SMTPPassword,
	env.SMTPFrom,
	env.SMTPHost != "" && env.SMTPPort != 0,
}

func GetEmailSMTPSender() *EmailSMTPSender {
	return emailSMTPSender
}
