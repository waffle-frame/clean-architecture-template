package server

import (
	"fmt"
	"net/http"

	"github.com/waffle-frame/clean-architecture-template/pkg/bootstrap/http/router"
	"github.com/waffle-frame/clean-architecture-template/pkg/config"
	"go.uber.org/fx"
)

// ServerModule ...
var ServerModule = fx.Provide(NewServer)

// Dependecies ...
type Dependecies struct {
	fx.In

	Config *config.Config
	Router *router.HTTPRouter
}

// NewServer ...
func NewServer(params Dependecies) *http.Server {
	url := fmt.Sprintf("%s:%s", params.Config.Server.Host, fmt.Sprint(params.Config.Server.Port))

	return &http.Server{
		Addr:    url,
		Handler: params.Router,
	}
}
