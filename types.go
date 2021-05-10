package spotify_web_api

import "time"

// TokenResponse represents a response from the `api/token` endpoint
type TokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

type PlaylistObj struct {
	Collaborative bool
	Description   string
	ExternalURLS  ExternalURLObj
	Followers     FollowersObj
	Href          string
	ID            string
	Images        []ImageObj
	Name          string
	Owner         PublicUserObj
	Public        bool
	SnapshotID    string
	Tracks        []PlaylistTrackObj
	Type          string
	URI           string
}

type ExternalURLObj struct {
	Spotify string `json:"spotify"`
}

type FollowersObj struct {
	Href  string `json:"href"`
	Total int    `json:"total"`
}

type ImageObj struct {
	Height int    `json:"height"`
	URL    string `json:"url"`
	Width  int    `json:"width"`
}

type PublicUserObj struct {
	DisplayName  string         `json:"display_name"`
	ExternalURLS ExternalURLObj `json:"external_urls"`
	Followers    FollowersObj   `json:"followers"`
	Href         string         `json:"string"`
	ID           string         `json:"id"`
	Images       []ImageObj     `json:"images"`
	Type         string         `json:"type"`
	URI          string         `json:"uri"`
}

type PlaylistTrackObj struct {
	AddedAt time.Time      `json:"added_at"`
	AddedBy PublicUserObj  `json:"added_by"`
	IsLocal bool           `json:"is_local"`
	Track   TrackOrEpisode `json:"track"`
}

type TrackObj struct {
	Album            AlbumObj            `json:"album"`
	Artists          []ArtistObj         `json:"artists"`
	AvailableMarkets []string            `json:"available_markets"`
	DiscNumber       int                 `json:"disc_number"`
	DurationMS       int                 `json:"duration_ms"`
	Explicit         bool                `json:"explicit"`
	ExternalIDs      ExternalIDObj       `json:"external_ids"`
	ExternalURLs     ExternalURLObj      `json:"external_urls"`
	Href             string              `json:"href"`
	ID               string              `json:"id"`
	IsLocal          bool                `json:"is_local"`
	IsPlayable       bool                `json:"is_playable"`
	Name             string              `json:"name"`
	Popularity       int                 `json:"popularity"`
	PreviewURL       string              `json:"preview_url"`
	Restrictions     TrackRestrictionObj `json:"restrictions"`
	TrackNumber      int                 `json:"track_number"`
	Type             string              `json:"type"`
	URI              string              `json:"uri"`
}

type TrackOrEpisode interface {
	IsTrack()
}

func (TrackObj) IsTrack() {}

type EpisodeObj struct {
	AudioPreviewURL      string                `json:"audio_preview_url"`
	Description          string                `json:"description"`
	DurationMS           int                   `json:"duration_ms"`
	Explicit             bool                  `json:"explicit"`
	ExternalURLs         ExternalURLObj        `json:"external_urls"`
	Href                 string                `json:"href"`
	HTMLDescription      string                `json:"html_description"`
	ID                   string                `json:"id"`
	Images               []ImageObj            `json:"images"`
	IsExternallyHosted   bool                  `json:"is_externally_hosted"`
	IsPlayable           bool                  `json:"is_playable"`
	Language             string                `json:"language"`
	Languages            []string              `json:"languages"`
	Name                 string                `json:"name"`
	ReleaseDate          string                `json:"release_date"`
	ReleaseDatePrecision string                `json:"release_date_precision"`
	Restrictions         EpisodeRestrictionObj `json:"restrictions"`
	ResumePoint          ResumePointObj        `json:"resume_point"`
	Show                 SimplifiedShowObj     `json:"show"`
	Type                 string                `json:"type"`
	URI                  string                `json:"uri"`
}

func (EpisodeObj) IsTrack() {}

type EpisodeRestrictionObj struct {
	Reason string `json:"reason"`
}

type ResumePointObj struct {
	FullyPlayed      bool `json:"fully_played"`
	ResumePositionMS int  `json:"resume_position_ms"`
}

