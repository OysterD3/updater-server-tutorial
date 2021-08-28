package service

import "net/http"

// Service :
type Service struct {
	client *http.Client
}

func New() *Service {
	return &Service{
		client: &http.Client{},
	}
}
