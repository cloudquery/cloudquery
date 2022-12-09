package resources

import (
	"github.com/tailscale/tailscale-client-go/tailscale"
)

func recipes() []*Resource {
	return []*Resource{
		{
			Struct: new(tailscale.ACL),
		},
	}
}
