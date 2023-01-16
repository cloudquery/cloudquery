package glue

import (
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func SecurityConfigurations() *schema.Table {
	return &schema.Table{
		Name:        "aws_glue_security_configurations",
		Description: `https://docs.aws.amazon.com/glue/latest/webapi/API_SecurityConfiguration.html`,
		Resolver:    fetchGlueSecurityConfigurations,
		Transform:   transformers.TransformWithStruct(&types.SecurityConfiguration{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer("glue"),
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
				Name: "name",
				Type: schema.TypeString,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
