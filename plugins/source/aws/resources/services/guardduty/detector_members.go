package guardduty

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/guardduty"
	"github.com/aws/aws-sdk-go-v2/service/guardduty/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/guardduty/models"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
)

func detectorMembers() *schema.Table {
	tableName := "aws_guardduty_detector_members"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/guardduty/latest/APIReference/API_Member.html`,
		Resolver:    fetchDetectorMembers,
		Transform: transformers.TransformWithStruct(&types.Member{},
			transformers.WithTypeTransformer(client.TimestampTypeTransformer),
			transformers.WithResolverTransformer(client.TimestampResolverTransformer),
		),
		Columns: []schema.Column{
			client.DefaultRegionColumn(false),
			{
				Name:     "detector_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
		},
	}
}

func fetchDetectorMembers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	detector := parent.Item.(*models.DetectorWrapper)
	c := meta.(*client.Client)
	svc := c.Services().Guardduty
	config := &guardduty.ListMembersInput{DetectorId: aws.String(detector.Id)}
	paginator := guardduty.NewListMembersPaginator(svc, config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- page.Members
	}
	return nil
}
