package ssm

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func PatchBaselines() *schema.Table {
	tableName := "aws_ssm_patch_baselines"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/systems-manager/latest/APIReference/API_PatchBaselineIdentity.html`,
		Resolver:    fetchSsmPatchBaselines,
		Transform:   transformers.TransformWithStruct(&types.PatchBaselineIdentity{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "ssm"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
			{
				Name:       "baseline_id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("BaselineId"),
				PrimaryKey: true,
			},
		},
	}
}

func fetchSsmPatchBaselines(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Ssm

	paginator := ssm.NewDescribePatchBaselinesPaginator(svc, nil)
	for paginator.HasMorePages() {
		v, err := paginator.NextPage(ctx, func(o *ssm.Options) {
			o.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- v.BaselineIdentities
	}
	return nil
}
