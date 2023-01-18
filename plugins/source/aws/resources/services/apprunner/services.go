package apprunner

import (
	"github.com/aws/aws-sdk-go-v2/service/apprunner/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Services() *schema.Table {
	return &schema.Table{
		Name:                "aws_apprunner_services",
		Description:         `https://docs.aws.amazon.com/apprunner/latest/api/API_Service.html`,
		Resolver:            fetchApprunnerServices,
		PreResourceResolver: getService,
		Multiplex:           client.ServiceAccountRegionMultiplexer("apprunner"),
		Transform:           transformers.TransformWithStruct(&types.Service{}),
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
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ServiceArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveApprunnerTags("ServiceArn"),
			},
		},
		Relations: []*schema.Table{
			Operations(),
			CustomDomains(),
		},
	}
}
