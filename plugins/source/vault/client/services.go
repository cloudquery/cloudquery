package client

import (
	internalAPI "github.com/cloudquery/plugins/vault/client/services/api"
	"github.com/hashicorp/vault/api"
)

type Services struct {
	Sys internalAPI.Sys
}

func NewServices(vc *api.Client) *Services {
	return &Services{Sys: vc.Sys()}
}
