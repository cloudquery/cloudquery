package zones

import (
	"github.com/apache/arrow/go/v14/arrow"
	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudquery/cloudquery/plugins/source/cloudflare/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func Zones() *schema.Table {
	return &schema.Table{
		Name:      "cloudflare_zones",
		Resolver:  fetchZones,
		Multiplex: client.AccountMultiplex,
		Transform: client.TransformWithStruct(&cloudflare.Zone{}),
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Type:        arrow.BinaryTypes.String,
				Resolver:    client.ResolveAccountID,
				Description: `The Account ID of the resource.`,
			},
			{
				Name:       "id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("ID"),
				PrimaryKey: true,
			},
		},
	}
}
