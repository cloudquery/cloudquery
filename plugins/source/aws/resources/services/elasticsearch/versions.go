package elasticsearch

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Versions() *schema.Table {
	return &schema.Table{
		Name:        "aws_elasticsearch_versions",
		Description: `https://docs.aws.amazon.com/opensearch-service/latest/APIReference/API_ListVersions.html`,
		Resolver:    fetchElasticsearchVersions,
		Multiplex:   client.ServiceAccountRegionMultiplexer("es"),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "version",
				Type:     schema.TypeString,
				Resolver: resolveVersion,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "instance_types",
				Type:     schema.TypeJSON,
				Resolver: resolveInstanceTypes,
			},
		},
	}
}
