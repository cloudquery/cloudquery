package client

import (
	"context"
	"fmt"
	"net"
	"strings"

	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/spf13/cast"
	"github.com/thoas/go-funk"
)

func ResolveResourceTypeFromUrn(_ context.Context, meta schema.ClientMeta, r *schema.Resource, c schema.Column) error {
	value := funk.Get(r.Item, "URN", funk.WithAllowZero())
	urn := cast.ToString(value)
	if urn == "" {
		return nil
	}
	parts := strings.Split(urn, ":")
	return r.Set(c.Name, parts[1])
}

func ResolveResourceIdFromUrn(_ context.Context, meta schema.ClientMeta, r *schema.Resource, c schema.Column) error {
	value := funk.Get(r.Item, "URN", funk.WithAllowZero())
	urn := cast.ToString(value)
	if urn == "" {
		return nil
	}
	parts := strings.Split(urn, ":")
	if len(parts) < 2 {
		return nil
	}
	return r.Set(c.Name, parts[2])
}

func IPAddressResolver(path string) schema.ColumnResolver {
	return func(_ context.Context, meta schema.ClientMeta, r *schema.Resource, c schema.Column) error {
		ipStr, err := cast.ToStringE(funk.Get(r.Item, path, funk.WithAllowZero()))
		if err != nil {
			return err
		}
		ip := net.ParseIP(ipStr)
		if ipStr != "" && ip == nil {
			return fmt.Errorf("failed to parse IP from %s", ipStr)
		}
		if ip.To4() != nil {
			return r.Set(c.Name, ip.To4())
		}
		return r.Set(c.Name, ip)
	}
}
