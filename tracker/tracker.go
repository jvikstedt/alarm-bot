package tracker

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

type TrackResult struct {
	TargetURL         string
	TargetText        string
	TargetStatusCode  int
	ResultStatusCode  int
	ResultTextMatched bool
}

func Perform(targetURL, targetText string, targetStatusCode int) (TrackResult, error) {
	trackResult := TrackResult{targetURL, targetText, targetStatusCode, 0, false}

	resp, err := http.Get(targetURL)
	if err != nil {
		return trackResult, err
	}

	trackResult.ResultStatusCode = resp.StatusCode

	if trackResult.TargetStatusCode == 0 {
		// TODO Log that status comparison was skipped because it was not set
	} else if trackResult.ResultStatusCode != targetStatusCode {
		return trackResult, fmt.Errorf("StatusCodeMatchError: Looked for (%d), but found (%d)", trackResult.TargetStatusCode, trackResult.ResultStatusCode)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return trackResult, err
	}

	match, err := regexp.MatchString(targetText, string(body))
	if err != nil {
		return trackResult, err
	}

	trackResult.ResultTextMatched = match

	if !trackResult.ResultTextMatched {
		return trackResult, fmt.Errorf("TextMatchError: Looked for (%s)", trackResult.TargetText)
	}

	return trackResult, nil
}
