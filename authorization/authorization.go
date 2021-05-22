package authorization

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/raulcodes/spotify-web-api/types"
)

// Client contains credentials for the Spotify Web API
type Client struct {
	ClientId  string
	ClientKey string
	URL       string
}

// ClientImpl implements functions for authorizing against and interacting with the Spotify Web API
type ClientImpl interface {
	AccessToken() (types.TokenResponse, error)
}

// NewClient returns an instance of Client, with provided client_id and client_key
func NewClient(clientId, clientKey string) Client {
	return Client{
		ClientId:  clientId,
		ClientKey: clientKey,
		URL:       "https://accounts.spotify.com/api/token",
	}
}

// AccessToken retrieves an access token from the `api/token` endpoint
func (c Client) AccessToken() (types.TokenResponse, error) {
	if c.ClientId == "" || c.ClientKey == "" {
		return types.TokenResponse{}, fmt.Errorf("AccessToken: Ensure that your clientID and clientSecret are set.")
	}

	body := url.Values{}
	body.Set("grant_type", "client_credentials")

	req, err := http.NewRequest(http.MethodPost, c.URL, strings.NewReader(body.Encode()))
	if err != nil {
		return types.TokenResponse{}, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", encodeAuth(c.ClientId, c.ClientKey))

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return types.TokenResponse{}, err
	}

	tokenResponse, err := parseTokenResponse(res)
	if err != nil {
		return types.TokenResponse{}, err
	}

	return tokenResponse, nil
}

func encodeAuth(id, key string) string {
	userPass := fmt.Sprintf("%s:%s", id, key)
	encodedStr := base64.StdEncoding.EncodeToString([]byte(userPass))

	return fmt.Sprintf("Basic %s", encodedStr)
}

func parseTokenResponse(res *http.Response) (types.TokenResponse, error) {
	body, _ := ioutil.ReadAll(res.Body)
	accessTokenRes := types.TokenResponse{}

	err := json.Unmarshal(body, &accessTokenRes)
	if err != nil {
		return types.TokenResponse{}, err
	}

	return accessTokenRes, nil
}
