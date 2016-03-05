package configuration

import (
	"encoding/json"
	"io/ioutil"
)

type Configuration struct {
	TestObjects []TestObject `json:"testObjects"`
	MailSetting Mail         `json:"mailSetting"`
}

type TestObject struct {
	URL         string `json:"url"`
	MatchString string `json:"matchString"`
	Status      int    `json:"status"`
}

type Mail struct {
	Host     string `json:"host"`
	From     string `json:"from"`
	Password string `json:"password"`
	Port     string `json:"port"`
}

func NewConfiguration(filePath string) *Configuration {
	var conf Configuration
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(file, &conf)
	return &conf
}
