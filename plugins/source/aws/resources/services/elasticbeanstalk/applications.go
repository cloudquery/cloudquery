package elasticbeanstalk

import (
	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Applications() *schema.Table {
	tableName := "aws_elasticbeanstalk_applications"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/elasticbeanstalk/latest/api/API_ApplicationDescription.html`,
		Resolver:    fetchElasticbeanstalkApplications,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "elasticbeanstalk"),
		Transform:   client.TransformWithStruct(&types.ApplicationDescription{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ApplicationArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name: "date_created",
				Type: schema.TypeTimestamp,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveElasticbeanstalkApplicationTags,
			},
		},
	}
}
