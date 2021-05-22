package main_test

import (
	"testing"

	spotify "github.com/raulcodes/spotify-web-api"
)

func TestSetToken(t *testing.T) {
	client := spotify.NewClient("", "")

	client.SetToken("token")

	if client.Token != "token" {
		t.Errorf("Expected: %s, Actual: %s", "token", client.Token)
	}
}

var AccessTokenTestCases = []struct {
	Name          string
	ClientID      string
	ClientSecret  string
	ExpectedError error
}{
	{
		"Happy Path: AccessToken with valid creds",
		"",
		"",
		nil,
	},
}

func TestAccessToken(t *testing.T) {
	for _, test := range AccessTokenTestCases {
		t.Run(test.Name, func(t *testing.T) {
			if test.ExpectedError != nil {
				t.Fail()
			}
		})
	}
}
