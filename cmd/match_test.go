package cmd

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/yucealiosman/riot-cli/pkg"
	"github.com/yucealiosman/riot-cli/riot"
)

type Select struct {
	RunMock func() (int, string, error)
}

func (s *Select) Run() (int, string, error) {
	return s.RunMock()
}

type HTTPClientMock struct {
	DoMock func(*http.Request) (*http.Response, error)
}

func (H HTTPClientMock) Do(r *http.Request) (*http.Response, error) {
	return H.DoMock(r)
}

type ConfigHandleMock struct {
	GetStringMock func(string) string
}

func (c *ConfigHandleMock) InitConfig() (err error) {
	return nil
}

func (c *ConfigHandleMock) WriteLocalConfigFile(key string, value string) {
}

func (c *ConfigHandleMock) GetString(key string) string {
	return c.GetStringMock(key)
}

func TestNormalizationFuncGlobalExistence(t *testing.T) {
	selectObj := &Select{}
	selectObj.RunMock = func() (int, string, error) {
		return 1, "ok", nil
	}
	regionPrompt = selectObj

	configHandle := &ConfigHandleMock{}

	configHandle.GetStringMock = func(key string) string {
		var keys = map[string]string{
			"region":    "testRegion",
			"riotToken": "testToken",
			"riotUrl":   "https://%s.api.riotgames.com/lol/",
		}
		value := keys[key]
		return value
	}

	configHandler = configHandle

	mockClient := HTTPClientMock{}

	matchDetail := riot.MatchDetail{
		PlatformID:   "TR1",
		GameID:       1145942841,
		QueueID:      420,
		SeasonID:     13,
		GameCreation: 22,
		GameDuration: 33,
	}

	b, err := json.Marshal(matchDetail)
	if err != nil {
		t.Errorf("An error occurred in the marshal process.")
		return
	}

	mockClient.DoMock = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			Body:       io.NopCloser(strings.NewReader(string(b))),
			StatusCode: 200,
		}, nil
	}
	client, _ = pkg.NewClient("http://testurl.com", mockClient)

	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	args := []string{"match", "-m", "555"}
	rootCmd.SetArgs(args)
	rootCmd.ExecuteC()
	result := buf.String()

	value, err := json.MarshalIndent(matchDetail, "", "  ")
	if err != nil {
		t.Errorf("An error occurred in the marshal process.")
		return
	}
	expectedValue := string(value)

	if expectedValue != result {
		t.Errorf("Output was incorrect, got: %s, expected: %s.", result, expectedValue)
	}

}
