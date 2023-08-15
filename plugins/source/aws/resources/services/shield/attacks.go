package shield

import (
	"context"
	"time"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/service/shield"
	"github.com/aws/aws-sdk-go-v2/service/shield/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Attacks() *schema.Table {
	tableName := "aws_shield_attacks"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/waf/latest/DDOSAPIReference/API_AttackDetail.html`,
		Resolver:            fetchShieldAttacks,
		PreResourceResolver: getAttack,
		Transform:           transformers.TransformWithStruct(&types.AttackDetail{}),
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "shield"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			{
				Name:        "id",
				Type:        arrow.BinaryTypes.String,
				Resolver:    schema.PathResolver("AttackId"),
				Description: `The unique identifier (ID) of the attack`,
				PrimaryKey:  true,
			},
		},
	}
}

func fetchShieldAttacks(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Shield
	end := time.Now()
	start := end.Add(-time.Hour * 24)
	config := shield.ListAttacksInput{
		EndTime:   &types.TimeRange{ToExclusive: &end},
		StartTime: &types.TimeRange{FromInclusive: &start},
	}
	pagintor := shield.NewListAttacksPaginator(svc, &config)
	for pagintor.HasMorePages() {
		page, err := pagintor.NextPage(ctx, func(o *shield.Options) {
			o.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.AttackSummaries
	}
	return nil
}

func getAttack(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Shield
	a := resource.Item.(types.AttackSummary)

	attack, err := svc.DescribeAttack(ctx, &shield.DescribeAttackInput{AttackId: a.AttackId}, func(o *shield.Options) {
		o.Region = cl.Region
	})
	if err != nil {
		return err
	}

	resource.Item = attack.Attack
	return nil
}
