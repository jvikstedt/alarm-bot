package mailer

import "net/smtp"

type Mailer struct {
	Host     string `json:"host"`
	From     string `json:"from"`
	Password string `json:"password"`
	Port     string `json:"port"`
}

func NewMailer(host, from, password, port string) *Mailer {
	return &Mailer{host, from, password, port}
}

func (m Mailer) Send(subject, body, to string) error {
	msg := []byte("To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" + body + "\r\n")

	err := smtp.SendMail(m.Host+":"+m.Port, smtp.PlainAuth("", m.From, m.Password, m.Host), m.From, []string{to}, msg)
	return err
}
