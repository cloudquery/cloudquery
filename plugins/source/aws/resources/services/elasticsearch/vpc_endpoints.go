// Code generated by codegen; DO NOT EDIT.

package elasticsearch

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func VpcEndpoints() *schema.Table {
	return &schema.Table{
		Name:        "aws_elasticsearch_vpc_endpoints",
		Description: `https://docs.aws.amazon.com/opensearch-service/latest/APIReference/API_VpcEndpoint.html`,
		Resolver:    fetchElasticsearchVpcEndpoints,
		Multiplex:   client.ServiceAccountRegionMultiplexer("es"),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
			{
				Name:     "domain_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DomainArn"),
			},
			{
				Name:     "endpoint",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Endpoint"),
			},
			{
				Name:     "status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Status"),
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("VpcEndpointId"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "vpc_endpoint_owner",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("VpcEndpointOwner"),
			},
			{
				Name:     "vpc_options",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("VpcOptions"),
			},
		},
	}
}
