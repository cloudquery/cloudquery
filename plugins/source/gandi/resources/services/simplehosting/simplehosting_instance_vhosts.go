package simplehosting

import (
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/types"
)

func SimplehostingInstanceVhosts() *schema.Table {
	return &schema.Table{
		Name:     "gandi_simplehosting_instance_vhosts",
		Resolver: fetchSimplehostingInstanceVhosts,
		Columns: []schema.Column{
			{
				Name:       "instance_id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.ParentColumnResolver("id"),
				PrimaryKey: true,
			},
			{
				Name:     "created_at",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.PathResolver("CreatedAt"),
			},
			{
				Name:       "fqdn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("FQDN"),
				PrimaryKey: true,
			},
			{
				Name:     "is_a_test_vhost",
				Type:     arrow.FixedWidthTypes.Boolean,
				Resolver: schema.PathResolver("IsATestVhost"),
			},
			{
				Name:     "linked_dns_zone",
				Type:     types.ExtensionTypes.JSON,
				Resolver: schema.PathResolver("LinkedDNSZone"),
			},
			{
				Name:     "status",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.PathResolver("Status"),
			},
			{
				Name:     "application",
				Type:     types.ExtensionTypes.JSON,
				Resolver: schema.PathResolver("Application"),
			},
		},
	}
}
