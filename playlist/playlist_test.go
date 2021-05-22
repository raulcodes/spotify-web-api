package playlist_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/raulcodes/spotify-web-api/playlist"

	"github.com/stretchr/testify/assert"
)

var testCases = []struct{
	Name          string
	Server        func() *httptest.Server
	ExpectedError error
}{
	{
		"GetPlaylist: Happy Path, empty response",
		func() *httptest.Server {
			testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintln(w, "{}")
			}))
			return testServer
		},
		nil,
	},
	{
		"GetPlaylist: invalid response",
		func() *httptest.Server {
			testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintln(w, "invalid response")
			}))
			return testServer
		},
		fmt.Errorf("parsePlaylistResponse: Error encountered while trying to unmarshal response to PlaylistObj"),
	},
}

func TestGetPlaylist(t *testing.T) {
	for _, test := range testCases {
		t.Run(test.Name, func(t *testing.T) {
			testServer := test.Server()

			client := playlist.Client{
				Token: "test-token",
				URL:   testServer.URL,
			}

			_, err := client.GetPlaylist("test-id")

			if test.ExpectedError != nil {
				assert.EqualError(t, err, test.ExpectedError.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}