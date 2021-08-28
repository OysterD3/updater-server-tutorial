package service

import (
	"errors"
	"io"
	"net/http"
)

func (s *Service) Download(url string) (io.ReadCloser, http.Header, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Add("Accept", "application/octet-stream")
	resp, err := s.client.Do(req)
	req.Close = true

	if err != nil {
		return nil, nil, err
	}

	if resp.Body == nil {
		return nil, nil, errors.New("body is empty")
	}

	return resp.Body, resp.Header, nil
}
