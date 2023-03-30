package iam

import (
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Groups() *schema.Table {
	tableName := "aws_iam_groups"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/IAM/latest/APIReference/API_Group.html`,
		Resolver:    fetchIamGroups,
		Transform:   transformers.TransformWithStruct(&types.Group{}, transformers.WithPrimaryKeys("Arn")),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "iam"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
		},

		Relations: []*schema.Table{
			groupAttachedPolicies(),
			groupPolicies(),
			groupLastAccessedDetails(),
		},
	}
}
