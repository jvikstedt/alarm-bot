package main

import (
	"fmt"

	"os"

	"github.com/jvikstedt/alarm-bot/configuration"
	"github.com/jvikstedt/alarm-bot/tracker"
)

var conf *configuration.Configuration

func main() {
	setupConf()
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
