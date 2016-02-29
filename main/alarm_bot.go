package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"

	"github.com/jvikstedt/alarm-bot/configuration"
)

var conf = configuration.NewConfiguration("./config.json")

func main() {
	for _, c := range conf.TestObjects {
		matchInfo, err := perform(c)
		fmt.Print(matchInfo)
		fmt.Print(err)
	}
}

type MatchInfo struct {
	StatusCode int
	MatchFound bool
}

func perform(testObject configuration.TestObject) (MatchInfo, error) {
	matchInfo := MatchInfo{}

	resp, err := http.Get(testObject.URL)
	if err != nil {
		return matchInfo, err
	}

	matchInfo.StatusCode = resp.StatusCode

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return matchInfo, err
	}

	match, err := regexp.MatchString(testObject.MatchString, string(body))
	if err != nil {
		return matchInfo, err
	}

	matchInfo.MatchFound = match

	return matchInfo, nil
}
