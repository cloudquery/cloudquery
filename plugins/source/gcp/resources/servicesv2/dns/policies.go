// Code generated by codegen; DO NOT EDIT.

package dns

import (
	"context"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/pkg/errors"
)

func Policies() *schema.Table {
	return &schema.Table{
		Name:      "gcp_dns_policies",
		Resolver:  fetchPolicies,
		Multiplex: client.ProjectMultiplex,
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
			},
			{
				Name:     "alternative_name_server_config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("AlternativeNameServerConfig"),
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
			},
			{
				Name:     "enable_inbound_forwarding",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("EnableInboundForwarding"),
			},
			{
				Name:     "enable_logging",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("EnableLogging"),
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
				Name:     "kind",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Kind"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "networks",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Networks"),
			},
		},
	}
}

func fetchPolicies(ctx context.Context, meta schema.ClientMeta, r *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	nextPageToken := ""
	for {
		output, err := c.Services.Dns.Policies.List(c.ProjectId).PageToken(nextPageToken).Do()
		if err != nil {
			return errors.WithStack(err)
		}
		res <- output.Policies

		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}
	return nil
}
