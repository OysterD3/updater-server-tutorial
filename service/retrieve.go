package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/OysterD3/updater-server-tutorial/env"
	"io/ioutil"
	"math"
	"net/http"
	"strings"
	"time"
)

type LatestRelease struct {
	Version string
	Notes string
	PublishDate time.Time
	Platforms map[string]Platform
}

type Platform struct {
	APIUrl string
	Filename string
	DownloadURL string
	ContentType string
	Size float64
}

type GitHubAuthor struct {
	Login string `json:"login"`
	ID uint64 `json:"id"`
	NodeID string `json:"node_id"`
	AvatarURL string `json:"avatar_url"`
	GravatarURL string `json:"gravatar_url"`
	URL string `json:"url"`
	HTMLUrl string `json:"html_url"`
	FollowersURL string `json:"followers_url"`
	FollowingURL string `json:"following_url"`
	GistsURL string `json:"gists_url"`
	StarredURL string `json:"starred_url"`
	SubscriptionsURL string `json:"subscriptions_url"`
	OrganizationsURL string `json:"organizations_url"`
	ReposURL string `json:"repos_url"`
	EventsURL string `json:"events_url"`
	ReceivedEventsURL string `json:"received_events_url"`
	Type string `json:"type"`
	SiteAdmin bool `json:"site_admin"`
}

type GitHubReleaseAsset struct {
	URL string `json:"url"`
	ID uint64 `json:"id"`
	Name string `json:"name"`
	NodeID string `json:"node_id"`
	Label string `json:"label"`
	Uploader GitHubAuthor `json:"uploader"`
	ContentType string `json:"content_type"`
	State string `json:"state"`
	Size uint64 `json:"size"`
	DownloadCount uint64 `json:"download_count"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	BrowserDownloadURL string `json:"browser_download_url"`
}

type GitHubRelease struct {
	URL string `json:"url"`
	AssetsURL string `json:"assets_url"`
	UploadURL string `json:"upload_url"`
	ID uint64 `json:"id"`
	Author GitHubAuthor `json:"author"`
	NodeID string `json:"node_id"`
	TagName string `json:"tag_name"`
	TargetCommitish string `json:"target_commitish"`
	Name string `json:"name"`
	Draft bool `json:"draft"`
	Prerelease bool `json:"prerelease"`
	CreatedAt time.Time `json:"created_at"`
	PublishedAt time.Time `json:"published_at"`
	Assets []GitHubReleaseAsset `json:"assets"`
	TarballURL string `json:"tarball_url"`
	ZipballURL string `json:"zipball_url"`
	Body string `json:"body"`
}


func (s *Service) RetrieveGitHubReleases(version string) (*LatestRelease, error) {
	repo := fmt.Sprintf("%s/%s", env.Config.GitHub.Account, env.Config.GitHub.Repository)
	releases := make([]GitHubRelease, 0)

	req, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("https://api.github.com/repos/%s/releases?per_page=100", repo),
		nil)

	req.Header.Add("Authorization", fmt.Sprintf("token %s", env.Config.GitHub.AccessToken))
	req.Header.Add("Accept", "application/json")

	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &releases)
	if err != nil {
		return nil, err
	}

	idx := 0

	if version != "latest" {
		for i, release := range releases {
			if release.TagName == version {
				idx = i
				break
			} else if i == len(releases)-1 {
				return nil, errors.New("version not found")
			}
		}
	}

	latest := new(LatestRelease)
	latest.Version = releases[idx].TagName
	latest.Notes = releases[idx].Body
	latest.PublishDate = releases[idx].PublishedAt
	latest.Platforms = make(map[string]Platform)

	for _, asset := range releases[idx].Assets {
		str := strings.Split(asset.Name, ".")
		ext := str[len(str)-1]
		platform := Platform{
			APIUrl: asset.URL,
			Filename: asset.Name,
			DownloadURL: asset.BrowserDownloadURL,
			ContentType: asset.ContentType,
			Size:        math.Round(float64(asset.Size / 1000000 * 10)) / 10,
		}

		if ext == "msi" {
			latest.Platforms["windows"] = platform
		} else if ext == "dmg" {
			latest.Platforms["mac"] = platform
		}
	}

	return latest, nil
}
