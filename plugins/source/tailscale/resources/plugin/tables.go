package plugin

import (
	"github.com/cloudquery/cloudquery/plugins/source/tailscale/resources/services/acl"
	"github.com/cloudquery/cloudquery/plugins/source/tailscale/resources/services/device"
	"github.com/cloudquery/cloudquery/plugins/source/tailscale/resources/services/dns"
	"github.com/cloudquery/cloudquery/plugins/source/tailscale/resources/services/key"
	"github.com/cloudquery/plugin-sdk/schema"
)

func tables() []*schema.Table {
	return []*schema.Table{
		acl.Acls(),
		device.Devices(),
		dns.Nameservers(),
		dns.Preferences(),
		dns.Searchpaths(),
		key.Keys(),
	}
}
