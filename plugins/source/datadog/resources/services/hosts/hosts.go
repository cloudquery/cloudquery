// Code generated by codegen; DO NOT EDIT.

package hosts

import (
	"github.com/cloudquery/cloudquery/plugins/source/datadog/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Hosts() *schema.Table {
	return &schema.Table{
		Name:      "datadog_hosts",
		Resolver:  fetchHosts,
		Multiplex: client.AccountMultiplex,
		Columns: []schema.Column{
			{
				Name:     "account_name",
				Type:     schema.TypeString,
				Resolver: client.ResolveAccountName,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "id",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Id"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "aliases",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("Aliases"),
			},
			{
				Name:     "apps",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("Apps"),
			},
			{
				Name:     "aws_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AwsName"),
			},
			{
				Name:     "host_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("HostName"),
			},
			{
				Name:     "is_muted",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("IsMuted"),
			},
			{
				Name:     "last_reported_time",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("LastReportedTime"),
			},
			{
				Name:     "meta",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Meta"),
			},
			{
				Name:     "metrics",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Metrics"),
			},
			{
				Name:     "mute_timeout",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("MuteTimeout"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "sources",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("Sources"),
			},
			{
				Name:     "tags_by_source",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("TagsBySource"),
			},
			{
				Name:     "up",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Up"),
			},
		},
	}
}
