package iot

import (
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Policies() *schema.Table {
	tableName := "aws_iot_policies"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/iot/latest/apireference/API_Policy.html`,
		Resolver:    fetchIotPolicies,
		Transform:   transformers.TransformWithStruct(&types.Policy{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "iot"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: ResolveIotPolicyTags,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PolicyArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
