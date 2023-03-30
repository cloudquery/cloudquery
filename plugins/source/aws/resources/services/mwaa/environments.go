package mwaa

import (
	"github.com/aws/aws-sdk-go-v2/service/mwaa/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Environments() *schema.Table {
	tableName := "aws_mwaa_environments"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/mwaa/latest/API/API_Environment.html`,
		Resolver:            fetchMwaaEnvironments,
		Transform:           client.TransformWithStruct(&types.Environment{}),
		PreResourceResolver: getEnvironment,
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "airflow"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Arn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
