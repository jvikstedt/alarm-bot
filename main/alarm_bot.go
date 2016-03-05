package main

import (
	"fmt"

	"os"

	"github.com/jvikstedt/alarm-bot/configuration"
	"github.com/jvikstedt/alarm-bot/mailer"
	"github.com/jvikstedt/alarm-bot/tracker"
)

var mail *mailer.Mailer
var conf *configuration.Configuration

func init() {
	setupConf()
	setupMailer()
}

func main() {
	for _, c := range conf.TestObjects {
		trackResult, err := tracker.Perform(c.URL, c.MatchString, c.Status)
		if err != nil {
			fmt.Print(err)
		} else {
			fmt.Print(trackResult)
		}
	}
}

func setupConf() {
	confName := os.Getenv("ALARM_BOT_CONFIG")
	if confName == "" {
		confName = "./config.json"
	}
	conf = configuration.NewConfiguration(confName)
}

func setupMailer() {
	mail = mailer.NewMailer(conf.MailSetting.Host, conf.MailSetting.From, conf.MailSetting.Password, conf.MailSetting.Port)
}
