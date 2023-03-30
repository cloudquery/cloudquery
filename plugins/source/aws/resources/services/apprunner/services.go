package apprunner

import (
	"github.com/aws/aws-sdk-go-v2/service/apprunner/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Services() *schema.Table {
	tableName := "aws_apprunner_services"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/apprunner/latest/api/API_Service.html`,
		Resolver:            fetchApprunnerServices,
		PreResourceResolver: getService,
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "apprunner"),
		Transform:           client.TransformWithStruct(&types.Service{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
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
