package bootstrap

import (
	"context"
	"github.com/OysterD3/updater-server-tutorial/service"
)

// Bootstrap :
type Bootstrap struct {
	Service *service.Service
}

// New :
func New(ctx context.Context) *Bootstrap {
	bs := new(Bootstrap)
	bs.Service = service.New()

	return bs
}
