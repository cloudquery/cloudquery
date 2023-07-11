package detective

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/detective"
	"github.com/aws/aws-sdk-go-v2/service/detective/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func members() *schema.Table {
	tableName := "aws_detective_graph_members"
	return &schema.Table{
		Name: tableName,
		Description: `https://docs.aws.amazon.com/detective/latest/APIReference/API_GetMembers.html
The 'request_account_id' and 'request_region' columns are added to show the account and region of where the request was made from.`,
		Resolver:  fetchMembers,
		Multiplex: client.ServiceAccountRegionMultiplexer(tableName, "api.detective"),
		Transform: transformers.TransformWithStruct(&types.MemberDetail{}, transformers.WithPrimaryKeys("AccountId", "GraphArn")),
		Columns: []schema.Column{
			{
				Name:     "request_account_id",
				Type:     arrow.BinaryTypes.String,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:       "request_region",
				Type:       arrow.BinaryTypes.String,
				Resolver:   client.ResolveAWSRegion,
				PrimaryKey: true,
			},
		},
	}
}

func fetchMembers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Detective
	graph := parent.Item.(types.Graph)
	config := detective.ListMembersInput{
		GraphArn:   graph.Arn,
		MaxResults: aws.Int32(50),
	}
	paginator := detective.NewListMembersPaginator(svc, &config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *detective.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}

		res <- page.MemberDetails
	}

	return nil
}
