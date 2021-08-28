package handler

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/OysterD3/updater-server-tutorial/env"
	"github.com/labstack/echo/v4"
)

type VersionResponse struct {
	Filename    string    `json:"filename"`
	Version     string    `json:"version"`
	Notes       string    `json:"notes"`
	PublishDate time.Time `json:"publishDate"`
	DownloadURL string    `json:"downloadUrl"`
	Size        float64   `json:"size"`
}

func (h Handler) GetReleasesByPlatformAndVersion(c echo.Context) error {
	platform := strings.ToLower(strings.TrimSpace(c.Param("platform")))
	version := strings.ToLower(strings.TrimSpace(c.Param("version")))
	release, err := h.service.RetrieveGitHubReleases(version)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}
	p, ok := release.Platforms[platform]

	if !ok {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "platform not found",
		})
	}

	v := new(VersionResponse)
	v.Notes = release.Notes
	v.Version = release.Version
	v.PublishDate = release.PublishDate
	v.Filename = p.Filename
	v.DownloadURL = fmt.Sprintf("%s/download/%s/%s", env.Config.ServiceURL, platform, version)
	v.Size = p.Size

	return c.JSON(http.StatusOK, v)
}
