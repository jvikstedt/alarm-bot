package main

import (
	"fmt"

	"github.com/jvikstedt/alarm-bot/configuration"
	"github.com/jvikstedt/alarm-bot/tracker"
)

var conf = configuration.NewConfiguration("./config.json")

func main() {
	for _, c := range conf.TestObjects {
		trackResult, _ := tracker.Perform(c.URL, c.MatchString)
		fmt.Print(trackResult)
	}
}
