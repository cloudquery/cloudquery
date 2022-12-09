package client

import (
	"context"

	"github.com/tailscale/tailscale-client-go/tailscale"
)

// Interface defines a read-only portion of tailscale.Client calls.
type Interface interface {
	// ACL retrieves the tailscale.ACL that is currently set for the given tailnet.
	ACL(ctx context.Context) (*tailscale.ACL, error)

	// DNSNameservers lists the DNS nameservers for a tailnet
	DNSNameservers(ctx context.Context) ([]string, error)

	// DNSPreferences retrieves the DNS preferences that are currently set for the given tailnet.
	// Supply the tailnet of interest in the path.
	DNSPreferences(ctx context.Context) (*tailscale.DNSPreferences, error)

	// DNSSearchPaths retrieves the list of search paths that is currently set for the given tailnet.
	DNSSearchPaths(ctx context.Context) ([]string, error)

	// DeviceSubnetRoutes Retrieves the list of subnet routes that a device is advertising, as well as those that are
	// enabled for it. Enabled routes are not necessarily advertised (e.g. for pre-enabling), and likewise, advertised
	// routes are not necessarily enabled.
	DeviceSubnetRoutes(ctx context.Context, deviceID string) (*tailscale.DeviceRoutes, error)

	// Devices lists the devices in a tailnet.
	Devices(ctx context.Context) ([]tailscale.Device, error)

	// GetKey returns all information on a key whose identifier matches the one provided.
	// This will not return the authentication key itself, just the metadata.
	GetKey(ctx context.Context, id string) (tailscale.Key, error)

	// Keys returns all keys within the tailnet.
	// The only fields set for each key will be its identifier.
	// The keys returned are relative to the user that owns the API key used to authenticate the client.
	Keys(ctx context.Context) ([]tailscale.Key, error)
}