type SimplifiedShowObj struct {
	AvailableMarkets   []string       `json:"available_markets"`
	Copyrights         []CopyrightObj `json:"copyrights"`
	Description        string         `json:"description"`
	Explicit           bool           `json:"explicit"`
	ExternalURLs       ExternalURLObj `json:"external_urls"`
	Href               string         `json:"href"`
	ID                 string         `json:"id"`
	Images             []ImageObj     `json:"images"`
	IsExternallyHosted bool           `json:"is_externally_hosted"`
	Languages          []string       `json:"languages"`
	MediaType          string         `json:"media_type"`
	Name               string         `json:"name"`
	Publisher          string         `json:"publisher"`
	Type               string         `json:"type"`
	URI                string         `json:"uri"`
}

type AlbumObj struct {
	AlbumType            string               `json:"album_type"`
	Artists              []ArtistObj          `json:"artists"`
	AvailableMarkets     []string             `json:"available_markets"`
	Copyrights           []CopyrightObj       `json:"copyrights"`
	ExternalIDs          ExternalIDObj        `json:"external_ids"`
	ExternalURLs         ExternalURLObj       `json:"external_urls"`
	Genres               []string             `json:"genres"`
	Href                 string               `json:"href"`
	ID                   string               `json:"id"`
	Images               []ImageObj           `json:"images"`
	Label                string               `json:"label"`
	Name                 string               `json:"name"`
	Popularity           int                  `json:"popularity"`
	ReleaseDate          string               `json:"release_date"`
	ReleaseDatePrecision string               `json:"release_date_precision"`
	Restrictions         AlbumRestrictionObj  `json:"restrictions"`
	Tracks               []SimplifiedTrackObj `json:"tracks"`
	Type                 string               `json:"type"`
	URI                  string               `json:"uri"`
}

type ArtistObj struct {
	ExternalIDs ExternalIDObj `json:"external_ids"`
	Followers   FollowersObj  `json:"followers"`
	Genres      []string      `json:"genres"`
	Href        string        `json:"href"`
	ID          string        `json:"id"`
	Images      []ImageObj    `json:"images"`
	Name        string        `json:"name"`
	Popularity  int           `json:"popularity"`
	Type        string        `json:"type"`
	URI         string        `json:"uri"`
}

type ExternalIDObj struct {
	EAN  string `json:"ean"`
	ISRC string `json:"isrc"`
	UPC  string `json:"upc"`
}

type CopyrightObj struct {
	Text string `json:"text"`
	Type string `json:"type"`
}

type AlbumRestrictionObj struct {
	Reason string `json:"reason"`
}

type TrackRestrictionObj struct {
	Reason string `json:"reason"`
}

type SimplifiedTrackObj struct {
	Artists          []SimplifiedArtistObj `json:"artists"`
	AvailableMarkets []string              `json:"available_markets"`
	DiscNumber       int                   `json:"disc_number"`
	DurationMS       int                   `json:"duration_ms"`
	Explicit         bool                  `json:"explicit"`
	ExternalURLs     ExternalURLObj        `json:"external_urls"`
	Href             string                `json:"href"`
	ID               string                `json:"id"`
	IsLocal          bool                  `json:"is_local"`
	IsPlayable       bool                  `json:"is_playable"`
	LinkedFrom       LinkedTrackObj        `json:"linked_from"`
	Name             string                `json:"name"`
	PreviewURL       string                `json:"preview_url"`
	Restrictions     TrackRestrictionObj   `json:"restrictions"`
	TrackNumber      int                   `json:"track_number"`
	Type             string                `json:"type"`
	URI              string                `json:"uri"`
}

type SimplifiedArtistObj struct {
	ExternalURLs ExternalURLObj `json:"external_urls"`
	Href         string         `json:"href"`
	ID           string         `json:"id"`
	Name         string         `json:"name"`
	Type         string         `json:"type"`
	URI          string         `json:"uri"`
}

type LinkedTrackObj struct {
	ExternalURLs ExternalURLObj `json:"external_urls"`
	Href         string         `json:"href"`
	ID           string         `json:"id"`
	Name         string         `json:"name"`
	Type         string         `json:"type"`
	URI          string         `json:"uri"`
}
