package authorization_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/raulcodes/spotifyWebAPI/authorization"
	"github.com/raulcodes/spotifyWebAPI/types"

	"github.com/stretchr/testify/assert"
)

var AccessTokenTestCases = []struct {
	Name          string
	ClientID      string
	ClientKey     string
	Server        func() *httptest.Server
	ExpectedError error
	ExpectedResponse types.TokenResponse
}{
	{
		"Happy Path: AccessToken with valid creds, empty response",
		"test-id",
		"test-key",
		func() *httptest.Server {
			testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintln(w, "{}")
			}))
			return testServer
		},
		nil,
		types.TokenResponse{},
	},
	{
		"Happy Path: AccessToken with valid cred, valid response",
		"test-id",
		"test-key",
		func() *httptest.Server {
			testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintln(w, `
					{
						"access_token": "NgCXRKcMzYjw", 
						"token_type": "bearer", 
						"expires_in": 3600
					}`)
			}))
			return testServer
		},
		nil,
		types.TokenResponse{
			AccessToken: "NgCXRKcMzYjw",
			TokenType: "bearer",
			ExpiresIn: 3600,
		},
	},
	{
		"Error Path: AccessToken with empty clientID",
		"",
		"test-key",
		func() *httptest.Server {
			testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintln(w, "{}")
			}))
			return testServer
		},
		fmt.Errorf("AccessToken: Ensure that your clientID and clientSecret are set."),
		types.TokenResponse{},
	},
	{
		"Error Path: AccessToken with empty clientSecret",
		"test-id",
		"",
		func() *httptest.Server {
			testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintln(w, "{}")
			}))
			return testServer
		},
		fmt.Errorf("AccessToken: Ensure that your clientID and clientSecret are set."),
		types.TokenResponse{},
	},
}

func TestAccessToken(t *testing.T) {
	for _, test := range AccessTokenTestCases {
		t.Run(test.Name, func(t *testing.T) {
			authClient := authorization.Client{test.ClientID, test.ClientKey, test.Server().URL}
			token, err := authClient.AccessToken()
			
			if test.ExpectedError != nil {
				assert.EqualError(t, test.ExpectedError, err.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.ExpectedResponse, token)
			}
		})
	}
}
