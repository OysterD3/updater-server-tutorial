package handler

import (
	"fmt"
	"github.com/OysterD3/updater-server-tutorial/env"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

func (h Handler) Download(c echo.Context) error {
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

	url := strings.Replace(p.APIUrl, "https://api.github.com", fmt.Sprintf("https://%s@api.github.com", env.Config.GitHub.AccessToken), 1)
	body, header, err := h.service.Download(url)

	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEOctetStream)
	c.Response().Header().Set(echo.HeaderContentLength, header.Get("Content-Length"))
	c.Response().Header().Set(echo.HeaderContentDisposition, header.Get("Content-Disposition"))
	c.Response().WriteHeader(http.StatusOK)

	defer body.Close()

	return c.Stream(http.StatusOK, "application/octet-stream", body)
}
