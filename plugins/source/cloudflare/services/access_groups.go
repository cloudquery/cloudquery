package services

import (
	"context"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudquery/cloudquery/plugins/source/cloudflare/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func AccessGroups() *schema.Table {
	return &schema.Table{
		Name:        "cloudflare_access_groups",
		Description: "AccessGroup defines a group for allowing or disallowing access to one or more Access applications.",
		Resolver:    fetchAccessGroups,
		Multiplex:   client.ZoneMultiplex,
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The Account ID of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAccountId,
			},
			{
				Name:        "zone_id",
				Description: "Zone identifier tag.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveZoneId,
			},
			{
				Name:            "id",
				Description:     "The unique identifier for the Access group.",
				Type:            schema.TypeString,
				Resolver:        schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:        "created_at",
				Description: "Hashed script content, can be used in a If-None-Match header when updating.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "updated_at",
				Description: "Size of the script, in bytes.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "name",
				Description: "The name of the Access group.",
				Type:        schema.TypeString,
			},
			{
				Name:        "include",
				Description: "Rules evaluated with an OR logical operator. A user needs to meet only one of the Include rules.",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "exclude",
				Description: "Rules evaluated with a NOT logical operator. To match a policy, a user cannot meet any of the Exclude rules.",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "require",
				Description: "Rules evaluated with an AND logical operator. To match a policy, a user must meet all of the Require rules.",
				Type:        schema.TypeJSON,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchAccessGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client)
	zoneID := svc.ZoneId

	pagination := cloudflare.PaginationOptions{
		Page:    1,
		PerPage: client.MaxItemsPerPage,
	}

	for {
		resp, info, err := svc.ClientApi.ZoneLevelAccessGroups(ctx, zoneID, pagination)
		if err != nil {
			return err
		}
		res <- resp

		if !info.HasMorePages() {
			break
		}
		pagination.Page++
	}
	return nil
}
