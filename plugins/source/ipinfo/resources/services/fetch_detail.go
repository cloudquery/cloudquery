package services

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/ipinfo/client"
	"github.com/cloudquery/cloudquery/plugins/source/ipinfo/internal/ipinfo"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func FetchIPInfo() *schema.Table {
	return &schema.Table{
		Name:      "ipinfo_details",
		Resolver:  fetchIPInfoDetails,
		Transform: transformers.TransformWithStruct(&ipinfo.IPinfo{}),
	}
}

func fetchIPInfoDetails(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	info, err := cl.IPInfo.GetIPInfo(cl.Spec.IP)
	if err != nil {
		return err
	}
	res <- info
	return nil
}
