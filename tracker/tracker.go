package tracker

import (
	"io/ioutil"
	"net/http"
	"regexp"
)

type TrackResult struct {
	TargetURL         string
	TargetText        string
	StatusCode        int
	TargetTextMatched bool
}

func Perform(targetURL, targetText string) (TrackResult, error) {
	var trackResult = TrackResult{TargetURL: targetURL, TargetText: targetText}

	resp, err := http.Get(targetURL)
	if err != nil {
		return trackResult, err
	}

	trackResult.StatusCode = resp.StatusCode

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return trackResult, err
	}

	match, err := regexp.MatchString(targetText, string(body))
	if err != nil {
		return trackResult, err
	}

	trackResult.TargetTextMatched = match

	return trackResult, nil
}
