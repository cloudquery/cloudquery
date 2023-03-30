package apprunner

import (
	"github.com/aws/aws-sdk-go-v2/service/apprunner/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func VpcConnectors() *schema.Table {
	tableName := "aws_apprunner_vpc_connectors"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/apprunner/latest/api/API_VpcConnector.html`,
		Resolver:    fetchApprunnerVpcConnectors,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "apprunner"),
		Transform:   client.TransformWithStruct(&types.VpcConnector{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("VpcConnectorArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveApprunnerTags("VpcConnectorArn"),
			},
		},
	}
}
