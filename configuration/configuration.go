package configuration

import (
	"encoding/json"
	"io/ioutil"
)

type Configuration struct {
	TestObjects []TestObject `json:"testObjects"`
}

type TestObject struct {
	URL         string `json:"url"`
	MatchString string `json:"matchString"`
	Status      int    `json:"status"`
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
