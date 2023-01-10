package simplehosting

import (
	"github.com/cloudquery/plugin-sdk/schema"
)

func SimplehostingInstanceVhosts() *schema.Table {
	return &schema.Table{
		Name:     "gandi_simplehosting_instance_vhosts",
		Resolver: fetchSimplehostingInstanceVhosts,
		Columns: []schema.Column{
			{
				Name:     "instance_id",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("id"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "created_at",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CreatedAt"),
			},
			{
				Name:     "fqdn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("FQDN"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "is_a_test_vhost",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("IsATestVhost"),
			},
			{
				Name:     "linked_dns_zone",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("LinkedDNSZone"),
			},
			{
				Name:     "status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Status"),
			},
			{
				Name:     "application",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Application"),
			},
		},
	}
}
