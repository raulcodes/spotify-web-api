package playlist

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/raulcodes/spotify-web-api/types"
)

type Client struct {
	Token string
	URL   string
}

func NewClient(token string) Client {
	return Client{
		Token: token,
		URL: "https://api.spotify.com/v1/playlists",
	}
}

type API interface {
	GetPlaylist(id string) (types.PlaylistObj, error) 
}

// GetPlaylist - Get a playlist owned by a Spotify user
func (c Client) GetPlaylist(id string) (types.PlaylistObj, error) {
	url := fmt.Sprintf("%s/%s", c.URL, id)

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
		return types.PlaylistObj{}, fmt.Errorf("parsePlaylistResponse: Error encountered while trying to unmarshal json to PlaylistObj")
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