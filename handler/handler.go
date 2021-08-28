package handler

import (
	"github.com/OysterD3/updater-server-tutorial/bootstrap"
	"github.com/OysterD3/updater-server-tutorial/service"
)

// Handler :
type Handler struct {
	service *service.Service
}

// New :
func New(bs *bootstrap.Bootstrap) *Handler {
	return &Handler{
		service: bs.Service,
	}
}
