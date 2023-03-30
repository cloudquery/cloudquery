package inspector2

import (
	"github.com/aws/aws-sdk-go-v2/service/inspector2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Findings() *schema.Table {
	tableName := "aws_inspector2_findings"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/inspector/v2/APIReference/API_Finding.html`,
		Resolver:    fetchInspector2Findings,
		Transform:   client.TransformWithStruct(&types.Finding{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "inspector2"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("FindingArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
