package ssm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Sessions() *schema.Table {
	tableName := "aws_ssm_sessions"
	return &schema.Table{
		Name: tableName,
		Description: `https://docs.aws.amazon.com/systems-manager/latest/APIReference/API_Session.html. 
Only Active sessions are fetched.`,
		Resolver:  fetchSsmSessions,
		Transform: transformers.TransformWithStruct(&types.Session{}, transformers.WithPrimaryKeys("SessionId")),
		Multiplex: client.ServiceAccountRegionMultiplexer(tableName, "ssm"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
		},
	}
}

func fetchSsmSessions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Ssm

	params := ssm.DescribeSessionsInput{
		State:   types.SessionStateActive,
		Filters: []types.SessionFilter{{Key: types.SessionFilterKeyOwner, Value: aws.String("Self")}},
	}
	paginator := ssm.NewDescribeSessionsPaginator(svc, &params)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(o *ssm.Options) {
			o.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.Sessions
	}
	return nil
}
