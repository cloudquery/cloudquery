package iam

import (
	"context"

	"github.com/apache/arrow/go/v14/arrow"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func userGroups() *schema.Table {
	tableName := "aws_iam_user_groups"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/IAM/latest/APIReference/API_Group.html`,
		Resolver:    fetchIamUserGroups,
		Transform:   transformers.TransformWithStruct(&types.Group{}, transformers.WithPrimaryKeys("Arn")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			{
				Name:       "user_arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.ParentColumnResolver("arn"),
				PrimaryKey: true,
			},
			{
				Name:     "user_id",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.ParentColumnResolver("user_id"),
			},
		},
	}
}
func fetchIamUserGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var config iam.ListGroupsForUserInput
	p := parent.Item.(*types.User)
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceIam).Iam
	config.UserName = p.UserName
	paginator := iam.NewListGroupsForUserPaginator(svc, &config)
	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx, func(options *iam.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- output.Groups
	}
	return nil
}
