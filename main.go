package main

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
	Token     string
}

// ClientImpl implements functions for authenticating against and interacting with the Spotify Web API
type ClientImpl interface {
	SetToken(string)
	AccessToken() (types.TokenResponse, error)

	GetPlaylist(id string)
}

// NewClient returns an instance of Client, with provided client_id and client_key
func NewClient(clientId, clientKey string) Client {
	return Client{
		ClientId:  clientId,
		ClientKey: clientKey,
		Token:     "",
	}
}

// SetToken sets the Token field of a Client instance
func (c *Client) SetToken(token string) {
	c.Token = token
}

// AccessToken retrieves an access token from the `api/token` endpoint
func (c Client) AccessToken() (types.TokenResponse, error) {
	body := url.Values{}
	body.Set("grant_type", "client_credentials")

	req, err := http.NewRequest(http.MethodPost, "https://accounts.spotify.com/api/token", strings.NewReader(body.Encode()))
	if err != nil {
		return types.TokenResponse{}, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", c.encodeAuth())

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return types.TokenResponse{}, err
	}

	tokenResponse, err := c.parseTokenResponse(res)
	if err != nil {
		return types.TokenResponse{}, err
	}

	return tokenResponse, nil
}

func (c Client) encodeAuth() string {
	userPass := fmt.Sprintf("%s:%s", c.ClientId, c.ClientKey)
	encodedStr := base64.StdEncoding.EncodeToString([]byte(userPass))

	return fmt.Sprintf("Basic %s", encodedStr)
}

func (c Client) parseTokenResponse(res *http.Response) (types.TokenResponse, error) {
	body, _ := ioutil.ReadAll(res.Body)
	accessTokenRes := types.TokenResponse{}

	err := json.Unmarshal(body, &accessTokenRes)
	if err != nil {
		return types.TokenResponse{}, err
	}

	return accessTokenRes, nil
}

func (c Client) GetPlaylist(id string) (types.PlaylistObj, error) {
	url := fmt.Sprintf("https://api.spotify.com/v1/playlists/%s", id)

	res, err := c.performRequest(http.MethodGet, url)
	if err != nil {
		return types.PlaylistObj{}, err
	}

	playlistObj, err := c.parsePlaylistResponse(res)
	if err != nil {
		return types.PlaylistObj{}, err
	}

	return playlistObj, nil
}

func (c Client) parsePlaylistResponse(res *http.Response) (types.PlaylistObj, error) {
	body, _ := ioutil.ReadAll(res.Body)
	playlistObj := types.PlaylistObj{}

	err := json.Unmarshal(body, &playlistObj)
	if err != nil {
		return types.PlaylistObj{}, err
	}

	return playlistObj, nil
}

func (c Client) performRequest(method, url string) (*http.Response, error) {
	if c.Token == "" {
		return nil, fmt.Errorf("performRequest: Missing token")
	}

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	bearer := fmt.Sprintf("Bearer %s", c.Token)
	req.Header.Set("Authorization", bearer)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}