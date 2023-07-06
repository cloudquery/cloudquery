package quicksight

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/quicksight"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func groupMembers() *schema.Table {
	tableName := "aws_quicksight_group_members"
	return &schema.Table{
		Name:        tableName,
		Description: "https://docs.aws.amazon.com/quicksight/latest/APIReference/API_GroupMember.html",
		Resolver:    fetchQuicksightGroupMembers,
		Transform:   transformers.TransformWithStruct(&types.GroupMember{}, transformers.WithPrimaryKeys("Arn")),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "quicksight"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
			{
				Name:       "group_arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.ParentColumnResolver("arn"),
				PrimaryKey: true,
			},
		},
	}
}

func fetchQuicksightGroupMembers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	item := parent.Item.(types.Group)
	cl := meta.(*client.Client)
	svc := cl.Services().Quicksight

	input := quicksight.ListGroupMembershipsInput{
		AwsAccountId: aws.String(cl.AccountID),
		Namespace:    aws.String(defaultNamespace),
		GroupName:    item.GroupName,
	}
	// No paginator available
	for {
		out, err := svc.ListGroupMemberships(ctx, &input, func(options *quicksight.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- out.GroupMemberList

		if aws.ToString(out.NextToken) == "" {
			break
		}
		input.NextToken = out.NextToken
	}
	return nil
}
