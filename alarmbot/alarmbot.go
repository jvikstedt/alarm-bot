package main

import (
	"fmt"
	"strings"

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
	fmt.Println("Starting AlarmBot")
	for _, c := range conf.TestObjects {
		t := tracker.NewTracker(c.Name, c.HistoryCount)
		trackResult, err := t.Perform(c.URL, c.MatchString, c.Status)
		if err != nil {
			if trackResult.Changed || len(t.TrackResults) <= 1 {
				fmt.Println(err)
				mail.Send("AlarmBot Error @ "+trackResult.TargetURL, err.Error(), c.MailTo)
			}
		} else {
			if trackResult.Changed {
				mail.Send("AlarmBot back to normal @ "+trackResult.TargetURL, strings.Join(trackResult.ChangeInfo, ","), c.MailTo)
			}
			fmt.Println(trackResult)
		}
		t.SaveHistory()
	}
	fmt.Println("Closing AlarmBot")
}

func setupConf() {
	confName := os.Getenv("ALARM_BOT_CONFIG")
	if confName == "" {
		confName = "../config.json"
	}
	conf = configuration.NewConfiguration(confName)
}

func setupMailer() {
	mail = mailer.NewMailer(conf.MailSetting.Host, conf.MailSetting.From, conf.MailSetting.Password, conf.MailSetting.Port)
}
