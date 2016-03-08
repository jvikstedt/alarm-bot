package tracker

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"time"
)

const TrackResultsPath = "../storage/track_results"

type TrackResult struct {
	TargetURL         string
	TargetText        string
	TargetStatusCode  int
	ResultStatusCode  int
	ResultTextMatched bool
	Timestamp         time.Time
}

type Tracker struct {
	Name         string
	SaveCount    int
	TrackResults []TrackResult
}

func NewTracker(name string, saveCount int) *Tracker {
	trackResults := loadHistory(name)
	return &Tracker{name, saveCount, trackResults}
}

func init() {
	if _, err := os.Stat(TrackResultsPath); os.IsNotExist(err) {
		os.MkdirAll(TrackResultsPath, 0777)
	}
}

func (t *Tracker) SaveHistory() {
	firstIndex := len(t.TrackResults) - t.SaveCount
	if firstIndex < 0 {
		firstIndex = 0
	}
	b, err := json.Marshal(t.TrackResults[firstIndex:])
	if err != nil {
		fmt.Println(err)
	}
	err = ioutil.WriteFile(TrackResultsPath+"/"+t.Name+".json", b, 0777)
	if err != nil {
		fmt.Println(err)
	}
}

func (t *Tracker) Perform(targetURL, targetText string, targetStatusCode int) (TrackResult, error) {
	trackResult := TrackResult{targetURL, targetText, targetStatusCode, 0, false, time.Now()}
	t.TrackResults = append(t.TrackResults, trackResult)

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

func loadHistory(name string) []TrackResult {
	var trackResults []TrackResult
	file, err := ioutil.ReadFile(TrackResultsPath + "/" + name + ".json")
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(file, &trackResults)
	if err != nil {
		fmt.Println(err)
	}

	return trackResults
}
